import AbstractView from "./AbstractView.js";
import { throttle, usernameFromUserID } from "../utils/utils.js";

export default class Chat extends AbstractView {
  constructor() {
    super();
    this.setTitle("Chat");
    this.socket = null;
    this.currentUserID = Number(localStorage.getItem("id"));
    this.RecipientID = undefined;
    this.RecipientName = undefined;
    this.limit = 10;
    this.offset = 10;
    this.previousTime = null;
    this.previousDate = null;
    this._typingIndicator = null;
    this.idleTime = 400;
    this.idleTimer = null;
    this.inputValue;
    this.indicatorState = {
      active: "is-typing-active",
      init: "is-typing-init",
    };
  }

  async startWebsocket() {
    this.socket = new WebSocket(`wss://localhost:8080/api/websocket`);
    this.socket.addEventListener("open", (event) => {
      event.preventDefault();
      console.log("WebSocket connection is open.");
      this.onlineNotifier();
      this.processIncomingWebsocketMessage();
    });

    // Event listener for the logout button click
    const logoutButton = document.getElementById("logout"); // Replace "logoutButton" with the actual ID of your logout button
    if (logoutButton) {
      logoutButton.addEventListener("click", () => {
        console.log("here!");
        this.closeWebSocket();
      });
    }

    // Event listener for page refresh
    window.addEventListener("beforeunload", () => {
      console.log("reloaded here");
      this.closeWebSocket();
    });
  }

  closeWebSocket() {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(
        JSON.stringify({
          type: "connection_close",
          sender: this.currentUserID,
        })
      );
      this.socket.close();
      console.log("WebSocket connection is closed.");
    }
  }

  // List of users to click on to initialise chat
  async userList() {
    const response = await fetch(
      `https://localhost:8080/api/getusers?userId=${this.currentUserID}`,
      {
        credentials: "include",
      }
    );
    const users = await response.json();

    const userContainer = document.getElementById("userContainer");
    const userBox = document.createElement("div");
    userBox.id = "userBox";
    const recentChat = document.createElement("div");
    recentChat.id = "recentChat";
    const alphabeticalChat = document.createElement("div");
    alphabeticalChat.id = "alphabeticalChat";

    userBox.appendChild(recentChat);
    userBox.appendChild(alphabeticalChat);
    userContainer.appendChild(userBox);

    // If there's any entries in users.recentChat...
    if (users.recentChat != null) {
      for (const user of users.recentChat) {
        this.processUserEntry(recentChat, user);
      }
    }

    // If there's any entries in users.alphabetical...
    if (users.alphabetical != null) {
      for (const user of users.alphabetical) {
        this.processUserEntry(alphabeticalChat, user);
      }
    }
  }

  processUserEntry(userContainer, user) {
    // Make a div to put the other two divs in
    let userEntry = document.createElement("div");
    userEntry.classList.add("userEntry");
    userEntry.id = `UserID${user.id}`;

    // Div 1, the persons name with an event listener to load up the chat
    const usernameLink = document.createElement("a");
    usernameLink.classList.add("usernameLink");
    usernameLink.textContent = user.username;
    usernameLink.addEventListener("click", (event) => {
      event.preventDefault();
      this.RecipientID = user.id;
      this.renderHTML();
      // if (userContainer.id == "alphabeticalChat") {
      // 	this.chatInitialiser()
      // }
    });
    // Div 2, the users online status indicator, starts off as hidden
    const statusIndicator = document.createElement("div");
    statusIndicator.id = `statusIndicator${user.id}`;
    statusIndicator.classList.add("statusIndicator");
    statusIndicator.style.display = "none";

    userEntry.appendChild(usernameLink);
    userEntry.appendChild(statusIndicator);

    userContainer.appendChild(userEntry);
  }

  async renderHTML() {
    const chatContainer = document.getElementById("chatContainer");

    if (this.RecipientID == undefined) {
      chatContainer.innerHTML = "Chat (click on Username)";
    } else {
      await this.displayChatHistory();
    }
  }

  async processIncomingWebsocketMessage() {
    if (!this.socket) {
      // Check if the socket is available
      console.error(
        "[processIncomingWebsocketMessage] WebSocket connection is not open."
      );
      return;
    }

    // when user receives a message...
    const handleMessage = async (event) => {
      console.log("Received a WebSocket message:", event.data);
      let message = JSON.parse(event.data);

      if (message.type == "typing") {
        console.log("this is message.sender1", message.sender);
        if (this.RecipientID == message.sender) {
          console.log("this is message.sender2", message.sender);
		  this.showIndicator();
        }
      } else if (message.type == "chat") {
        this.chatHandler(message);
      } else if (message.type == "online-notification") {
        this.onlineHandler(message.onlineUsers);
      }
    };

    this.socket.addEventListener("message", handleMessage);
  }
  // Handle incoming chat messages
  chatHandler(message) {
    if (message.sender == this.RecipientID) {
      console.log(3);
      const chatHistory = document.getElementById("chatHistory");
      console.log(4, chatHistory);
      const senderClassName =
        message.sender === this.currentUserID ? "sent" : "received";
      let messageUsername;
      if (senderClassName === "received") {
        messageUsername = this.RecipientName;
      } else {
        messageUsername = "You";
      }
      const time = this.formatTimestamp(message.time);
      const chatElement = document.createElement("div");
      chatElement.classList.add(senderClassName);
      chatElement.innerHTML = `
				<div id="message-content">
				<div id="recipient">${messageUsername}:</div>
					<div id="body-time-container">
					 <div id="body">${message.body}</div>
					   <div id="time"><i>${time}</i></div>
				   </div>
				</div>
			`;

      // chatHistory displays from the bottom up, this adds new messages to the bottom
      chatHistory.insertBefore(chatElement, chatHistory.firstChild);
      // and then scrolls to the bottom of chatBox
      chatHistory.scrollTop = chatHistory.scrollHeight;
    }

    this.changeUserlistOrder(message.sender);
    this.showMessageReceivedNotification();
  }

  onlineHandler(message) {
    console.log("[onlineHandler]", message);

    setTimeout(() => {
      // sets all status indicators to offline unless their userID is in the list of online users
      const allStatusIndicators =
        document.getElementsByClassName("statusIndicator");
      [...allStatusIndicators].forEach((statusIndicator) => {
        const userId = parseInt(
          statusIndicator.id.replace("statusIndicator", ""),
          10
        );

        statusIndicator.style.display = message.includes(userId)
          ? "block"
          : "none";
      });
    }, 1500);
  }

  changeUserlistOrder(userID) {
    const divToMove = document.getElementById(`UserID${userID}`);
    const recentChat = document.getElementById("recentChat");
    const alphabeticalChat = document.getElementById("alphabeticalChat");
    if (divToMove) {
      if (recentChat != null) {
        if (recentChat.contains(divToMove)) {
          recentChat.removeChild(divToMove);
        }
      } else if (alphabeticalChat != null) {
        if (alphabeticalChat.contains(divToMove)) {
          alphabeticalChat.removeChild(divToMove);
        }
      }
    }
    const divElements = recentChat.querySelectorAll("div");
    const numberOfDivs = divElements.length;

    if (recentChat != null && numberOfDivs > 0) {
      recentChat.insertBefore(divToMove, recentChat.firstChild);
    } else {
      recentChat.appendChild(divToMove);
    }
  }

  async onlineNotifier() {
    console.log("Notifying Server that user is online");

    if (this.socket) {
      this.socket.send(
        JSON.stringify({
          type: "user_online",
          sender: this.currentUserID,
        })
      );
    } else {
      console.log("[onlineNotifier] Websocket not open");
    }
  }

  sendMessage() {
    const messageInput = document.getElementById("messageInput");
    const Message = messageInput.value.trim();

    if (Message !== "" && /\S/.test(Message)) {
      console.log("Sending message:", Message);
      // Create a Date object to get the current date and time
      const now = new Date();

      // Format the date and time as a string, adding a 0 to the left of the numbers, if they were to be single digit
      const formattedTime = `${now.getHours()}:${String(
        now.getMinutes()
      ).padStart(2, "0")}:${String(now.getSeconds()).padStart(
        2,
        "0"
      )} ${now.getDate()}-${now.getMonth() + 1}-${now.getFullYear()}`;

      // Create a message object with the formatted time
      const newMessage = {
        type: "chat",
        body: Message,
        sender: this.currentUserID,
        recipient: this.RecipientID,
        time: formattedTime,
      };

      this.socket.send(JSON.stringify(newMessage));
      messageInput.value = "";
      this.prependMessageToChatHistory(newMessage);
      this.changeUserlistOrder(this.RecipientID);
    }
  }

  // displays chat history (if any) between two users
  async displayChatHistory() {
    this.offset = 10;

    const chatContainer = document.getElementById("chatContainer");
    chatContainer.innerHTML = "";

    const allChat = document.createElement("div");
    allChat.id = "allChat";

    this.RecipientName = await usernameFromUserID(this.RecipientID);

    const recipientHeader = document.createElement("div");
    recipientHeader.id = "recipientHeader";

    const recipientHeaderName = document.createElement("div");
    recipientHeaderName.id = "recipientHeaderName";
    recipientHeaderName.innerHTML = this.RecipientName;

    const recipientHeaderIsTyping = document.createElement("div");
    recipientHeaderIsTyping.id = "recipientHeaderIsTyping";

    const typing = document.createElement("div");
    typing.classList.add("typing");

    typing.innerHTML = `
<span class="typing__bullet"></span>
        <span class="typing__bullet"></span>
        <span class="typing__bullet"></span>
		`;

    recipientHeaderIsTyping.appendChild(typing);

    this._typingIndicator = document.querySelector(".typing");

    // const typing = document.getElementsByClassName("typing")
    console.log(document.getElementsByClassName("typing"));

    recipientHeader.appendChild(recipientHeaderName);
    recipientHeader.appendChild(recipientHeaderIsTyping);

    const chatHistory = document.createElement("div");
    chatHistory.id = "chatHistory";
    const chatTextBox = document.createElement("div");
    chatTextBox.id = "chatTextBox";
    const messageInput = document.createElement("input");
    messageInput.id = "messageInput";
    const sendButton = document.createElement("button");
    sendButton.id = "sendButton";
    sendButton.innerText = "send";
    chatTextBox.appendChild(messageInput);
    chatTextBox.appendChild(sendButton);

    allChat.appendChild(recipientHeader);
    allChat.appendChild(chatHistory);
    allChat.appendChild(chatTextBox);

    chatContainer.appendChild(allChat);

    //deals with sending new messages to the backend when sendButton is clicked or enter is pressed
    chatTextBox.addEventListener("click", (event) => {
      // Check if the clicked element is the sendButton or messageInput
      if (event.target === sendButton) {
        this.sendMessage();
      }
    });

    // Alternatively, you can still keep the "Enter" key functionality for messageInput
    messageInput.addEventListener("keydown", (event) => {
      if (event.key === "Enter") {
        this.sendMessage();
      }
    });

	messageInput.addEventListener("input", (event) => {
		this.sendTypingStatus("typing")
	}); 

    // Load and display an initial set of messages (e.g., 20)
    const initialMessages = await this.fetchMessagesInChunks(0, 10);
    if (initialMessages != null) {
      for (const message of initialMessages) {
        this.appendMessageToChatHistory(message);
      }
    } else {
      // If no messages then add an encouraging message
      const chatElement = document.createElement("div");
      chatElement.innerHTML = `${"- Start your chat journey -"}`;
      chatHistory.appendChild(chatElement);
    }

    // Add a scroll event listener to the chat history container
    const throttleScroll = throttle(() => {
      const scrollThreshold = chatHistory.scrollHeight * 0.9;

      if ((chatHistory.scrollTop = scrollThreshold)) {
        // Load and append more messages
        this.loadMoreMessages(
          this.currentUserID,
          this.RecipientID,
          this.offset,
          this.limit
        );
        this.offset += 10;
      }
    }, 5000);

    chatHistory.addEventListener("scroll", throttleScroll);


    // this.initTypingIndicator(typing);
  }

  appendMessageToChatHistory(message) {
    const chatHistory = document.getElementById("chatHistory");
    const chatElement = document.createElement("div");
    const senderClassName =
      message.sender === this.currentUserID ? "sent" : "received";
    chatElement.classList.add(senderClassName);
    let messageUsername;
    if (senderClassName === "received") {
      messageUsername = this.RecipientName;
    } else {
      messageUsername = "You";
    }
    const time = this.formatTimestamp(message.time);
    chatElement.innerHTML = `
		<div id="message-content">
    		<div id="recipient">${messageUsername}:</div>
   			 <div id="body-time-container">
     			<div id="body">${message.body}</div>
       			<div id="time"><i>${time}</i></div>
   			</div>
		</div>
		`;
    chatHistory.appendChild(chatElement);
  }

  async prependMessageToChatHistory(message) {
    const chatHistory = document.getElementById("chatHistory");
    const chatElement = document.createElement("div");
    const senderClassName =
      message.sender === this.currentUserID ? "sent" : "received";
    chatElement.classList.add(senderClassName);

    let messageUsername;
    if (senderClassName === "received") {
      messageUsername = this.RecipientName;
    } else {
      messageUsername = "You";
    }

    const time = this.formatTimestamp(message.time);
    const messageContent = document.createElement("div");
    messageContent.id = "message-content";

    const recipientDiv = document.createElement("div");
    recipientDiv.id = "recipient";
    recipientDiv.textContent = `${messageUsername}:`;

    const bodyTimeContainerDiv = document.createElement("div");
    bodyTimeContainerDiv.id = "body-time-container";

    const bodyDiv = document.createElement("div");
    bodyDiv.id = "body";
    bodyDiv.textContent = message.body;

    const timeDiv = document.createElement("div");
    timeDiv.id = "time";
    timeDiv.innerHTML = `<i>${time}</i>`;

    bodyTimeContainerDiv.appendChild(bodyDiv);
    bodyTimeContainerDiv.appendChild(timeDiv);

    messageContent.appendChild(recipientDiv);
    messageContent.appendChild(bodyTimeContainerDiv);

    chatElement.appendChild(messageContent);

    // Set background color based on senderClassName
    chatHistory.insertBefore(chatElement, chatHistory.firstChild);
  }

  async loadMoreMessages() {
    if (!this.canScroll) {
      return;
    }

    const nextMessages = await this.fetchMessagesInChunks(
      this.offset,
      this.limit
    );
    if (nextMessages != null) {
      if (nextMessages.length > 0) {
        for (const message of nextMessages) {
          this.appendMessageToChatHistory(message);
        }
      }
    } else {
      this.canScroll = false;
    }
  }

  async fetchMessagesInChunks(offset, limit) {
    const response = await fetch(
      `https://localhost:8080/getChatHistory?user1=${this.currentUserID}&user2=${this.RecipientID}&offset=${offset}&limit=${limit}`,
      {
        credentials: "include", // Ensure cookies are included in the request
      }
    );
    // checkSessionTimeout(response)
    return await response.json();
  }

  canScroll = true;

  async canScrollfunc() {
    if (this.canScroll) {
      return true;
    } else {
      return false;
    }
  }

  // The chatbox for new messages
  getChatTextBoxHTML() {
    return `
				<div id="messages"></div>
				<input type="text" id="messageInput" />
				<button id="sendButton">Send</button>
		`;
  }

  formatTimestamp(timestamp) {
    const splitTimestamp = timestamp.split(" "); // Split the timestamp into time and date parts

    const timePart = splitTimestamp[0]; // "19:12:30"
    const datePart = splitTimestamp[1]; // "25-10-2023"

    // split up the time, ommiting seconds
    const time = timePart.split(":"); // Split the time into hours, minutes, and seconds

    const hours = time[0];
    const minutes = time[1];

    // Split up the date
    const dateParts = datePart.split("-");
    const day = dateParts[0];
    const month = dateParts[1];
    const year = dateParts[2];

    let currentTime = `${hours}:${minutes}`;
    let currentDate = `${day}/${month}/${year}`;
    let formattedTimestamp = `${hours}:${minutes} ${day}/${month}/${year}`;

    if (this.previousTime == currentTime && this.previousDate == currentDate) {
      formattedTimestamp = "";
    } else if (
      this.previousTime != currentTime &&
      this.previousDate == currentDate
    ) {
      formattedTimestamp = `${hours}:${minutes}`;
    } else {
      formattedTimestamp = `${hours}:${minutes} ${day}/${month}/${year}`;
    }

    this.previousTime = currentTime;
    this.previousDate = currentDate;
    return formattedTimestamp;
  }

  showMessageReceivedNotification() {
    // Create a notification element
    const notification = document.createElement("div");
    notification.id = "notification";
    notification.innerText = "New Message Received";

    // Append the notification to the body
    document.body.appendChild(notification);

    setTimeout(() => {
      // Remove the notification element after 5 seconds
      notification.remove();
    }, 5000);
  }

  showIndicator() {
    const typing = document.getElementsByClassName("typing");
    typing[0].classList.add(this.indicatorState.init);
    
  }

  activateIndicator(el) {
    const typing = document.getElementsByClassName("typing");
    typing[0].classList.add(this.indicatorState.active);
    this.inputValue = el.value;
    this.detectIdle(el);
    
  }

  removeIndicator() {
    const typing = document.getElementsByClassName("typing");
    typing[0].classList.remove(
      this.indicatorState.init,
      this.indicatorState.active
    )
  }

  detectIdle(el) {
    const typing = document.getElementsByClassName("typing");
    if (this.idleTimer) {
      clearInterval(this.idleTimer);
    }

    this.idleTimer = setTimeout(() => {
      if (this.getInputCurrentValue(el) === this.inputValue) {
        typing[0].classList.remove(this.indicatorState.active);
      }
    }, this.idleTime);

    
  }

  getInputCurrentValue(el) {
    var currentValue = el.value;
    return currentValue;
  }

  initTypingIndicator() {
    const _input = document.getElementById("messageInput");
    _input.onfocus = () => {
      this.showIndicator();
    };

    _input.onkeyup = () => {
      this.activateIndicator(this);
    };

    _input.onblur = () => {
      this.removeIndicator();
    };
  }

  sendTypingStatus(typingStatus) {
    const message = {
      type: "typing",
      body: typingStatus,
      sender: this.currentUserID,
      recipient: this.RecipientID,
    };
    this.socket.send(JSON.stringify(message));
  }
}

// function sendTypingStatus(typingStatus) {

//     const message = {
//         type: "typing",
//         isTyping: typingStatus,
//         sender: this.currentUserID,
//         recipient: this.RecipientID,
//     };
//     this.socket.send(JSON.stringify(message));
// }

// // Call this function when the sender focuses on chat box
// sendTypingStatus("focusTyping");

// // Call this function when the sender starts typing
// sendTypingStatus("typing");

// // Call this function when the sender stops typing
// sendTypingStatus("notTyping");

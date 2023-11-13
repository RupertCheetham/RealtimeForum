import AbstractView from "./AbstractView.js"
import { throttle, usernameFromUserID } from "../utils/utils.js"




export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
		this.socket = null;
		this.currentUserID = Number(localStorage.getItem("id"))
		this.RecipientID = undefined;
	}

	// List of users to click on to initialise chat
	async userList() {
		const userContainer = document.getElementById("userContainer")
		const userBox = document.createElement("div")
		userBox.id = "userBox"

		const response = await fetch(
			`https://localhost:8080/api/getusers?userId=${this.currentUserID}`,
			{
				credentials: "include",
			}
		)

		const users = await response.json()

		const recentChat = document.createElement("div")
		recentChat.id = "recentChat"

		if (users.recentChat != null) {
			for (const user of users.recentChat) {
				if (user.id != this.currentUserID) {
					let userEntry = document.createElement("div")
					userEntry.id = `UserID${user.id}`
					const usernameLink = document.createElement("a");
					usernameLink.classList.add("chatUserButton");
					usernameLink.textContent = user.username;
					usernameLink.addEventListener("click", (e) => {
						e.preventDefault();
						console.log(user.username)
						console.log(user.id)
						this.RecipientID = user.id
						// this.chatInitialiser(user.id)
						this.displayChatHistory(user.id)
						console.log("I have changed this.RecipientID to ", this.RecipientID)

					})

					userEntry.appendChild(usernameLink);
					recentChat.appendChild(userEntry)
				}
			}
		}
		userBox.appendChild(recentChat)
		userContainer.appendChild(userBox)

		let alphabeticalChat = document.createElement("div")
		alphabeticalChat.id = "alphabeticalChat"

		if (users.alphabetical != null) {
			for (const user of users.alphabetical) {
				if (user.id != this.currentUserID) {
					let userEntry = document.createElement("div")
					userEntry.id = `UserID${user.id}`


					// Create a clickable username link without reloading the page
					const usernameLink = document.createElement("a");
					usernameLink.classList.add("chatUserButton");
					usernameLink.textContent = user.username;
					usernameLink.addEventListener("click", (e) => {
						e.preventDefault();
						console.log(user.username)
						this.RecipientID = user.id
						console.log("I have changed this.RecipientID to ", this.RecipientID)
						this.displayChatHistory(user.id)
					});

					userEntry.appendChild(usernameLink);
					alphabeticalChat.appendChild(userEntry)
				}
			}
		}
		userBox.appendChild(alphabeticalChat)
		userContainer.appendChild(userBox)
	}

	async startWebsocket() {
		this.socket = new WebSocket(`wss://localhost:8080/api/websocketChat`);
		this.socket.addEventListener("open", (event) => {
			event.preventDefault();
			console.log("WebSocket connection is open.");
			this.onlineStatusHandler()
			this.webSocketChat()
		});

	}

	async renderHTML() {
		const chatContainer = document.getElementById("chatContainer");
		const allChat = document.createElement("div");
		allChat.id = "allChat";
		const chatHistory = document.createElement("div");
		chatHistory.id = "chatHistory";
		chatHistory.innerHTML = "Chat (click on Username)"
		allChat.appendChild(chatHistory)
		const chatTextBox = document.createElement("div");
		chatTextBox.id = "chatTextBox"
		chatTextBox.innerHTML = this.getChatTextBoxHTML();
		chatTextBox.style.display = "none";
		allChat.appendChild(chatTextBox)

		chatContainer.appendChild(allChat);

	}

	async webSocketChat() {

		//deals with sending new messages to the backend when sendButton is clicked or enter is pressed
		document.getElementById("sendButton").addEventListener("click", () => {
			this.sendMessage();
		});
		document.getElementById("messageInput").addEventListener("keydown", function (event) {
			if (event.key === "Enter") {
				this.sendMessage()
			}
		})


		if (!this.socket) {
			// Check if the socket is available
			console.error("[webSocketChat] WebSocket connection is not open.");
			return;
		}

		// when chat receives a message...
		const handleMessage = async (event) => {
			console.log("Received a WebSocket message:", event.data);
			let message = JSON.parse(event.data);
console.log("message.recipient", message.recipient)
console.log("this.RecipientID", this.RecipientID)
			// rudimentary notification system
			// const RecipientName = await usernameFromUserID(message.recipient);
			// alert("Message: " + message.body +  " from " + RecipientName)
			if (message.recipient == this.RecipientID || message.sender == this.RecipientID) {
				// Handle incoming messages
				let chatElement = document.createElement("div")
				const senderClassName = message.sender === this.currentUserID ? "sent" : "received"
				chatElement.classList.add(senderClassName)
				const time = this.formatTimestamp(message.time)
				chatElement.innerHTML = `
		<div id="message-content">
			<div id="body">${message.body}</div>
			<div id="time"><i>${time}</i></div>
		</div>
	`

				// chatHistory displays from the bottom up, this adds new messages to the bottom
				chatHistory.insertBefore(chatElement, chatHistory.firstChild)

				// and then scrolls to the bottom of chatBox
				chatHistory.scrollTop = chatHistory.scrollHeight

			}


			const divToMove = document.getElementById(`UserID${message, message.recipient}`)
			const recentChat = document.getElementById("recentChat")
			const alphabeticalChat = document.getElementById("alphabeticalChat")

			if (divToMove) {
				if (recentChat != null) {
					if (recentChat.contains(divToMove)) {
						recentChat.removeChild(divToMove)
					}
				} else if (alphabeticalChat != null) {
					if (alphabeticalChat.contains(divToMove)) {
						alphabeticalChat.removeChild(divToMove)
					}
				}
			}
			console.log("recentChat", recentChat)
			if (recentChat != null) {
				console.log(1)
				const divElements = recentChat.querySelectorAll("div");
				const numberOfDivs = divElements.length;
				console.log("Number of divs inside parentDiv: " + numberOfDivs);
				if (numberOfDivs > 0)
					console.log(2)
				console.log("Number of divs inside parentDiv: " + numberOfDivs);
				console.log("recentChat.firstChild", recentChat.firstChild)
				recentChat.insertBefore(divToMove, recentChat.firstChild);
			} else {
				recentChat.appendChild(divToMove)
			}
		}

		this.socket.addEventListener("message", handleMessage);

	}

	async onlineStatusHandler() {
		console.log("Notifying Server that user is online")
		this.socket.send(
			JSON.stringify({
				type: "user_online",
				body: "",
				sender: this.currentUserID,
			})
		)
	}

	chatInitialiser() {
		if (!this.socket) {
			// Check if the socket is available
			console.error("[chatInitialiser] WebSocket connection is not open.");
			return;
		}
		console.log("Priming Chat")
		this.socket.send(
			JSON.stringify({
				type: "chat_init",
				body: "",
				sender: this.currentUserID,
				recipient: this.RecipientID,
			})
		)
	}


	sendMessage() {
		const messageInput = document.getElementById("messageInput")
		const Message = messageInput.value.trim()

		if (Message !== "" && /\S/.test(Message)) {
			console.log("Sending message:", Message)
			this.socket.send(
				JSON.stringify({
					type: "chat",
					body: Message,
					sender: this.currentUserID,
					recipient: this.RecipientID,
				})
			)
			messageInput.value = ""
		}
	}
	// displays chat history (if any) between two users
	async displayChatHistory() {
		console.log("displayChatHistory")
		// Make the chatTextBox visible
		const chatTextBox = document.getElementById("chatTextBox");
		chatTextBox.style.display = "block";
		let messageOffset = 10
		let limit = 10
		const chatHistory = document.getElementById("chatHistory")
		chatHistory.innerHTML = ""
		console.log("[displayChatHistory] allChat", allChat)
		const RecipientName = await usernameFromUserID(this.RecipientID)
		// Create the <h1> element with RecipientName
		const recipientHeader = document.createElement("h1");
		recipientHeader.id = "recipient";
		recipientHeader.className = "chat-font";
		recipientHeader.textContent = RecipientName;
		console.log(this.RecipientID)

		// Load and display an initial set of messages (e.g., 20)
		const initialMessages = await this.fetchMessagesInChunks(
			0,
			10
		)
		if (initialMessages != null) {
			for (const message of initialMessages) {
				this.appendMessageToChatHistory(message, this.currentUserID)
			}
		} else {
			// If no messages then add an encouraging message
			const chatElement = document.createElement("div")
			chatElement.innerHTML = `${"- Start your chat journey -"}`
			chatHistory.appendChild(chatElement)
		}

		// Add a scroll event listener to the chat history container
		const throttleScroll = throttle(() => {
			const scrollThreshold = chatHistory.scrollHeight * 0.3

			if (chatHistory.scrollTop <= scrollThreshold) {
				// Load and append more messages
				this.loadMoreMessages(this.currentUserID, this.RecipientID, messageOffset, limit)
				messageOffset += 10
			}
		}, 100)

		chatHistory.addEventListener("scroll", throttleScroll)
	}

	async appendMessageToChatHistory(message) {
		const chatHistory = document.getElementById("chatHistory")
		const chatElement = document.createElement("div")
		const senderClassName = message.sender === this.currentUserID ? "sent" : "received"
		chatElement.classList.add(senderClassName)
		const time = this.formatTimestamp(message.time)
		chatElement.innerHTML = `<div id="message-content">
				<div id="body">${message.body}</div>
				<div id="time"><i>${time}</i></div>
			</div>`
		chatHistory.appendChild(chatElement)
	}

	async loadMoreMessages(messageOffset, limit) {
		if (!this.canScroll) {
			return
		}

		const nextMessages = await this.fetchMessagesInChunks(
			messageOffset,
			limit
		)
		if (nextMessages != null) {
			if (nextMessages.length > 0) {
				for (const message of nextMessages) {
					this.appendMessageToChatHistory(message)
				}
			}
		} else {
			this.canScroll = false

		}
	}

	async fetchMessagesInChunks(offset, limit) {
		const response = await fetch(
			`https://localhost:8080/getChatHistory?user1=${this.currentUserID}&user2=${this.RecipientID}&offset=${offset}&limit=${limit}`,
			{
				credentials: "include", // Ensure cookies are included in the request
			}
		)
		// checkSessionTimeout(response)
		return await response.json()
	}

	canScroll = true

	async canScrollfunc() {
		if (this.canScroll) {
			return true
		} else {
			return false
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


	previousTime = null
	previousDate = null

	formatTimestamp(timestamp) {
		const splitTimestamp = timestamp.split(" ") // Split the timestamp into time and date parts

		const timePart = splitTimestamp[0] // "19:12:30"
		const datePart = splitTimestamp[1] // "25-10-2023"

		// split up the time, ommiting seconds
		const time = timePart.split(":") // Split the time into hours, minutes, and seconds

		const hours = time[0]
		const minutes = time[1]

		// Split up the date
		const dateParts = datePart.split("-")
		const day = dateParts[0]
		const month = dateParts[1]
		const year = dateParts[2]

		let currentTime = `${hours}:${minutes}`
		let currentDate = `${day}/${month}/${year}`
		let formattedTimestamp = `${hours}:${minutes} ${day}/${month}/${year}`

		if (this.previousTime == currentTime && this.previousDate == currentDate) {
			formattedTimestamp = ""
		} else if (
			this.previousTime != currentTime &&
			this.previousDate == currentDate
		) {
			formattedTimestamp = `${hours}:${minutes}`
		} else {
			formattedTimestamp = `${hours}:${minutes} ${day}/${month}/${year}`
		}

		this.previousTime = currentTime
		this.previousDate = currentDate
		return formattedTimestamp
	}
}

import AbstractView from "./AbstractView.js"
import { userIDFromSessionID, usernameFromUserID } from "../utils/utils.js"

export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
	}


	async userList() {
		const currentUser = await userIDFromSessionID()
		const userContainer = document.getElementById("userContainer");
		userContainer.innerHTML = "";
		const userBox = document.createElement("div");

		const response = await fetch("https://localhost:8080/api/getusers", {
			credentials: "include", // Ensure cookies are included in the request
		});

		const users = await response.json();



		for (const user of users) {
			if (user.id != currentUser) {
				let userEntry = document.createElement("div");
				userEntry.id = "UserBox";

				userEntry.innerHTML = `
				<a href="/main?userId=${user.id}" class="chatUserButton">${user.username}</a>
				`;

				userBox.appendChild(userEntry);
			}

		}
		userContainer.appendChild(userBox);

	}

	async renderHTML() {
		const chatContainer = document.getElementById("chatContainer");
		const RecipientID = await this.getRecipientIDFromURL()
		console.log("in Chat, RecipientID", RecipientID)
		if (RecipientID != 0) {
			const Recipient = await usernameFromUserID(RecipientID)
			const chatTextBox = this.getChatTextBoxHTML();
			chatContainer.innerHTML = `
			
			<h1 id="recipient" class = "chat-font"> ${Recipient}</h1>
			<div id="chatHistory"></div>
			${chatTextBox}
			`
			await this.webSocketChat()
		}
	}

	// Function to extract a query parameter from the URL
	async getRecipientIDFromURL() {
		const queryString = window.location.search;
		const urlParams = new URLSearchParams(queryString);
		return Number(urlParams.get("userId"));
	}

	async webSocketChat() {

		const Sender = await userIDFromSessionID()

		const Recipient = await this.getRecipientIDFromURL()

		const socket = new WebSocket(`wss://localhost:8080/chat?sender=${Sender}&recipient=${Recipient}`);

		socket.addEventListener("open", (event) => {
			event.preventDefault();
			console.log("WebSocket connection is open.");
		});

		this.displayChatHistory(Sender, Recipient)

		// when chat receives a message...
		socket.addEventListener("message", (event) => {
			console.log("Received a WebSocket message:", event.data);

			let chat = JSON.parse(event.data);

			// Handle incoming messages

			let chatElement = document.createElement("div");
			const senderClassName = chat.sender === Sender ? "sent" : "received";
			chatElement.classList.add(senderClassName);

			chatElement.innerHTML = `
	  ${chat.body} <b>Time: </b> <i>${chat.time}</i>
	`;

			chatHistory.appendChild(chatElement);

			// Scroll to the bottom of chatBox
			chatHistory.scrollTop = chatHistory.scrollHeight;
		});


		document.getElementById("sendButton").addEventListener("click", () => {
			const messageInput = document.getElementById("messageInput");
			// Get the value of the input and trim leading/trailing spaces
			const Message = messageInput.value.trim();

			if (Message !== "") {
				// Check if the message contains at least one non-space character
				if (/\S/.test(Message)) {
					console.log("Sending message:", Message);
					// Send the message to the server via WebSocket
					socket.send(JSON.stringify({
						type: "chat",
						body: Message,
						sender: Sender,
						recipient: Recipient
					}));
					// Clear the input field
					messageInput.value = "";
				}
			}


		});

	}

	// displays chat history (if any) between two users
	async displayChatHistory(user1, user2) {

		const chatHistory = document.getElementById("chatHistory");

		const response = await fetch(`https://localhost:8080/getChatHistory?user1=${user1}&user2=${user2}`, {
			credentials: "include", // Ensure cookies are included in the request
		})
		const currentUser = user1
		const chat = await response.json();


		if (chat != null) {
			for (const message of chat) {
				let chatElement = document.createElement("div");
				const senderClassName = message.sender === currentUser ? "sent" : "received";
				chatElement.classList.add(senderClassName);

				chatElement.innerHTML = `
       ${message.body} <b>Time: </b> <i>${message.time}</i>
    `;


				chatHistory.appendChild(chatElement);
			}
			chatHistory.scrollTop = chatHistory.scrollHeight;

		} else {
			let chatElement = document.createElement("div");
			chatElement.classList.add("sent");

			chatElement.innerHTML = `
		- Your Chat Starts Here -
	  `;
			chatHistory.appendChild(chatElement);
		}
	}

	// The chatbox for new messages
	getChatTextBoxHTML() {
		return `
	<div id="chatTextBox">
    <div id="messages"></div>
    <input type="text" id="messageInput" />
    <button id="sendButton">Send</button>
  </div>
`
	}
}






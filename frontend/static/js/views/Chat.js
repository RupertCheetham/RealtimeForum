import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"
import { userIDFromSessionID, usernameFromUserID } from "../utils/utils.js"

export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
	}

	async renderHTML() {

		const nav = new Nav() // Create an instance of the Nav class
		const navHTML = await nav.renderHTML() // Get the HTML content for the navigation
		const RecipientID = await this.getRecipientIDFromURL()
		const Recipient = await usernameFromUserID(RecipientID)
		const chatTextBox = getChatTextBoxHTML();
		return `
		<body>
			${navHTML}
			<!DOCTYPE html>
        	<html>
       	 	<h1 id="chat-font" class = "chat-font"> cHaT iS hErE</h1>
			<div id="chatContainer" class="chatContainer">
				<div class="user-info">
				<img src="frontend/static/js/views/icons8-user-94.png" />
					<h1 id="recipient" class = "chat-font"> ${Recipient}</h1>
				</div>
				<div class="chat-box">
				<!-- Chat messages go here -->
				</div>
				<div class="inputContainer">
				${chatTextBox}
				</div>
			</div>
			</div>
		
			</html>
		</body>
        `
	}

	// Function to extract a query parameter from the URL
	async getRecipientIDFromURL() {
		const queryString = window.location.search;
		const urlParams = new URLSearchParams(queryString);
		return Number(urlParams.get("userId"));
	}

	async webSocketStuff() {


		const Sender = await userIDFromSessionID()

		const Recipient = await this.getRecipientIDFromURL()

		const socket = new WebSocket(`wss://localhost:8080/chat?sender=${Sender}&recipient=${Recipient}`);

		socket.addEventListener("open", (event) => {
			event.preventDefault();
			console.log("WebSocket connection is open.");
		});

		this.displayChatContainer(Sender, Recipient)

		// when chat receives a message...
		socket.addEventListener("message", (event) => {
			console.log("Received a WebSocket message:", event.data);

			let chat = JSON.parse(event.data)
			console.log("hello", chat.message, chat.sender)
			// Handle incoming messages
			const /* `chatB` is a variable that is used to store the HTML content for the chat box. It is a
			function that returns a string of HTML code representing the chat box. */
			chatBox = document.getElementById("chatBox");

			let chatElement = document.createElement("div");
			const senderClassName = chat.sender === Sender ? "sent" : "received";
			chatElement.classList.add(senderClassName);

			chatElement.innerHTML = `
			${chat.message} <b>Time: </b> <i>${chat.time}</i>
    `;
			chatBox.appendChild(chatElement)
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
						message: Message,
						sender: Sender,
						recipient: Recipient
					}));
					// Clear the input field
					messageInput.value = "";
				}
			}


		});

	}

	async displayChatContainer(user1, user2) {
		const chatContainer = document.getElementById("chatContainer");
		//chatContainer.innerHTML = "";

		const response = await fetch(`https://localhost:8080/getChatHistory?user1=${user1}&user2=${user2}`, {
			credentials: "include", // Ensure cookies are included in the request
		})
		const currentUser = user1
		const chats = await response.json();

		const chatBox = document.createElement("div");
		chatBox.className = "chatBox";
		chatBox.id = "chatBox";

		if (chats != null) {
			for (const chat of chats) {
				let chatElement = document.createElement("div");
				const senderClassName = chat.sender === currentUser ? "sent" : "received";
				chatElement.classList.add(senderClassName);

				// Determine the appropriate class name based on the sender


				chatElement.innerHTML = `
       ${chat.message} <b>Time: </b> <i>${chat.time}</i>
    `;


				chatBox.appendChild(chatElement);
			}

		} else {
			let chatElement = document.createElement("div");
			chatElement.classList.add("sent");

			chatElement.innerHTML = `
		No messages yet
	  `;
			chatBox.appendChild(chatElement);
		}
		chatContainer.appendChild(chatBox)
	}
}


// The chatbox for new messages
function getChatTextBoxHTML() {
	return `
	<div id="chat">
    <div id="messages"></div>
    <input type="text" id="messageInput" />
    <button id="sendButton">Send</button>
  </div>
`
}

export async function userList() {
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
			<a href="/chat?userId=${user.id}" class="chatUserButton">${user.username}</a>
			`;

			userBox.appendChild(userEntry);
		}

	}
	userContainer.appendChild(userBox);

}

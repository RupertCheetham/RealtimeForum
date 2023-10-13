import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"
import { userIDFromSessionID } from "../utils/utils.js"

export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
	}

	async renderHTML() {

		const nav = new Nav() // Create an instance of the Nav class
		const navHTML = await nav.renderHTML() // Get the HTML content for the navigation
		const chatTextBox = getChatTextBoxHTML();
		return `
		${navHTML}
        <h1 id="chat-font" class = "chat-font"> cHaT iS hErE</h1>
		<div id="chatContainer"></div>
		${chatTextBox}
        `
	}

	// Function to extract a query parameter from the URL
	async getUserIDFromURL() {
		const queryString = window.location.search;
		const urlParams = new URLSearchParams(queryString);
		return Number(urlParams.get("userId"));
	}

	async webSocketStuff() {

		// Example Sender
		const Sender = await userIDFromSessionID()
		console.log("This is Sender:", Sender)
		// Example Receiver
		const Recipient = await this.getUserIDFromURL()
		console.log("This is  Recipient:", Recipient)

		const socket = new WebSocket("wss://localhost:8080/chat");

		socket.addEventListener("open", (event) => {
			event.preventDefault();
			console.log("WebSocket connection is open.");

			// WebSocket connection established
		});

		this.displayChatContainer(Sender, Recipient)

		socket.addEventListener("message", (event) => {
			console.log("Received a WebSocket message:", event.data);
			// Handle incoming messages
		});

		document.getElementById("sendButton").addEventListener("click", () => {
			const messageInput = document.getElementById("messageInput");
			const Message = messageInput.value;
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

		});

	}
	async displayChatContainer(user1, user2) {
		const chatContainer = document.getElementById("chatContainer");
		chatContainer.innerHTML = "";

		const response = await fetch(`https://localhost:8080/getChatHistory?user1=${user1}&user2=${user2}`, {
			credentials: "include", // Ensure cookies are included in the request
		})
		const currentUser = user1
		const chats = await response.json();

		const chatBox = document.createElement("div");
		chatBox.className = "chatBox";
		// const chatTarget =  await fetch(`https://localhost:8080/getUsernameFromUswerID?userID=${user2}`, {
		// 	credentials: "include", // Ensure cookies are included in the request
		// })

		if (chats != null) {
			for (const chat of chats) {
				let chatElement = document.createElement("div");
				const senderClassName = chat.sender === currentUser ? "sent" : "received";
				chatElement.classList.add(senderClassName);

				// Determine the appropriate class name based on the sender
				

				chatElement.innerHTML = `
        <b>Username: </b> ${chat.sender}, <b>Message: </b> ${chat.message}, <b>Time: </b> ${chat.time}
    `;


				chatBox.appendChild(chatElement);
			}

		} else {
			let chatElement = document.createElement("div");
			chatElement.id = "Chat" + chat.id;
			chatElement.classList.add("chat");

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

			console.log("this is user.username:", user.username)
			userBox.appendChild(userEntry);
		}

	}
	userContainer.appendChild(userBox);

}

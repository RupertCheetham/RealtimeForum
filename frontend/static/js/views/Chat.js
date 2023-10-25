import AbstractView from "./AbstractView.js"
import { userIDFromSessionID, usernameFromUserID, throttle } from "../utils/utils.js"

export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
	}

	// List of users to click on to initialise chat
	async userList() {
		const currentUser = await userIDFromSessionID()
		const userContainer = document.getElementById("userContainer");
		userContainer.innerHTML = "";
		const userBox = document.createElement("div");

		const response = await fetch("https://localhost:8080/api/getusers", {
			credentials: "include",
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
		const Sender = await userIDFromSessionID();
		const Recipient = await this.getRecipientIDFromURL();
		const socket = new WebSocket(`wss://localhost:8080/chat?sender=${Sender}&recipient=${Recipient}`);



		socket.addEventListener("open", (event) => {
			event.preventDefault();
			console.log("WebSocket connection is open.");
		});


		// socket.addEventListener("close", (event) => {
		// 	event.preventDefault();
		// 	console.log("WebSocket connection is closed.");
		// 	console.log("socket: ", socket)
		// 	socket.send(JSON.stringify({ type: "disconnect", body: socket}));
		// });

		this.displayChatHistory(Sender, Recipient);

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

		// Handle the WebSocket connection's close event
		socket.addEventListener("close", () => {
				// An abrupt connection closure, e.g., a page reload
				// Send a disconnect message to the server before closing the connection
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
	// Modify the displayChatHistory function
	async displayChatHistory(user1, user2) {
		const chatHistory = document.getElementById("chatHistory");

		// Function to fetch messages in chunks
		async function fetchMessagesInChunks(offset, limit) {
			const response = await fetch(`https://localhost:8080/getChatHistory?user1=${user1}&user2=${user2}&offset=${offset}&limit=${limit}`, {
				credentials: "include", // Ensure cookies are included in the request
			});
			return await response.json();
		}

		// Load and display an initial set of messages (e.g., 20)
		const initialMessages = await fetchMessagesInChunks(0, 10);
		if (initialMessages.length > 0) {
			for (const message of initialMessages.reverse()) {
				appendMessageToChatHistory(message);
				chatHistory.scrollTop = chatHistory.scrollHeight;
			}
		}

		// Define a variable to keep track of the message offset
		let messageOffset = 10;

		// Function to append a single message to the chat history container
		function appendMessageToChatHistory(message) {
			const chatElement = document.createElement("div");
			const senderClassName = message.sender === user1 ? "sent" : "received";
			chatElement.classList.add(senderClassName);
			chatElement.innerHTML = `${message.body} <b>Time: </b> <i>${message.time}</i>`;
			chatHistory.appendChild(chatElement);
			chatHistory.scrollTop = chatHistory.scrollHeight;
		}

		// Define the scroll threshold for loading more messages (e.g., 10% of the chat history container's height)
		const scrollThreshold = chatHistory.scrollHeight * 0.1;

		// Function to load and append more messages
		async function loadMoreMessages() {
			const nextMessages = await fetchMessagesInChunks(messageOffset, 10);
			if (nextMessages.length > 0) {
				for (const message of nextMessages.reverse()) {
					appendMessageToChatHistory(message);
				}
				messageOffset += 10;
			}
		}

		// Add a scroll event listener to the chat history container
		chatHistory.addEventListener("scroll", () => {
			if (chatHistory.scrollTop < scrollThreshold) {
				// Load and append more messages
				loadMoreMessages();
			}
		});
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






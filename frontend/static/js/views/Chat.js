import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"

export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
	}

	async renderHTML() {
		
		const nav = new Nav() // Create an instance of the Nav class
		const navHTML = await nav.renderHTML() // Get the HTML content for the navigation
		const chat = getChatHTML();
		return `
		${navHTML}
        <h1 id="chat-font" class = "chat-font"> cHaT iS hErE</h1>
		${chat}
		<div id="chatContainer"></div>
        `
	}

	stylingBlue() {
		const chatFont = document.getElementById("chat-font")
		chatFont.style.color = "blue"
	}

	stylingBorder() {
		const chatFont = document.getElementById("chat-font")
		chatFont.style.border = "5px solid Red"
	}

	async webSocketStuff() {

		// Example Sender
		const Sender = 5
		// Example Receiver
		const Recipient = 2

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
		console.log("I am in displayChatContainer.")
		const chatContainer = document.getElementById("chatContainer");
		chatContainer.innerHTML = "";
	
		const response = await fetch(`https://localhost:8080/getChatHistory?user1=${user1}&user2=${user2}`, {
			credentials: "include", // Ensure cookies are included in the request
		})
		console.log("I am next to JSON", response)
		const chats = await response.json();
		console.log("I am chat", chats)
		console.log("I am chat length", chats.length)
	
		for (const chat of chats) {
			let chatElement = document.createElement("div");
			chatElement.id = "Chat" + chat.id;
			chatElement.classList.add("chat");
		
		 console.log("HERE")
			chatElement.innerHTML = `
	
				${chat.sender}
				${chat.message}
				${chat.time}
				<li><b>Username:</b> 	${chat.sender}</li>
			  `;
		
			  chatContainer.appendChild(chatElement);
		}
	}
}



function getChatHTML() {
	return `
return <div id="chat">
<div id="messages"></div>
<input type="text" id="messageInput" />
<button id="sendButton">Send</button>
</div>
`
}




//handle func for chats specific to uuid
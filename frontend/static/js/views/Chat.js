import AbstractView from "./AbstractView.js"
import { throttle, usernameFromUserID } from "../utils/utils.js"




export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
		this.socket = null;
		this.currentUserID = Number(localStorage.getItem("id"))
	}



	// List of users to click on to initialise chat
	async userList() {
		const userContainer = document.getElementById("userContainer")
		// userContainer.innerHTML = "";
		const userBox = document.createElement("div")
		userBox.id = "userBox"

		const response = await fetch(
			`https://localhost:8080/api/getusers?userId=${this.currentUserID}`,
			{
				credentials: "include",
			}
		)

		// checkSessionTimeout(response)

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
						this.renderHTML(user.id)
					})

					// 	userEntry.innerHTML = `
					// <a href="/main?userId=${user.id}" class="chatUserButton">${user.username}</a>
					// `
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
						this.renderHTML(user.id)


						//   this.loadUserChat(user.id);
					});







					// 	userEntry.innerHTML = `
					// <a href="/main?userId=${user.id}" class="chatUserButton">${user.username}</a>
					// `
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
		});
	}

	async renderHTML(RecipientID) {



		const chatContainer = document.getElementById("chatContainer")

		if (RecipientID != 0) {

			const RecipientName = await usernameFromUserID(RecipientID)
			const chatTextBox = this.getChatTextBoxHTML()
			chatContainer.innerHTML = `
      <div class = "allChat"
      <h1 id="recipient" class = "chat-font"> ${RecipientName}</h1>
			<div id="chatHistory"></div>
			${chatTextBox}
      </div>
			`
			await this.displayChatHistory(this.currentUserID, RecipientID)

		}
		this.connectionInitialiser(RecipientID)
		this.webSocketChat(RecipientID)
	}

	// // Function to extract a query parameter from the URL
	// async getRecipientIDFromURL() {
	// 	const queryString = window.location.search
	// 	const urlParams = new URLSearchParams(queryString)
	// 	return Number(urlParams.get("userId"))
	// }

	async webSocketChat(RecipientID) {

		if (!this.socket) {
			// Check if the socket is available
			console.error("[webSocketChat] WebSocket connection is not open.");
			return;
		}




		// when chat receives a message...
		this.socket.addEventListener("message", (event) => {
			console.log("Received a WebSocket message:", event.data)

			let message = JSON.parse(event.data)

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

			const divToMove = document.getElementById(`UserID${RecipientID}`)
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

			if (recentChat != null) {
				recentChat.insertBefore(divToMove, recentChat.firstChild)
			} else {
				recentChat.appendChild(divToMove)
			}
		})

		//deals with sending new messages to the backend when sendButton is clicked or enter is pressed
		document.getElementById("sendButton").addEventListener("click", () => {
			console.log("her3")
			const messageInput = document.getElementById("messageInput");
			const message = messageInput.value.trim();
			this.sendMessage(RecipientID); // Call the sendMessage method of the Chat class
		});
		document.getElementById("messageInput").addEventListener("keydown", function (event) {
			if (event.key === "Enter") {
				this.sendMessage(RecipientID)
			}
		})

	}


	connectionInitialiser(RecipientID) {

		console.log("Priming Chat")
		this.socket.send(
			JSON.stringify({
				type: "connection_init",
				body: "Connection initiated",
				sender: this.currentUserID,
				recipient: RecipientID,
			})
		)
		messageInput.value = ""
	}


	sendMessage(RecipientID) {
		const messageInput = document.getElementById("messageInput")
		const Message = messageInput.value.trim()

		if (Message !== "" && /\S/.test(Message)) {
			console.log("Sending message:", Message)
			this.socket.send(
				JSON.stringify({
					type: "chat",
					body: Message,
					sender: this.currentUserID,
					recipient: RecipientID,
				})
			)
			messageInput.value = ""
		}
	}
	// displays chat history (if any) between two users
	async displayChatHistory(user1, user2) {
		let messageOffset = 10
		let limit = 10
		const chatHistory = document.getElementById("chatHistory")

		// Load and display an initial set of messages (e.g., 20)
		const initialMessages = await this.fetchMessagesInChunks(
			user1,
			user2,
			0,
			10
		)
		if (initialMessages != null) {
			for (const message of initialMessages) {
				this.appendMessageToChatHistory(message, user1)
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
				this.loadMoreMessages(user1, user2, messageOffset, limit)
				messageOffset += 10
			}
		}, 100)

		chatHistory.addEventListener("scroll", throttleScroll)
	}

	async appendMessageToChatHistory(message, user1) {
		const chatElement = document.createElement("div")
		const senderClassName = message.sender === user1 ? "sent" : "received"
		chatElement.classList.add(senderClassName)
		const time = this.formatTimestamp(message.time)
		chatElement.innerHTML = `<div id="message-content">
				<div id="body">${message.body}</div>
				<div id="time"><i>${time}</i></div>
			</div>`
		chatHistory.appendChild(chatElement)
	}

	async loadMoreMessages(user1, user2, messageOffset, limit) {
		if (!this.canScroll) {
			return
		}

		const nextMessages = await this.fetchMessagesInChunks(
			user1,
			user2,
			messageOffset,
			limit
		)
		if (nextMessages != null) {
			if (nextMessages.length > 0) {
				for (const message of nextMessages) {
					this.appendMessageToChatHistory(message, user1)
				}
			}
		} else {
			this.canScroll = false

		}
	}

	async fetchMessagesInChunks(user1, user2, offset, limit) {
		const response = await fetch(
			`https://localhost:8080/getChatHistory?user1=${user1}&user2=${user2}&offset=${offset}&limit=${limit}`,
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
	<div id="chatTextBox">
    <div id="messages"></div>
    <input type="text" id="messageInput" />
    <button id="sendButton">Send</button>
  </div>
`
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

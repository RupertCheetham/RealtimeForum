import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"


export default class Chat extends AbstractView {
	constructor() {
		super()
		this.setTitle("Chat")
	}

	async getHTML() {
		const nav = new Nav() // Create an instance of the Nav class
		const navHTML = await nav.getHTML() // Get the HTML content for the navigation

		return `
		${navHTML}
        <h1 id="chat-font"> CHat is here</h1>
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
}
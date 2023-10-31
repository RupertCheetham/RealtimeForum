import AbstractView from "./AbstractView.js"
import { userNameFromSessionID } from "../utils/utils.js";

export default class Nav extends AbstractView {
	constructor() {
		super()
		this.setTitle("Posts")
	}

	async renderHTML() {
		
		const username = await userNameFromSessionID()

		return `
		<nav id="nav" class="nav">
			<a href="/" class="nav-link" data-link id="logout">Logout</a>
			<a href="/main" class="nav-link" data-link>Home</a>
			<a href="/chat" class="nav-link" data-link>Chat</a>
			<span id="username">${username}</span>
		</nav>
    `
	}

	
}

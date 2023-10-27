import AbstractView from "./AbstractView.js"
import { getUserName } from "../utils/utils.js"

export default class Nav extends AbstractView {
	constructor() {
		super()
		this.setTitle("Posts")
	}

	async renderHTML() {
		const username = await getUserName()

		return `
			<nav id="nav" class="nav">
				<a href="/" class="nav-link" data-link id="logout">Logout</a>
				<a href="/posts" class="nav-link" data-link>Posts</a>
				<a href="/chat" class="nav-link" data-link>Chat</a>
				<span id="cookie-value">${username}</span>
			</nav>
		  `
	}
}

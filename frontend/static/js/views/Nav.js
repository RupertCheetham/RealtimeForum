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
			<span id="username">${username}</span>
		</nav>
    `
	}
}

// Nav with home button still there
// `
// <nav id="nav" class="nav">
// 	<a href="/" class="nav-link" data-link id="logout">Logout</a>
// 	<a href="/main" class="nav-link" data-link>Home</a>
// 	<span id="username">${username}</span>
// </nav>
// `

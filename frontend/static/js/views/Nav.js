import AbstractView from "./AbstractView.js"

export default class Nav extends AbstractView {
	constructor() {
		super()
		this.setTitle("Posts")
	}

	async renderHTML() {
		return `
		<nav id="nav" class="nav">
			<a href="/" class="nav-link" data-link>Logout</a>
			<a href="/posts" class="nav-link" data-link>Posts</a>
			<a href="/chat" class="nav-link" data-link>Chat</a>
		</nav>
    `
	}
}

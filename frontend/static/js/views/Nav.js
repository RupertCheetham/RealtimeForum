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
				<span id="cookie-value">${username}</span>
			</nav>
		  `
	}

	async logout() {
		let logoutbtn = document.getElementById("logout")
		logoutbtn.addEventListener("click", (event) => {
			event.preventDefault()
			fetch("https://localhost:8080/api/logout", {
				headers: {
					Accept: "application/json",
					"Content-Type": "application/json",
				},
				credentials: "include", // Ensure cookies are included in the request
			})
		})
	}
}

import AbstractView from "./AbstractView.js"

export default class Nav extends AbstractView {
	constructor() {
		super()
		this.setTitle("Posts")
	}

	async renderHTML() {
		const username = localStorage.getItem("username")

		return `
		<header>
			<div class="nav-container">
				<div class="nav-wrapper">
					<a href="/" class="nav-link" data-link id="logout">Logout</a>
					<a href="/user" id="cookie-value">${username}</a>
				</div>
			</div>
		</header>
		  `
	}

	async logout() {
		let logoutbtn = document.getElementById("logout")
		logoutbtn.addEventListener("click", (event) => {
			event.preventDefault()
			localStorage.clear()
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

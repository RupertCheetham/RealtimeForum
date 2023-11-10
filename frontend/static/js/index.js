import { getCookie } from "./utils/utils.js"
import Auth from "./views/Auth.js"
import MainPage from "./views/MainPage.js"

const navigateTo = (url) => {
	history.pushState(null, null, url)
	router()
}

const timeout = 5

const router = async () => {
	const routes = [
		{ path: "/", view: Auth },
		{ path: "/main", view: MainPage },
	]

	// test each route for potential match
	const potentialMatches = routes.map((route) => {
		return {
			route: route,
			isMatch: location.pathname === route.path,
		}
	})

	let match = potentialMatches.find((potentialMatch) => potentialMatch.isMatch)

	if (!match) {
		match = {
			route: routes[0],
			isMatch: true,
		}
	}

	// const view = new match.route.view()
	// document.querySelector("#container").innerHTML = await view.renderHTML()
	const mainView = new MainPage()

	if (match.route.view === Auth) {
		let userId = localStorage.getItem("id")
		let cookie = getCookie("browserCookie")
		let expirationTime = new Date(cookie)
		let currentTime = new Date()

		console.log("cookie:", cookie)
		console.log("userId:", userId)

		if (!cookie && !userId) {
			document.cookie =
				"browserCookie=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/"
			const authView = new Auth()
			document.querySelector("#container").innerHTML =
				await authView.renderHTML()
			authView.submitForm()
			return
		}

		if (currentTime > expirationTime || !userId) {
			localStorage.clear()
			document.cookie =
				"browserCookie=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/"
			const authView = new Auth()
			document.querySelector("#container").innerHTML =
				await authView.renderHTML()
			authView.submitForm()
		} else {
			window.location.pathname = "/main"
		}
	}

	// Call the submitForm and displayPosts method here
	if (match.route.view === MainPage) {
		let userInfo = localStorage.getItem("id")

		if (!userInfo) {
			window.location.href = "/"
			return
		}
		document.querySelector("#container").innerHTML = await mainView.renderHTML()
		mainView.attachPostSubmitForm()
		mainView.displayUserContainer()
		mainView.displayPostContainer()
		mainView.displayChatContainer()
		mainView.Logout()
		mainView.reactions()
	}
}

window.addEventListener("popstate", router)

document.addEventListener("DOMContentLoaded", () => {
	document.body.addEventListener("click", (event) => {
		if (event.target.matches("[data-link]")) {
			event.preventDefault()
			navigateTo(event.target.href)
		}
	})
	router()
})

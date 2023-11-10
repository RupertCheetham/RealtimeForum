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
		let userInfo = localStorage.getItem("id")

		let currentTime = new Date()
		let expiration = new Date(currentTime)
		expiration.setMinutes(currentTime.getMinutes() + timeout)

		if (userInfo) {
			window.location.pathname = "/main"
		} else {
			if (currentTime > expiration) {
				localStorage.clear()
			}
			const authView = new Auth()
			document.querySelector("#container").innerHTML =
				await authView.renderHTML()
			authView.submitForm()
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
		mainView.runStartWebsocket()
		mainView.displayUserContainer()
		mainView.displayPostContainer()
		// mainView.displayChatContainer()
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

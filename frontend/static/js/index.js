import Auth from "./views/Auth.js"
import MainPage from "./views/MainPage.js"
import { getCookie } from "./utils/utils.js"

const navigateTo = (url) => {
	history.pushState(null, null, url)
	router()
}

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

	const view = new match.route.view()

	document.querySelector("#container").innerHTML = await view.renderHTML()

	if (match.route.view === Auth) {
		// document.querySelector("#container").innerHTML = await view.renderHTML()
		const authView = new Auth()
		authView.submitForm()
	}

	// Call the submitForm and displayPosts method here
	if (match.route.view === MainPage) {

		const mainView = new MainPage()
		mainView.attachPostSubmitForm()
		mainView.displayUserContainer()
		mainView.displayPostContainer()
		mainView.displayChatContainer()

		mainView.Logout()
		mainView.reactions()
	}

	console.log("match:", view)
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

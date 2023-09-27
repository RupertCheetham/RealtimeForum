import Auth from "./views/Auth.js"
import Comments from "./views/Comments.js"
import Posts from "./views/Posts.js"

const navigateTo = (url) => {
	history.pushState(null, null, url)
	router()
}

const router = async () => {
	const routes = [
		{ path: "/", view: Auth },
		{ path: "/posts", view: Posts },
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

	document.querySelector("#container").innerHTML = await view.getHTML()

	if (match.route.view === Auth) {
		const authView = new Auth()
		authView.submitForm()
	}

	// Call the submitForm method here
	if (match.route.view === Posts) {
		const postsView = new Posts()
		postsView.getPosts()
		postsView.submitForm()
	}

	// Call the submitForm method here
	if (match.route.view === Comments) {
		const postsView = new Posts()
		postsView.getPosts()
		postsView.submitCommentForm()
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
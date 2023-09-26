const router = async () => {
	const routes = [
		{ path: "/", view: () => console.log("viewing post and all") },
		{ path: "/auth", view: () => console.log("viewing authentication page") },
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

	console.log(match.route.view())
}

document.addEventListener("DOMContentLoaded", () => {
	router()
})

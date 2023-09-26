const router = async () => {
	const routes = [
		{ path: "/", view: () => console.log("viewing forum") },
		{ path: "/posts", view: () => console.log("viewing posts") },
	]

	// test each route for potential match
	const potentialMatches = routes.map((route) => {
		return {
			route: route,
			isMatch: location.pathname === route.path,
		}
	})

	console.log(potentialMatches)
}

document.addEventListener("DOMContentLoaded", () => {
	router()
})

import Auth from "./views/Auth.js"
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

	const postForm = document.getElementById("post-form")
	console.log("postForm:", postForm)
})

// const postForm = document.getElementById("post-form")
// console.log("postForm:", postForm)
// postForm.addEventListener("submit", function (event) {
// 	event.preventDefault()

// 	const postText = document.getElementById("postText").value
// 	const categories = document.getElementById("categories").value
// 	const image = document.getElementById("image").value

// 	console.log(postText, categories, image)

// 	fetch("http://localhost:8080/posts", {
// 		method: "POST",
// 		headers: {
// 			Accept: "application/json",
// 			"Content-Type": "application/json",
// 		},
// 		body: JSON.stringify({
// 			body: postText,
// 			categories: categories,
// 			img: image,
// 		}),
// 	})
// 		.then(async (response) => {
// 			if (response.ok) {
// 				// await viewPosts()
// 			}
// 		})
// 		.catch((error) => {
// 			console.log(error)
// 		})
// })

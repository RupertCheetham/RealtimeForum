import Posts from "./views/Posts.js"

import { setSessionCookie } from "./views/cookie.js"

async function switchToPostsView() {
	// Remove the current view content from the DOM
	const container = document.getElementById("container")

	// Create an instance of the Posts view and render it
	const postsView = new Posts()
	container.innerHTML = await postsView.getHTML()
}

function authenticate() {
	const authcontainer = document.getElementById("auth-container")
	const sign_in_btn = document.querySelector("#sign-in-btn")
	const sign_up_btn = document.querySelector("#sign-up-btn")

	sign_up_btn.addEventListener("click", () => {
		authcontainer.classList.add("sign-up-mode")
	})

	sign_in_btn.addEventListener("click", () => {
		authcontainer.classList.remove("sign-up-mode")
	})

	const signinForm = document.querySelector(".sign-in-form")
	signinForm.addEventListener("submit", function (event) {
		event.preventDefault()

		const userNameOrEmail = document.getElementById("usernameOrEmail").value
		const password = document.getElementById("password").value

		console.log(userNameOrEmail, password)

		fetch("http://localhost:8080/login", {
			method: "POST",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				username: userNameOrEmail,
				password: password,
			}),
		})
			.then((response) => {
				if (response.ok) {
					return response.json()
				} else {
					throw new Error("POST request failed!")
				}
			})
			.then(async (data) => {
				console.log("this is data", data)
				if (data.message === "Login successful") {
					setSessionCookie()
					switchToPostsView()
				}
			})
			.catch((error) => {
				console.log(error)
			})
	})

	const signupForm = document.querySelector(".sign-up-form")

	signupForm.addEventListener("submit", function (event) {
		event.preventDefault()

		const userName = document.getElementById("username").value
		const userAge = parseInt(document.getElementById("age").value)
		const userGender = document.getElementById("gender").value
		const firstName = document.getElementById("first_name").value
		const lastName = document.getElementById("last_name").value
		const email = document.getElementById("email").value
		const password = document.getElementById("new_password").value

		console.log(
			userName,
			userAge,
			userGender,
			firstName,
			lastName,
			email,
			password
		)

		fetch("http://localhost:8080/registrations", {
			method: "POST",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				username: userName,
				age: userAge,
				gender: userGender,
				first_name: firstName,
				last_name: lastName,
				email: email,
				password: password,
			}),
		}).catch((error) => {
			console.log(error)
		})
		console.log("registration complete")
	})
}

function makePost() {
	switchToPostsView()

	const postForm = document.getElementById("post-form")
	console.log("postForm:", postForm)
	postForm.addEventListener("submit", function (event) {
		event.preventDefault()
		const postText = document.getElementById("postText").value
		const categories = document.getElementById("categories").value
		const image = document.getElementById("image").value
		console.log(postText, categories, image)
		fetch("http://localhost:8080/posts", {
			method: "POST",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				body: postText,
				categories: categories,
				img: image,
			}),
		})
			.then(async (response) => {
				if (response.ok) {
					// await viewPosts()
					switchToPostsView()
				}
			})
			.catch((error) => {
				console.log(error)
			})
	})
}

async function getPosts() {
	const response = await fetch("http://localhost:8080/posts")
	const postContainer = document.getElementById("postContainer")
	const posts = await response.json()

	for (const post of posts) {
		const postElement = document.createElement("div")
		postElement.id = "Post" + post.id
		postElement.classList.add("post")

		const comments = await fetchComments(post.id) // Wait for the comments to be fetched

		postElement.textContent = `
      Id: ${post.id},
      Username: ${post.username},
      Img: ${post.img},
      Body: ${post.body},
      Categories: ${post.categories},
      Reaction: ${post.reaction},
    `

		let commentHTML = `
	<form id="comment-form">
	<div>
	<label for="commentText"><b>Comment</b></label>
	<input type="text" placeholder="Enter comment" name="commentText" id="commentText" required /><br>
	<button type="submit" id="submit">Submit Comment</button>
  	</div>
	</form>
	`

		if (comments.length > 0) {
			const commentsContainer = document.createElement("div")
			commentsContainer.id = "commentContainer"
			let commentsNum = 1
			comments.forEach((comment) => {
				const commentElement = document.createElement("div")
				commentElement.className = "comment" + commentsNum++
				commentElement.textContent = `Comment: ${comment.body}`
				commentsContainer.appendChild(commentElement)
			})

			postElement.appendChild(commentsContainer)
			console.log(postElement)
		}
		postElement.innerHTML += commentHTML
		postContainer.appendChild(postElement)
	}

	// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
	async function fetchComments(parentPostID) {
		const response = await fetch("http://localhost:8080/comments")
		const comments = await response.json()
		console.log("comments:", comments)
		return comments.filter((comment) => comment.parentPostId == parentPostID)
	}
}

document.addEventListener("DOMContentLoaded", async () => {
	// Your code here
	authenticate()
	makePost()
	await getPosts()
})

import { navbar } from "./nav.js"
import { setSessionCookie } from "./cookie.js"
import { showRegistrationForm } from "./registration.js"
import { viewPosts } from "./postsView.js"

const container = document.getElementById("container")

export function showLoginForm() {
	navbar()
	let html = `
		<h1>Login to Forum</h1>
		<form id="login-form">
			<div>
				<label for="username"><b>Username</b></label>
				<input
					type="text"
					placeholder="Enter Username"
					name="username"
					required
					id="usernameOrEmail"
				/>
			</div>
			<div>
				<label for="password"><b>Password</b></label>
				<input
					type="password"
					placeholder="Enter Password"
					name="psw"
					required
					id="password"
					/>
			</div>
			<div>
				<button type="submit">Login</button>
				<button type="button" id="signup">Sign Up</button>
			</div>
		</form>
 `
	container.innerHTML = html

	// Add event listener to the "Register Account" link
	const signup = document.getElementById("signup")
	signup.addEventListener("click", function (event) {
		event.preventDefault()
		showRegistrationForm()
	})

	const loginForm = document.getElementById("login-form")
	// Event listener for submitting to login form
	loginForm.addEventListener("submit", function (event) {
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
			.then((data) => {
				console.log(data.message)
				if (data.message === "Login successful") {
					setSessionCookie()
					viewPosts()
				}
			})
			.catch((error) => {
				console.log(error)
			})
	})
}

showLoginForm()

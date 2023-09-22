import { setSessionCookie } from "./cookie.js"
import { showRegistrationForm } from "./registration.js"

const container = document.getElementById("container")

export function showLoginForm() {
	let html = `
		<form id="login-form">
			<div>
				<h1>Login to Forum</h1>
			</div>
			<div class="container">
				<label for="username"><b>Username</b></label>
				<input
					type="text"
					placeholder="Enter Username"
					name="username"
					required
					id="username"
				/>
				<br />
				<p>
					<label for="password"><b>Password</b></label>
					<input
						type="password"
						placeholder="Enter Password"
						name="psw"
						required
						id="password"
					/>
					<br />
					<button type="submit">Login</button>
				</p>
			</div>
		</form>
		<span>
			<a href="#" id="registration-form">Register Account</a>
		</span>
	</div>
 `
	container.innerHTML = html
}

// showLoginForm()

// Add event listener to the "Register Account" link
const registrationForm = document.getElementById("registration-form")
registrationForm.addEventListener("click", function (event) {
	event.preventDefault()
	showRegistrationForm()
})

const loginForm = document.getElementById("login-form")

// Event listener for switching to the login form
loginForm.addEventListener("click", function (event) {
	event.preventDefault()
	showLoginForm()
})

// Event listener for submitting to login form
loginForm.addEventListener("submit", function (event) {
	event.preventDefault()

	const userName = document.getElementById("username").value
	const password = document.getElementById("password").value

	console.log(userName, password)

	fetch("http://localhost:8080/login", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			username: userName,
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
			console.log(data)
			setSessionCookie()
		})
		.catch((error) => {
			console.log(error)
		})
})

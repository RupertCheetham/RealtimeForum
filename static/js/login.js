import { setSessionCookie } from "./cookie.js"
// import { loadLoginScript } from "./registrationToBackend.js"

const container = document.getElementById("container")

export function showLoginForm() {
	let html = `
		<form id="login-form">
			<div>
				<h1>Login to Forum</h1>
			</div>
			<div class="container">
				<label for="usernameOrEmail"><b>Username or Email</b></label>
				<input
					type="text"
					placeholder="Enter Username or Email"
					name="username"
					required
					id="usernameOrEmail"
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

			<div class="container" style="background-color: #f1f1f1">
				<button type="button" class="signup" id="signup">Sign Up</button>
				<span class="psw">
					Forgotten
					<a href="#">password?</a>
				</span>
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
})

const loginForm = document.getElementById("login-form")

// Event listener for switching to the login form
loginForm.addEventListener("click", function (event) {
	event.preventDefault()
})

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
			}
		})
		.catch((error) => {
			console.log(error)
		})
})

const signupButton = document.getElementById("signup")

signupButton.addEventListener("click", function (event) {
	event.preventDefault()

	// Load or serve your registrationToBackend.js script here
	loadRegistrationToBackendScript()
})

function loadRegistrationToBackendScript() {
	// Create a script element
	const script = document.createElement("script")

	// Set the src attribute to your registrationToBackend.js file
	script.src = "../static/js/registrationToBackend.js"

	// Append the script element to the document's head
	document.head.appendChild(script)
}

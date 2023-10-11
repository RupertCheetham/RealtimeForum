import AbstractView from "./AbstractView.js"
import Posts from "./Posts.js"
import { setSessionCookie } from "../utils/cookie.js"

async function switchToPostsView() {
	const container = document.getElementById("container")

	// Create an instance of the Posts view and render it
	const postsView = new Posts()
	container.innerHTML = await postsView.getHTML()
}

export default class extends AbstractView {
	constructor() {
		super()
		this.setTitle("Sign in  or sign up")
	}

	async getHTML() {
		return `
		<div id="auth-container" class="auth-container">
		<div class="forms-container">
			<div class="signin-signup">
				<div class="input-field-container">
					<form id="sign-up-form" action="#" class="sign-in-form">
						<h2 class="title">Sign in</h2>
						<div class="input-field">
							<i class="fas fa-user"></i>
							<input
								type="text"
								placeholder="Username or Email"
								required
								id="usernameOrEmail"
							/>
						</div>
						<div class="input-field">
							<i class="fas fa-lock"></i>
							<input
								type="password"
								placeholder="Password"
								required
								id="password"
							/>
						</div>

						<input type="submit" value="Login" class="btn solid" />

						<p class="social-text">Or Sign in with social platforms</p>
						<div class="social-media">
							<a href="#" class="social-icon">
								<i class="fab fa-facebook-f"></i>
							</a>
							<a href="#" class="social-icon">
								<i class="fab fa-twitter"></i>
							</a>
							<a href="#" class="social-icon">
								<i class="fab fa-google"></i>
							</a>
							<a href="#" class="social-icon">
								<i class="fab fa-linkedin-in"></i>
							</a>
						</div>
					</form>

					<form class="sign-up-form">
						<h2 class="title">Sign up</h2>
						<div =class"progress green">
							<div =class"indeterminate"></div>
						</div>
						
						<div class="input-field">
							<i class="fas fa-user"></i>
							<input
								type="text"
								placeholder="Username"
								required
								id="username"
							/>
						</div>
						<div class="input-field">
							<i class="fas fa-user"></i>
							<input type="text" placeholder="User Age" required id="age" />
						</div>
						<div class="input-field">
							<i class="fas fa-user"></i>
							<input
								type="text"
								placeholder="First Name"
								required
								id="first_name"
							/>
						</div>
						<div class="input-field">
							<i class="fas fa-user"></i>
							<input
								type="text"
								placeholder="Last Name"
								required
								id="last_name"
							/>
						</div>
						<div class="input-field">
							<i class="fas fa-envelope"></i>
							<input type="email" placeholder="Email" required id="email" />
						</div>
						<div class="input-field">
							<i class="fas fa-lock"></i>
							<input
								type="password"
								placeholder="Password"
								required
								id="new_password"
							/>
						</div>
						<div class="input-field">
							<i class="fas fa-lock"></i>
							<input
								type="password"
								placeholder="Repeat Password"
								required
								id="password-repeat"
							/>
						</div>
						<div class="input-field">
							<i class="fas fa-user"></i>
							<input type="text" placeholder="Gender" required id="gender" />
						</div>

						<input type="submit" value="Sign up" class="btn" />

						<p class="social-text">Or Sign up with social platforms</p>
						<div class="social-media">
							<a href="#" class="social-icon">
								<i class="fab fa-facebook-f"></i>
							</a>
							<a href="#" class="social-icon">
								<i class="fab fa-twitter"></i>
							</a>
							<a href="#" class="social-icon">
								<i class="fab fa-google"></i>
							</a>
							<a href="#" class="social-icon">
								<i class="fab fa-linkedin-in"></i>
							</a>
						</div>
					</form>
				</div>
			</div>
		</div>

		<div class="panels-container">
			<div class="panel left-panel">
				<div class="content">
					<h3>New here ?</h3>
					<p>
						Become a member and let's get you started!
					</p>
					<button class="btn transparent" id="sign-up-btn">Sign up</button>
				</div>
				<img src="img/log.svg" class="image" alt="" />
			</div>
			<div class="panel right-panel">
				<div class="content">
					<h3>One of us ?</h3>
					<p>
						Check out the latest posts and share your thoughts and news.
					</p>
					<button class="btn transparent" id="sign-in-btn">Sign in</button>
				</div>
				<img src="img/register.svg" class="image" alt="" />
			</div>
		</div>
		</div>
</div>
 `
	}

	async submitForm() {
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
						window.location.pathname = "/posts"
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
}

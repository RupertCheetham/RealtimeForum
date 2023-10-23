import AbstractView from "./AbstractView.js"
import { getCookie } from "../utils/utils.js"

export default class Auth extends AbstractView {
	constructor() {
		super()
		this.setTitle("Sign in or sign up")
	}

	async renderHTML() {
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
                <!-- Step 1: User Information -->
                <form id="step1-form" action="#" class="sign-up-form">
                    <h2 class="title">Step 1: User Information</h2>
                    <div class="input-field">
                        <i class="fas fa-user"></i>
                        <input type="text" placeholder="Username" required id="username" />
                    </div>
                   
                    <div class="input-field">
                        <i class="fas fa-user"></i>
                        <input type="text" placeholder="First Name" required id="first_name" />
                    </div>
                    <div class="input-field">
                        <i class="fas fa-user"></i>
                        <input type="text" placeholder="Last Name" required id="last_name" />
                    </div>
					<div class="input-field">
                        <i class="fas fa-envelope"></i>
                        <input type="email" placeholder="Email" required id="email" />
                    </div>

                    <button class="btn transparent" id="next-step2-btn">Next</button>
					
                </form>

                <!-- Step 2: Account Information -->
                <form id="step2-form" action="#" class="sign-up-form">
                    <h2 class="title">Step 2: Account Information</h2>
                    
                    <div class="input-field">
                        <i class="fas fa-lock"></i>
                        <input type="password" placeholder="Password" required id="new_password" />
                    </div>
                    <div class="input-field">
                        <i class="fas fa-lock"></i>
                        <input type="password" placeholder="Repeat Password" required id="password-repeat" />
                    </div>
					<div class="input-field">
						<i class="fas fa-calendar"></i> <!-- Use a calendar icon for D.o.B -->
						<input type="dete" placeholder="Date of Birth" required id="date_of_birth" />
					</div>
					<div class="input-field">
                        <i class="fas fa-user"></i>
                        <input type="text" placeholder="Gender" required id="gender" />
                    </div>
					<div class="button-container">
					<button class="btn transparent" id="back-step1-btn">Back</button>
                    <input class="btn" type="submit" value="Sign up" id="submit-btn" />

                </form>
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

<script>

</script>
 `
	}

	async submitForm() {
		const authcontainer = document.getElementById("auth-container")
		const sign_in_btn = document.querySelector("#sign-in-btn")
		const sign_up_btn = document.querySelector("#sign-up-btn")
		const signupForm = document.querySelector(".sign-up-form")

		sign_up_btn.addEventListener("click", () => {
			authcontainer.classList.add("sign-up-mode")
		})

		sign_in_btn.addEventListener("click", () => {
			authcontainer.classList.remove("sign-up-mode")
		})

		const signinForm = document.querySelector(".sign-in-form")
		signinForm.addEventListener("submit", async function (event) {
			event.preventDefault()

			const userNameOrEmail = document.getElementById("usernameOrEmail").value
			const password = document.getElementById("password").value

			console.log(userNameOrEmail, password)
			try {
				const response = await fetch("https://localhost:8080/api/auth", {
					method: "POST",
					headers: {
						Accept: "application/json",
						"Content-Type": "application/json",
					},
					body: JSON.stringify({
						username: userNameOrEmail,
						password: password,
					}),
					credentials: "include", // Ensure cookies are included in the request
				})

				if (response.ok) {
					// Authentication successful, redirect to protected page
					console.log("cookie in auth is:", document.cookie)
					let cookie = getCookie("sessionID")
					if (!cookie) {
						window.location.href = "/"
					} else {
						window.location.href = "posts" // Update the URL
					}
				} else {
					throw new Error("Authentication failed!")
				}
			} catch (error) {
				console.error(error)
			}
		})

		signupForm.addEventListener("submit", function (event) {
			event.preventDefault()

			const userName = document.getElementById("username").value
			const userAge = parseInt(document.getElementById("age").value)
			const userGender = document.getElementById("gender").value
			const firstName = document.getElementById("first_name").value
			const lastName = document.getElementById("last_name").value
			const email = document.getElementById("email").value
			const password = document.getElementById("new_password").value

			const stepOne = document.getElementById("step1-form")
			const stepTwo = document.getElementById("step2-form")

			const next1 = document.getElementById("next-step2-btn")
			const back1 = document.getElementById("back-step1-btn")

			next1.onclick = function(){
				stepOne.style.left = "-900px"
				stepTwo.style.left = "900px"	
			}

			back1.onclick = function(){
				stepOne.style.right = "900px"
				stepTwo.style.right = "900px"	
			}


			console.log(userName, userAge)

			fetch("https://localhost:8080/api/registrations", {
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
			},
			console.log("I'm in the post registrations JS"))
				.then((response) => {
					if (response.ok) {
						const userError = document.querySelector(".username-error")
						userError.style.display = "none"
					}
					if (response.status === 400) {
						const userError = document.querySelector(".username-error")
						userError.style.display = "block"

						setTimeout(() => {
							userError.style.display = "none"
						}, 4000)

						throw new Error("Unable to create user")
					}
				})
				.catch((error) => {
					console.log(error)
				})
		})
	}
}

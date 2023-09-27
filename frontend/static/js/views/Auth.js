import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
	constructor() {
		super()
		this.setTitle("Sign in  or sign up")
	}

	async getHTML() {
		return `
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

					<form action="#" class="sign-up-form">
						<h2 class="title">Sign up</h2>
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
							<input type="text" placeholder="Gender" required id="gender" />
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
						Lorem ipsum, dolor sit amet consectetur adipisicing elit. Debitis,
						ex ratione. Aliquid!
					</p>
					<button class="btn transparent" id="sign-up-btn">Sign up</button>
				</div>
				<img src="img/log.svg" class="image" alt="" />
			</div>
			<div class="panel right-panel">
				<div class="content">
					<h3>One of us ?</h3>
					<p>
						Lorem ipsum dolor sit amet consectetur adipisicing elit. Nostrum
						laboriosam ad deleniti.
					</p>
					<button class="btn transparent" id="sign-in-btn">Sign in</button>
				</div>
				<img src="img/register.svg" class="image" alt="" />
			</div>
		</div>
 `
	}
}

const container = document.getElementById("container")

export function showRegistrationForm() {
	let html = `
	<form id="registration-form">
		<div>
			<h1>Register</h1>
			<p>Kindly fill in this form to register.</p>
			<label for="username">
				<b>Username</b>
			</label>
			<input
				type="text"
				placeholder="Enter Username"
				name="username"
				id="username"
				required
			/>
			<br />
			<label for="age">
				<b>Age</b>
			</label>
			<input type="text" placeholder="Enter Age" name="age" id="age" required />
			<br />
			<label for="gender">
				<b>Gender</b>
			</label>
			<input
				type="text"
				placeholder="Enter Gender"
				name="gender"
				id="gender"
				required
			/>
			<br />
			<label for="first_name">
				<b>First Name</b>
			</label>
			<input
				type="text"
				placeholder="Enter First Name"
				name="first_name"
				id="first_name"
				required
			/>
			<br />
			<label for="last_name">
				<b>Last Name</b>
			</label>
			<input
				type="text"
				placeholder="Enter Last Name"
				name="last_name"
				id="last_name"
				required
			/>
			<br />
			<label for="email">
				<b>Email</b>
			</label>
			<input
				type="text"
				placeholder="Enter Email"
				name="email"
				id="email"
				required
			/>
			<br />
			<label for="password">
				<b>Password</b>
			</label>
			<input
				type="text"
				placeholder="Enter Password"
				name="password"
				id="password"
				required
			/>
			<br />
			<label for="password-repeat">
				<b>Repeat Password</b>
			</label>
			<input
				type="password"
				placeholder="Repeat Password"
				name="password-repeat"
				id="password-repeat"
				required
			/>
			<br />
			<button type="submit" id="submit">
				Register
			</button>
		</div>
		<div>
			<p>
				Already have an account?
				<a href="#" id="login-form">Log in</a>
			</p>
		</div>
	</form>
`
	container.innerHTML = html
}

const registrationForm = document.getElementById("registration-form")
// Event listener for switching to the registration form
registrationForm.addEventListener("click", function (event) {
	event.preventDefault()
	showRegistrationForm()
})

registrationForm.addEventListener("submit", function (event) {
	event.preventDefault()
	const userName = document.getElementById("username").value
	const userAge = parseInt(document.getElementById("age").value)
	const userGender = document.getElementById("gender").value
	const firstName = document.getElementById("first_name").value
	const lastName = document.getElementById("last_name").value
	const email = document.getElementById("email").value
	const password = document.getElementById("password").value

	console.log(userName, userAge)

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
})

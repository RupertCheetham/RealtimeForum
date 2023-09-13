const registrationForm = document.getElementById("registration-form")

registrationForm.addEventListener("submit", function (event) {
	event.preventDefault()

	const userName = document.getElementById("username").value
	const userAge = parseInt(document.getElementById("age").value, 10)
	const userGender = document.getElementById("gender").value
	const firstName = document.getElementById("first_name").value
	const lastName = document.getElementById("last_name").value
	const email = document.getElementById("email").value
	const password = document.getElementById("password").value

	console.log("userName:", userName)
	console.log("userAge:", userAge)

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

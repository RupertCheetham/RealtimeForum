import { setSessionCookie } from './cookie.js';

const loginForm = document.getElementById("login-form")

loginForm.addEventListener("submit", function (event) {
	event.preventDefault()

	//console.log("you are logged in?")

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
		})
	})
		.then(response => {
			if (response.ok) {
				return response.json()
			} else {
				throw new Error("POST request failed!")
			}
		})
		.then(data => {
			console.log("data:", data)
			setSessionCookie();
		})
		.catch(error => {
			console.log(error)
		})
})
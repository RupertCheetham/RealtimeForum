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
		.then((response) =>{
			if (response === 201 || response === 200){
				//setSessionCookie()
				return response.json()
			}	
		})
		.then((data) => {

			console.log(data)

		// 	// if (data.success) {
		// 		// Authentication successful, set session cookie and redirect
	 		setSessionCookie();
		// 	// } else {
		// 		// Authentication failed, display an error message
		// 		// alert("Authentication failed. Please check your username and password.");
		// 	// }
		})
		// .catch((error) => {
		// 	console.log(error)
		// })
})
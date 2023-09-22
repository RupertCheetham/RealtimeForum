/* The code is adding an event listener to the registration form. When the form is submitted, the
function will be executed. */

const container = document.getElementById("container")
let html = `
<form id="registration-form">
<div>
	<h1>Register</h1>
</div>
	<div class="container">
<p>Kindly fill in this form to register.</p>
 <!-- label and input for Username -->
 <label for="username"><b>Username</b></label> 
 <input
   type="text"
   placeholder="Enter Username"
   name="username"
   id="username"
   required
 /><br>
   <!-- label and input for Age -->
   <label for="age"><b>Age</b></label> 
   <input
	 type="text"
	 placeholder="Enter Age"
	 name="age"
	 id="age"
	 required
   /><br>
	 <!-- label and input for gender -->
 <label for="gender"><b>Gender</b></label> 
 <input
   type="text"
   placeholder="Enter Gender"
   name="gender"
   id="gender"
   required
 /><br>
  <!-- label and input for first name -->
<label for="first_name"><b>First Name</b></label>
<input
  type="text"
  placeholder="Enter First Name"
  name="first_name"
  id="first_name"
  required
/><br>
 <!-- label and input for last name -->
 <label for="last_name"><b>Last Name</b></label>
 <input
   type="text"
   placeholder="Enter Last Name"
   name="last_name"
   id="last_name"
   required
 /><br>
  <!-- label and input for email -->
  <label for="email"><b>Email</b></label>
  <input
	type="text"
	placeholder="Enter Email"
	name="email"
	id="email"
	required
  /><br>
	<!-- label and input for password -->
	<label for="password"><b>Password</b></label>
	<input
	  type="text"
	  placeholder="Enter Password"
	  name="password"
	  id="password"
	  required
	/><br>
	<!-- label and input for repeat password -->
	<label for="password-repeat"><b>Repeat Password</b></label>
	<input
	  type="password"
	  placeholder="Repeat Password"
	  name="password-repeat"
	  id="password-repeat"
	  required
	/><br>


	 <!-- submit button -->
	 
<button type="submit" id="submit">Register</button>
</div>
 <!-- wrapping the text inside the p tag with a tag for routing to the login page URL-->
<!-- the # must ideally be replaced by the login page URL -->
<div>
<div>
<p>Already have an account? <a href="#">Log in</a>.</p>

<button type="button" class="login-button" id="login-button">Log In</button>

</div>
</form>
 `
container.innerHTML = html;

const registrationForm = document.getElementById("registration-form")

registrationForm.addEventListener("submit", function (event) {
	event.preventDefault()

	const userName = document.getElementById("username").value
	const userAge = parseInt(document.getElementById("age").value)
	const userGender = document.getElementById("gender").value
	const firstName = document.getElementById("first_name").value
	const lastName = document.getElementById("last_name").value
	const email = document.getElementById("email").value
	const password = document.getElementById("password").value

	console.log(userName, userAge, userGender, firstName, lastName, email, password)

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

// const loginButton = document.getElementById("login-button");

// loginButton.addEventListener("click", function (event) {
//     event.preventDefault();
    
//     // Load or serve your login.js script here
//     loadLoginScript();
// });

// function loadLoginScript() {
//     // Create a script element
//     const script = document.createElement("script");
    
//     // Set the src attribute to your login.js file
//     script.src = "../static/js/login.js";
    
//     // Append the script element to the document's head
//     document.head.appendChild(script);
// }

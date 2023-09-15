const loginForm = document.getElementById("login-form")

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
            if (response.status === 200 || response.status === 201) {
                return response.json(); // Only parse JSON if the status code is 200 or 201
            } else {
                throw new Error("HTTP status code: " + response.status);
            }
        })
        .then((data) => {
            console.log(data);
            if (data.success) {
                // Authentication successful, set session cookie and redirect
                // setSessionCookie();
                console.log("success");
            } else {
                // Authentication failed, display an error message
                alert("Authentication failed. Please check your username and password.");
            }
        })
        .catch((error) => {
            console.log(error);
            // Handle other types of responses or errors here
        });
})

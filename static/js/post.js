const postForm = document.getElementById("post-form")

postForm.addEventListener("submit", function (event) {
	event.preventDefault()
	console.log("clicked post form")

	const postimage = document.getElementById("postimage")
	// const postbody = document.getElementById("postbody").value
	// console.log(postbody)
	console.log(postimage)
	// const categories = document.getElementById("first_name").value
	// const creationDate = document.getElementById("last_name").value
	// const likes = document.getElementById("email").value
	// const dislikes = document.getElementById("password").value
	// const WhoLiked = document.getElementById("password").value
	// const whoDisliked = document.getElementById("password").value

	fetch("http://localhost:8080/postupload", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			// username: userName,
			img: postimage,
			// body: postbody,
			// categories: firstName,
			// likes: lastName,
			// dislikes: email,
			// whoLiked: lastName,
			// whoDisliked: email,
			// password: password,
		}),
	}).catch((error) => {
		console.log(error)
	})
})

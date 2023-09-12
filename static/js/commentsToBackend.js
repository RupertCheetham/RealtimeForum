const commentForm = document.getElementById("comment-form")

commentForm.addEventListener("submit", function (event) {
	event.preventDefault()

	const comment = document.getElementById("commentText").value

	console.log(comment)

	fetch("http://localhost:8080/comments", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			Body: comment,
		}),
	}).catch((error) => {
		console.log(error)
	})
	console.log("bottom of page")
})

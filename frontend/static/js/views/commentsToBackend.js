const commentForm = document.getElementById("comment-form")

commentForm.addEventListener("submit", function (event) {
	event.preventDefault()

	const username = document.getElementById("commentUsername").value
	const parentPostID = parseInt(
		document.getElementById("parentPostID").value,
		10
	)
	const comment = document.getElementById("commentText").value
	console.log(username)
	console.log(parentPostID)
	console.log(comment)

	fetch("http://localhost:8080/comments", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			Username: username,
			ParentPostID: parentPostID,
			Body: comment,
		}),
	}).catch((error) => {
		console.log(error)
	})
})

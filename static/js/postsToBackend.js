/* The code is adding an event listener to the form with the id "post-form". When the form is
submitted, the event listener function is executed. */
const postForm = document.getElementById("post-form")

postForm.addEventListener("submit", function (event) {
	event.preventDefault()

	const postText = document.getElementById("postText").value
	const categories = document.getElementById("categories").value
    const image = document.getElementById("image").value

	console.log(postText, categories, image)

	fetch("http://localhost:8080/posts", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			Body: postText,
			Categories: categories,
			Img: image,
		}),
	}).catch((error) => {
		console.log(error)
	})
	console.log("bottom of page")
})

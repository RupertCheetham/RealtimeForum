/* The code is adding an event listener to the form with the id "post-form". When the form is
submitted, the event listener function is executed. */

let html = `

<form id="post-form">
    <div>
      <h1>Post</h1>
      <p>Kindly fill in this form to post.</p>
      <!-- label and input for postText -->
      <label for="post"><b>Post</b></label>
      <input type="text" placeholder="Enter Message" name="postText" id="postText" required /><br>
      <!-- label and input for categories -->
      <label for="categories"><b>Categories</b></label>
      <input type="text" placeholder="Enter Categories" name="categories" id="categories" required /><br>
      <!-- label and input for img -->
      <label for="image"><b>Image</b></label>
      <input type="text" placeholder="Enter Image String" name="image" id="image" required /><br>
      <!-- submit button -->

      <button type="submit" id="submit">Submit Post</button>
    </div>
    <!-- wrapping the text inside the p tag with a tag for routing to the login page URL-->
    <!-- the # must ideally be replaced by the login page URL -->
    <div>

  </form>

 `
container.innerHTML = html

function postsToBackend() {

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
}
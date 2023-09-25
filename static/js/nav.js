const navContainer = document.getElementById("nav")

export function navbar() {
	let html = `
  <div>Welcome to forum</div>
  <form id="post-form">
  <div>
  <p>Kindly fill in this form to post.</p>
  <label for="post"><b>Post</b></label>
  <input type="text" placeholder="Enter Message" name="postText" id="postText" required /><br>
  <label for="categories"><b>Categories</b></label>
  <input type="text" placeholder="Enter Categories" name="categories" id="categories" required /><br>
  <label for="image"><b>Image</b></label>
  <input type="text" placeholder="Enter Image String" name="image" id="image" required /><br>
  <button type="submit" id="submit">Submit Post</button>
  </div>
  </form>
  `
	navContainer.innerHTML = html

	// make a post
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
		})
			.then((response) => {
				if (response.ok) {
				}
			})
			.catch((error) => {
				console.log(error)
			})
	})
}

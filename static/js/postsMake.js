import { viewPosts } from "./postsView.js"

const container = document.getElementById("container")

export function makePosts() {
	let html = `
    <h2>Post</h2>
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
    <form id="comment-form">
    <h2>Comment</h2>
    <div>
      <p>Kindly fill in this form to comment.</p>
      <label for="username"><b>Username</b></label>
      <input type="text" placeholder="Enter Username" name="commentUsername" id="commentUsername" required /><br>
      <label for="parentPostID"><b>parentPostID</b></label>
      <input type="text" placeholder="Enter parentPostID" name="parentPostID" id="parentPostID" required /><br>
      <label for="commentText"><b>Comment</b></label>
      <input type="text" placeholder="Enter comment" name="commentText" id="commentText" required /><br>
      <button type="submit" id="submit">Submit Comment</button>
    </div>
    </form>
  `
	container.innerHTML = html

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
					viewPosts()
				}
			})
			.catch((error) => {
				console.log(error)
			})
	})
}

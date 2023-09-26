import { navbar } from "./nav.js"

const container = document.getElementById("container")

// fetch posts
export async function viewPosts() {
	navbar()
	let html = `
		<div>
			<div id="postContainer"></div>
		</div>
	`

	container.innerHTML = html

	container.classList.remove("container")

	const response = await fetch("http://localhost:8080/posts")
	const postContainer = document.getElementById("postContainer")
	const posts = await response.json()

	for (const post of posts) {
		const postElement = document.createElement("div")
		postElement.id = "Post" + post.id
		postElement.classList.add("post")

		const comments = await fetchComments(post.id) // Wait for the comments to be fetched

		postElement.textContent = `
      Id: ${post.id},
      Username: ${post.username},
      Img: ${post.img},
      Body: ${post.body},
      Categories: ${post.categories},
      Reaction: ${post.reaction},
    `

		let commentHTML = `
	<form id="comment-form">
	<div>
	<label for="commentText"><b>Comment</b></label>
	<input type="text" placeholder="Enter comment" name="commentText" id="commentText" required /><br>
	<button type="submit" id="submit">Submit Comment</button>
  	</div>
	</form>
	`

		if (comments.length > 0) {
			const commentsContainer = document.createElement("div")
			commentsContainer.id = "commentContainer"
			let commentsNum = 1
			comments.forEach((comment) => {
				const commentElement = document.createElement("div")
				commentElement.className = "comment" + commentsNum++
				commentElement.textContent = `Comment: ${comment.body}`
				commentsContainer.appendChild(commentElement)
			})

			postElement.appendChild(commentsContainer)
			console.log(postElement)
		}
		postElement.innerHTML += commentHTML
		postContainer.appendChild(postElement)
	}
}

// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
async function fetchComments(parentPostID) {
	const response = await fetch("http://localhost:8080/comments")
	const comments = await response.json()
	console.log("comments:", comments)
	return comments.filter((comment) => comment.parentPostId == parentPostID)
}

import { makePosts } from "./postsMake.js"

const container = document.getElementById("container")

// fetch posts
export async function viewPosts() {
	let html = `
		<div>
			<button type="button" id="makePost">Make a new post</button>
			<div id="postContainer"></div>
		</div>
	`

	container.innerHTML = html

	const makePost = document.getElementById("makePost")
	makePost.addEventListener("click", (event) => {
		event.preventDefault()
		makePosts()
	})

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

import AbstractView from "./AbstractView.js"
import {
	fetchComments,
	attachCommentForm,
	attachCommentsToPost,
} from "./Comments.js"

export default class Posts extends AbstractView {
	async renderHTML() {
		const postContainer = document.getElementById("postContainer")
		postContainer.innerHTML = ""

		const response = await fetch("https://localhost:8080/api/getposts", {
			// checkSessionTimeout(response)
			credentials: "include", // Ensure cookies are included in the request
		})

		if (response.status == 408) {
			window.location.href = "/"
		}
		// checkSessionTimeout(response)

		const posts = await response.json()
		const username = localStorage.getItem("username")

		for (const post of posts) {
			let postBox = document.createElement("div")
			postBox.id = "PostBox" + post.id
			postBox.classList.add("postBox")
			let postElement = document.createElement("div")
			postElement.id = "Post" + post.id
			postElement.classList.add("post")
			postElement.setAttribute("reactionID", post.reactionID)

			postElement.innerHTML = `
			<ul>
			  <li><b>Username:</b> ${username}</li>
			  <li><b>Img:</b> ${post.img}</li>
			  <li><b>Body:</b> ${post.body}</li>
			  <li><b>Categories:</b> ${post.categories}</li>
			  <button class="reaction-button" reaction-parent-class="post" reaction-parent-id="${
					post.id
				}" reaction-action="like" reaction-id = "${postElement.getAttribute(
				"reactionID"
			)}">👍 ${post.postLikes}</button>
			  <button class="reaction-button" reaction-parent-class="post" reaction-parent-id="${
					post.id
				}" reaction-action="dislike" reaction-id = "${postElement.getAttribute(
				"reactionID"
			)}">👎 ${post.postDislikes}</button>
			  </li>
			</ul>
		  `

			// attaches the comment form to the bottom of each post
			attachCommentForm(post, postElement)

			// fetch comments, if any, for this post
			postBox.appendChild(postElement)
			let comments = await fetchComments(post.id) // Wait for the comments to be fetched

			if (comments !== null) {
				let postComments = attachCommentsToPost(comments)
				postBox.appendChild(postComments)
				postComments.style.display = "none"
				postElement.addEventListener("click", () => {
					if (postComments.style.display === "none") {
						postComments.style.display = "block"
					}
				})
			}
			postContainer.appendChild(postBox)
		}
	}
}

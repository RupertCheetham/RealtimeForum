import AbstractView from "./AbstractView.js"
import {
	attachCommentForm,
	attachCommentsToPost,
	createCloseCommentButton,
} from "./Comments.js"

export default class Posts extends AbstractView {
	async renderHTML() {
		const response = await fetch("https://localhost:8080/api/getposts", {
			credentials: "include",
		})
		if (response.status == 408) {
			localStorage.clear()
			window.location.href = "/"
		}

		const posts = await response.json()
		const postsContainer = document.getElementById("postsContainer")
		postsContainer.innerHTML = ""

		for (const post of posts) {
			this.processPost(postsContainer, post)
		}
		
	}

	async processPost(postsContainer, post) {
		//makes the container for the individual posts and all its contents
		let postContainer = document.createElement("div")
		postContainer.id = "postContainer" + post.id
		postContainer.classList.add("postContainer")
		let postElement = document.createElement("div")
		postElement.id = "Post" + post.id
		postElement.classList.add("post")
		postElement.setAttribute("reactionID", post.reactionID)

		postElement.innerHTML = `
		<ul>
  			<li><b>Username:</b> ${post.username}</li>
  			<li><b>Img:</b> ${post.img}</li>
  			<li><b>Body:</b> ${post.body}</li>
  			<li><b>Categories:</b> ${post.categories}</li>
 				<button class="reaction-button" reaction-parent-class="post"
					reaction-parent-id="${post.id}" reaction-action="like" reaction-id = "${postElement.getAttribute("reactionID")}">👍 ${post.postLikes}</button>
  				<button class="reaction-button" reaction-parent-class="post"
					reaction-parent-id="${post.id}" reaction-action="dislike" reaction-id = "${postElement.getAttribute("reactionID")}">👎 ${post.postDislikes}</button>
 			</li>
		</ul>
`

		// attaches the comment form to the bottom of each post
		attachCommentForm(post, postElement)

		// appends the post(element) to the post Box
		postContainer.appendChild(postElement)

		// makes the commentsContainer, that comments (if any) will be appended to
		const commentsContainer = document.createElement("div")
		commentsContainer.id = "commentsContainer" + post.id
		commentsContainer.className = "commentsContainer"
		postContainer.appendChild(commentsContainer)
		commentsContainer.style.display = "none"
		const closeCommentButton = createCloseCommentButton(commentsContainer)
		// attaches closeComment Button to bottom of commentscontainer
		postContainer.appendChild(closeCommentButton)
		closeCommentButton.style.display = "none"

		postElement.addEventListener("click", () => {
			if (commentsContainer.style.display === "none") {
				commentsContainer.style.display = "block"
				// if (commentsContainer.querySelectorAll("div").length != 0) {
				closeCommentButton.style.display = "block"
				// }

			}
		})
		let comments = post.comments // Wait for the comments to be fetched

		if (comments !== null) {
			attachCommentsToPost(commentsContainer, comments)
		}


		postsContainer.insertBefore(postContainer, postsContainer.firstChild);

	}


}

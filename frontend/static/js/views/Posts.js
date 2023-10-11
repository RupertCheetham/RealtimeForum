import AbstractView from "./AbstractView.js";
import Nav from "./Nav.js";
import { handleReactions } from "../utils/reactions.js";
import { fetchComments, attachCommentForm, attachCommentsToPost } from "./Comments.js";

// Contains what the Posts page can do, including rendering itself
export default class Posts extends AbstractView {
	constructor() {
		super();
		this.setTitle("Posts");
	}

	async renderHTML() {
		const nav = new Nav(); // Create an instance of the Nav class
		const navHTML = await nav.renderHTML(); // Get the HTML content for the navigation
		const postForm = getPostFormHTML()
		return `
      ${navHTML}
	  ${postForm}
      <div class="contentContainer">
        <div id="leftContainer" class="contentContainer-left">left container</div>
        <div id="postContainer" class="contentContainer-post"></div>
        <div id="rightContainer" class="contentContainer-right">right container, probably chat</div>
      </div>
    `;
	}

	async clearCookie() {
		let logoutBtn = document.getElementById("logout")
		logoutBtn.addEventListener("click", () => {
			document.cookie =
				"sessionID=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/;"
		})
	}

	// The event listener for the post form
	async postSubmitForm() {
		const postForm = document.getElementById("post-form");

		postForm.addEventListener(
			"submit",
			async function (event) {
				event.preventDefault();
				const postText = document.getElementById("postText").value;
				const categories = document.getElementById("categories").value;
				const image = document.getElementById("image").value;
				console.log("submitted post:", postText, categories, image);

				try {
					const response = await fetch("https://localhost:8080/api/addposts", {
						method: "POST",
						headers: {
							Accept: "application/json",
							"Content-Type": "application/json",
						},
						body: JSON.stringify({
							body: postText,
							categories: categories,
							img: image,
						}),
						credentials: "include",
					});

					if (response.ok) {
						// clears the submitted form values, unsure if this helps but apparently it's good practice
						document.getElementById("postText").value = "";
						document.getElementById("categories").value = "";
						document.getElementById("image").value = "";

						// Call displayPostContainer to refresh the post container
						await this.displayPostContainer();
					}
				} catch (error) {
					console.log(error);
				}
			}.bind(this)
		);
	}

	// Gets and displays posts; attaches a comments form to the bottom of each
	async displayPostContainer() {
		const postContainer = document.getElementById("postContainer");
		postContainer.innerHTML = "";

		const response = await fetch("https://localhost:8080/api/getposts", {
			credentials: "include", // Ensure cookies are included in the request
		})

		const posts = await response.json();

		for (const post of posts) {
			let postBox = document.createElement("div");
			postBox.id = "PostBox" + post.id;
			postBox.classList.add("postBox");
			let postElement = document.createElement("div");
			postElement.id = "Post" + post.id;
			postElement.classList.add("post");
			postElement.setAttribute('reactionID', post.reactionID);


			postElement.innerHTML = `
			<ul>
			  <li><b>Id:</b> ${post.id}</li>
			  <li><b>Username:</b> ${post.username}</li>
			  <li><b>Img:</b> ${post.img}</li>
			  <li><b>Body:</b> ${post.body}</li>
			  <li><b>Categories:</b> ${post.categories}</li>
			  <li><b>ReactionID:</b> ${postElement.getAttribute('reactionID')}</li>
			  <li>
			  <button class="reaction-button" reaction-parent-class="post" reaction-parent-id="${post.id}" reaction-action="like" reaction-id = "${postElement.getAttribute('reactionID')}">👍 ${post.postLikes}</button>
			  <button class="reaction-button" reaction-parent-class="post" reaction-parent-id="${post.id}" reaction-action="dislike" reaction-id = "${postElement.getAttribute('reactionID')}">👎 ${post.postDislikes}</button>
			 
			  </li>
			</ul>
		  `;

			// attaches the comment form to the bottom of each post
			attachCommentForm(post, postElement)

			// fetch comments, if any, for this post

			postBox.appendChild(postElement);
			let comments = await fetchComments(post.id); // Wait for the comments to be fetched

			if (comments !== null) {
				let postComments = attachCommentsToPost(comments)
				postBox.appendChild(postComments);
				postComments.style.display = "none";
				postElement.addEventListener("click", () => {
					if (postComments.style.display === "none") {
						postComments.style.display = "block";
					}
				});

			}
			postContainer.appendChild(postBox);
		}
	}

	// Adds reactions to db
	async reactions() {
		handleReactions();
	}

}

function getPostFormHTML() {

	return `<div class="post-form">
        <form id="post-form" method="POST">
          <p>Kindly fill in this form to post.</p>
         
            <div class="post-form-input-field">
              <label for="postText"><b>Post</b></label>
              <input
                type="text"
                placeholder="Enter Message"
                name="postText"
                id="postText"
                required
              />
            </div>
            <div class="post-form-input-field">
              <label for="categories"><b>Categories</b></label>
              <input
                type="text"
                placeholder="Enter Categories"
                name="categories"
                id="categories"
                required
              />
            </div>
            <div class="post-form-input-field">
              <label for="image"><b>Image</b></label>
              <input
                type="text"
                placeholder="Enter Image String"
                name="image"
                id="image"
                required
              />
            </div>
          
          <button class="postSubmitButton" id="submit">Submit Post</button>
        </form>
      </div>`
}
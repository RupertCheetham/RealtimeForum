import AbstractView from "./AbstractView.js";
import Nav from "./Nav.js";

export default class Posts extends AbstractView {
	constructor() {
		super();
		this.setTitle("Posts");
	}

	async getHTML() {
		const nav = new Nav(); // Create an instance of the Nav class
		const navHTML = await nav.getHTML(); // Get the HTML content for the navigation

		return `
      ${navHTML}
      <div class="post-form">
        <form id="post-form" method="POST">
          <p>Kindly fill in this form to post.</p>
          <div class="input-row">
            <div class="input-field">
              <label for="postText"><b>Post</b></label>
              <input
                type="text"
                placeholder="Enter Message"
                name="postText"
                id="postText"
                required
              />
            </div>
            <div class="input-field">
              <label for="categories"><b>Categories</b></label>
              <input
                type="text"
                placeholder="Enter Categories"
                name="categories"
                id="categories"
                required
              />
            </div>
            <div class="input-field">
              <label for="image"><b>Image</b></label>
              <input
                type="text"
                placeholder="Enter Image String"
                name="image"
                id="image"
                required
              />
            </div>
          </div>
          <button class="btn" id="submit">Submit Post</button>
        </form>
      </div>

      <div class="contentContainer">
        <div id="leftContainer" class="contentContainer-left">left container</div>
        <div id="postContainer" class="contentContainer-post"></div>
        <div id="rightContainer" class="contentContainer-right">right container, probably chat</div>
      </div>
    `;
	}

	// The post form at the top of the main page
	async submitForm() {
		const postForm = document.getElementById("post-form");
		console.log("postform is:", postForm);

		postForm.addEventListener(
			"submit",
			async function (event) {
				event.preventDefault();
				const postText = document.getElementById("postText").value;
				const categories = document.getElementById("categories").value;
				const image = document.getElementById("image").value;
				console.log("submitted post:", postText, categories, image);

				try {
					const response = await fetch("http://localhost:8080/posts", {
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
					});

					if (response.ok) {
						document.getElementById("postText").value = "";
						document.getElementById("categories").value = "";
						document.getElementById("image").value = "";

						await this.getPosts();
					}
				} catch (error) {
					console.log(error);
				}
			}.bind(this)
		);
	}

	// Gets and displays posts; attaches a comments form to the bottom of each
	async getPosts() {
		let html = `
      <div>
        <div id="postContainer"></div>
      </div>
    `;

		container.innerHTML += html;

		const response = await fetch("http://localhost:8080/posts");
		const postContainer = document.getElementById("postContainer");
		postContainer.innerHTML = "";

		const posts = await response.json();

		for (const post of posts) {
			let postElement = document.createElement("div");
			postElement.id = "Post" + post.id;
			postElement.classList.add("post");

			let comments = await fetchComments(post.id); // Wait for the comments to be fetched

			postElement.innerHTML = `
			<ul>
			  <li><b>Id:</b> ${post.id}</li>
			  <li><b>Username:</b> ${post.username}</li>
			  <li><b>Img:</b> ${post.img}</li>
			  <li><b>Body:</b> ${post.body}</li>
			  <li><b>Categories:</b> ${post.categories}</li>
			  <li><b>Reaction:</b> ${post.reactionID}</li>
			  <li>
			  <button class="reaction-button" reaction-type="POSTREACTIONS" reaction-parent-id="${post.id}" reaction-action="like" reaction-id = ${post.reactionID}">👍</button>
			  <button class="reaction-button" reaction-type="POSTREACTIONS" reaction-parent-id="${post.id}" reaction-action="dislike" reaction-id = ${post.reactionID}">👎</button>
			  </li>
			</ul>
		  `;

			const commentFormElement = document.createElement("form");
			commentFormElement.id = "comment-form"
			commentFormElement.className = "comment-form"
			commentFormElement.method = "POST"
			commentFormElement.setAttribute('parentPostID', post.id);
			postElement.appendChild(commentFormElement);
			// Adds a comment form to each post.  Laughs at you, scornfully, when you try and figure out why post.id is always the most recent post

			// <form id="comment-form" class="comment-form" method="POST">
			//  </form>
			let commentFormHTML = `
          <label for="commentText"><b>Comment</b></label>
          <input type="text" placeholder="Enter comment" name="commentText" commentText="commentText" id="commentText" required"/><br>
          <input type="submit" value="REPLY" class="btn">
		   <input type="hidden" id="postID" name="postID" value="${post.id}"></input>
      `;
			commentFormElement.innerHTML = commentFormHTML;
			commentFormElement.addEventListener("submit", async function (event) {
				event.preventDefault();

				// Extract data from the submitted form
				const form = event.target;
				const commentText = form.querySelector("#commentText").value;
				const postID = form.querySelector("#postID").value;

				const response = await fetch("http://localhost:8080/comments", {
					method: "POST",
					headers: {
						Accept: "application/json",
						"Content-Type": "application/json",
					},
					body: JSON.stringify({
						body: commentText,
						parentPostId: Number(postID),
					}),
				}).catch((error) => {
					console.log(error)
				})
				if (response.ok) {
					document.getElementById("commentText").value = "";
					// temporary measure
					window.location.reload(); // Refresh the page
					// temporary measure
				}

				console.log(`Comment for post ID ${postID}: ${commentText}`);
			});

			// shows comments underneath post, if it has any to show
			// currently very inefficient, needs reworked
			if (comments.length > 0) {
				const commentsContainer = document.createElement("div");
				commentsContainer.id = "commentContainer";
				commentsContainer.className = "commentContainer";
				let commentsNum = 1;
				comments.forEach((comment) => {
					const commentElement = document.createElement("div");
					commentElement.className = "comment" + commentsNum++;
					// commentElement.textContent = `Comment: ${comment.body}`;
					commentElement.innerHTML = `
					Comment: ${comment.body}
					<ul>
					<button class="reaction-button" reaction-type="COMMENTREACTIONS" reaction-parent-id="${comment.id}" reaction-action="like" reaction-id = ${comment.reactionID}">👍</button>
					<button class="reaction-button" reaction-type="COMMENTREACTIONS" reaction-parent-id="${comment.id}" reaction-action="dislike" reaction-id = ${comment.reactionID}">👎</button>
					</ul>`
					commentsContainer.appendChild(commentElement);
				});

				postElement.appendChild(commentsContainer);
			}
			// commentForm.setAttribute('parentPostID', post.id);
			postContainer.appendChild(postElement);
		}

		// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
		async function fetchComments(parentPostID) {
			const response = await fetch("http://localhost:8080/comments");
			const comments = await response.json();
			return comments.filter((comment) => comment.parentPostId == parentPostID);
		}
	}

	// Adds reactions to db
	async reactions() {
		const reactionButtons = document.querySelectorAll('.reaction-button');
		reactionButtons.forEach((reactButton) => {
			reactButton.addEventListener('click', async (event) => {
				event.preventDefault();
				const Action = reactButton.getAttribute('reaction-action');
				const Type = reactButton.getAttribute('reaction-type');
				const ParentID = reactButton.getAttribute('reaction-parent-id');
				const ReactionID = reactButton.getAttribute('reaction-id');
				// Placeholder UserID
				const UserID = 1;
				// Placeholder UserID

				console.log(`Reacted to ${Type} ${ParentID} with action: ${Action}, whilst reactionID is ${ReactionID}`);

				fetch("http://localhost:8080/reaction", {
					method: "POST",
					headers: {
						Accept: "application/json",
						"Content-Type": "application/json",
					},
					body: JSON.stringify({
						userID: UserID,
						type: Type,
						parentID: parseInt(ParentID),
						action: Action,
						reactionID: parseInt(ReactionID),
					}),
				}).catch((error) => {
					console.log(error)
				})
				// window.location.reload();
			});
		});

	}
}





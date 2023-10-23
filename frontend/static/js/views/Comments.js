import { userIDFromSessionID } from "../utils/utils.js";

// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
export async function fetchComments(parentPostID) {
	const response = await fetch(`https://localhost:8080/api/getcomments?postID=${parentPostID}`, {
		credentials: "include", // Ensure cookies are included in the request
	});
	const comments = await response.json();
	return comments
}

// attaches the comment form to the bottom of each post, if you've placed it in the right place
export function attachCommentForm(post, postElement) {
	const commentFormElement = document.createElement("form");
	commentFormElement.id = "comment-form"
	commentFormElement.className = "comment-form"
	commentFormElement.method = "POST"
	commentFormElement.setAttribute('parentPostID', post.id);
	postElement.appendChild(commentFormElement);
	// Adds a comment form to each post.  Laughs at you, scornfully, when you try and figure out why post.id is always the most recent post

	// <form id="comment-form" class="comment-form" method="POST">
	//  </form>
	commentFormElement.innerHTML = getCommentFormHTML(post.id);
	commentFormElement.addEventListener("submit", async function (event) {
		event.preventDefault();

		// Extract data from the submitted form
		const form = event.target;
		const currentUserID = await userIDFromSessionID()
		const commentText = form.querySelector("#commentText").value;
		const postID = form.querySelector("#postID").value;

		const response = await fetch("https://localhost:8080/api/addcomments", {

			method: "POST",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				userID: currentUserID,
				body: commentText,
				parentPostId: Number(postID),
			}),
			credentials: "include",
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
}

// adds the com
export function attachCommentsToPost(comments) {
	const commentsContainer = document.createElement("div");
	commentsContainer.id = "commentContainer";
	commentsContainer.className = "commentContainer";
	comments.forEach((comment) => {
		const commentElement = document.createElement("div");
		commentElement.id = "comment" + comment.id
		commentElement.className = "comment";
		commentElement.setAttribute('reactionID', comment.reactionID);
		commentElement.innerHTML = `
		
		<li>Username: ${comment.username}</li>
		<li>Comment: ${comment.body}</li>
					<ul>
					<button class="reaction-button" reaction-action="like" reaction-parent-class="comment" reaction-parent-id="${comment.id}"  reaction-id = "${commentElement.getAttribute('reactionID')}">👍 ${comment.commentLikes}</button>
					<button class="reaction-button" reaction-action="dislike" reaction-parent-class="comment" reaction-parent-id="${comment.id}"  reaction-id = "${commentElement.getAttribute('reactionID')}">👎 ${comment.commentDislikes}</button>
					</ul>
		`
		commentsContainer.appendChild(commentElement);
	});
	const closeCommentsButton = document.createElement("button");
	closeCommentsButton.id = "closeCommentsButton";
	closeCommentsButton.innerText = "Close Comments";

	// Add an event listener to the "Close" button to hide the comments container
	closeCommentsButton.addEventListener("click", () => {
		commentsContainer.style.display = "none";
		console.log("made")
	});

	// Append the "Close" button to the comments container
	commentsContainer.appendChild(closeCommentsButton);
	return commentsContainer;
}

// The comment submission form
function getCommentFormHTML(postID) {
	return `
	<label for="commentText"><b>Comment</b></label>
	<input type="text" placeholder="Enter comment" name="commentText" commentText="commentText" id="commentText" required"/><br>
	<input type="submit" value="REPLY" class="btn">
	<input type="hidden" id="postID" name="postID" value="${postID}"></input>
	`
}
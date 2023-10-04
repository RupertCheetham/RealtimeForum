// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
export async function fetchComments(parentPostID) {
	const response = await fetch("http://localhost:8080/comments");
	const comments = await response.json();
	return comments.filter((comment) => comment.parentPostId == parentPostID);
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
}

export function attachCommentsToPost(comments) {
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

	return commentsContainer;
}
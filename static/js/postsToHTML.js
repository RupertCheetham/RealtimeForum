const container = document.getElementById("container")

export async function postsToHTML() {
	let html = `
    <h1>Post</h1>
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
    <br>
    <form id="comment-form">
    <h1>Comment</h1>
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

	const response = await fetch("http://localhost:8080/posts")
	const postContainer = document.getElementById("postContainer")
	const posts = await response.json()

	for (const post of posts) {
		const postElement = document.createElement("div")
		postElement.id = "Post" + post.id
		postElement.classList.add("post")

		const comments = await fetchComments(post.id) // Wait for the comments to be fetched
		console.log(comments)

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
		}

		postContainer.appendChild(postElement)
	}
}
postsToHTML()

// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
async function fetchComments(parentPostID) {
	const response = await fetch("http://localhost:8080/comments")
	const comments = await response.json()
	console.log("comments:", comments)
	return comments.filter((comment) => comment.parentPostId == parentPostID)
}

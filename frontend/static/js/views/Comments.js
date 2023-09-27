import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"

export default class extends AbstractView {
    
	constructor() {
		super()
		this.setTitle("Comments")
	}

	async getHTML() {
		const nav = new Nav() // Create an instance of the Nav class
		const navHTML = await nav.getHTML() // Get the HTML content for the navigation

		return `
		${navHTML}
    <h1>Posts</h1>
		<form id="post-form" method="POST">
			<div>
				<p>Kindly fill in this form to post.</p>
				<label for="post"><b>Post</b></label>
				<input
					type="text"
					placeholder="Enter Message"
					name="postText"
					id="postText"
					required
				/>
				<br />
				<label for="categories"><b>Categories</b></label>
				<input
					type="text"
					placeholder="Enter Categories"
					name="categories"
					id="categories"
					required
				/>
				<br />
				<label for="image"><b>Image</b></label>
				<input
					type="text"
					placeholder="Enter Image String"
					name="image"
					id="image"
					required
				/>
				<br />
				<button id="submit">Submit Post</button>
			</div>
		</form>
		<div id="postContainer"></div>
    `
	}

	// async submitForm() {
	// 	const postForm = document.getElementById("post-form")
	// 	console.log("postform is:", postForm)

	// 	postForm.addEventListener(
	// 		"submit",
	// 		function (event) {
	// 			event.preventDefault()
	// 			const postText = document.getElementById("postText").value
	// 			const categories = document.getElementById("categories").value
	// 			const image = document.getElementById("image").value
	// 			console.log("submitted post:", postText, categories, image)

	// 			fetch("http://localhost:8080/posts", {
	// 				method: "POST",
	// 				headers: {
	// 					Accept: "application/json",
	// 					"Content-Type": "application/json",
	// 				},
	// 				body: JSON.stringify({
	// 					body: postText,
	// 					categories: categories,
	// 					img: image,
	// 				}),
	// 			})
	// 				.then(async (response) => {
	// 					if (response.ok) {
	// 						document.getElementById("postText").value = ""
	// 						document.getElementById("categories").value = ""
	// 						document.getElementById("image").value = ""

	// 						await this.getPosts()
	// 					}
	// 				})
	// 				.catch((error) => {
	// 					console.log(error)
	// 				})
	// 		}.bind(this)
	// 	)
	// }

	async getPosts() {
		let html = `
		<div>
			<div id="postContainer"></div>
		</div>
	`

		container.innerHTML += html

		const response = await fetch("http://localhost:8080/posts")
		const postContainer = document.getElementById("postContainer")
		postContainer.innerHTML = ""

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

			}
			postElement.innerHTML += commentHTML
			postContainer.appendChild(postElement)
		}

		// Comments need to be reworked, currently very inefficient.  Probably foreign keys will be involved
		async function fetchComments(parentPostID) {
			const response = await fetch("http://localhost:8080/comments")
			const comments = await response.json()
			return comments.filter((comment) => comment.parentPostId == parentPostID)
		}
	}
	/* The `async submitCommentForm()` function is responsible for handling the submission of the comment
    form. It listens for the "submit" event on the comment form, prevents the default form submission
    behavior, and retrieves the comment text from the input field. */
    async submitCommentForm() {
        
        console.log("I'm comment form man; look at me go!")
        const commentForm = document.getElementById("comment-form")
        //console.log(commentForm)
    
        commentForm.addEventListener(
             "submit",
            function (event) {
                event.preventDefault()
                const commentText = document.getElementById("commentText").value
                
                console.log("this is comment text", commentText)
    
                fetch("http://localhost:8080/comments", {
                    method: "POST",
                    headers: {
                        Accept: "application/json",
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        body: commentText,
                    }),
                })
                    .then(async (response) => {
                        if (response.ok) {
                            document.getElementById("commentText").value = ""
                            await this.getPosts()
                        }
                    })
                    .catch((error) => {
                        console.log(error)
                    })
            }.bind(this)
        )

        }
}


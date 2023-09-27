import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"

export default class extends AbstractView {
	constructor() {
		super()
		this.setTitle("Posts")
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
}

import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"
import Post from "./Post.js"
const nav = new Nav()
const postClass = new Post()

export default class UserPage extends AbstractView {
	async renderHTML() {
		const navHTML = await nav.renderHTML()
		return `
    ${navHTML}
      <div class="contentContainer">
        <div id="likedContainer" class="contentContainer-user">Liked Posts</div>
        <div id="postsContainer" class="contentContainer-post"></div>
        <div id="chatContainer" class="contentContainer-chat"></div>
      </div>
    `
	}

	async getAllPostsByUser() {
		const response = await fetch("https://localhost:8080/api/getuserposts", {
			credentials: "include", // Ensure cookies are included in the request
		})

		if (response.status == 408) {
			localStorage.clear()
			window.location.href = "/"
		}

		const posts = await response.json()
		const postsContainer = document.getElementById("postsContainer")
		postsContainer.innerHTML = ""
		for (const post of posts) {
			postClass.processPost(postsContainer, post)
		}
	}

	async getLikedPostsByUser() {
		const response = await fetch("https://localhost:8080/api/getuserposts", {
			credentials: "include", // Ensure cookies are included in the request
		})

		if (response.status == 408) {
			localStorage.clear()
			window.location.href = "/"
		}

		const posts = await response.json()
		const postsContainer = document.getElementById("likedContainer")
		postsContainer.innerHTML = ""

		renderPosts(posts, postsContainer)
	}

	async Logout() {
		nav.logout()
	}
}

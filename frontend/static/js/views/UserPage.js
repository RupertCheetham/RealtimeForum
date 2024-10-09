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
        <div id="likedContainer" class="contentContainer-user"></div>
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
		const heading = document.createElement("h1");
		heading.textContent = "Posts By User";

		if (posts) {
			for (const post of posts) {
				postClass.processPost(postsContainer, post)
			}

			postsContainer.insertBefore(heading, postsContainer.firstChild);
		} else {
			postsContainer.appendChild(heading);
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

		const likedContainer = document.getElementById("likedContainer")
		const heading = document.createElement("h1");
		heading.textContent = "Liked Posts By User";

		if (posts) {
			for (const post of posts) {
				postClass.processPost(likedContainer, post)
			}

			likedContainer.insertBefore(heading, likedContainer.firstChild);
		} else {
			likedContainer.append(heading)
		}
	}

	async Logout() {
		nav.logout()
	}
}

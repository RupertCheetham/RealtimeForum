import AbstractView from "./AbstractView.js"
import Nav from "./Nav.js"
import { clearCookie } from "./Auth.js"
import PostSubmitForm from "./PostSubmitForm.js"
import Posts from "./Post.js"
import Chat from "./Chat.js"
import { handleReactions } from "../utils/reactions.js"

const postSubmitForm = new PostSubmitForm()
const post = new Posts()
const chat = new Chat()
const nav = new Nav()

// Contains what the main page can do, including rendering itself
export default class Mainpage extends AbstractView {
	constructor() {
		super()
		this.setTitle("Mainpage")
	}

	async renderHTML() {
		const navHTML = await nav.renderHTML()
		const postForm = await postSubmitForm.renderHTML()
		return `
    ${navHTML}
	  ${postForm}
      <div class="contentContainer">
        <div id="userContainer" class="contentContainer-user">Users</div>
        <div id="postContainer" class="contentContainer-post"></div>
        <div id="chatContainer" class="contentContainer-chat">Chat (click on Username)</div>
      </div>
    `
	}

	async Logout() {
		nav.logout()
	}

	// The event listener for the post form
	async attachPostSubmitForm() {
		await postSubmitForm.handlePostSubmission()
	}

	async displayUserContainer() {
		await chat.userList()
	}

	async displayPostContainer() {
		await post.renderHTML()
	}

	async displayChatContainer() {
		await chat.renderHTML()
	}

	// Adds reactions to db
	async reactions() {
		handleReactions()
	}
}

import AbstractView from "./AbstractView.js";
import Nav from "./Nav.js";
import { clearCookie } from "./Auth.js";
import { getPostFormHTML, postSubmitForm, handlePostContainer} from "./Post.js";
import { handleReactions } from "../utils/reactions.js";
import { userList } from "./Chat.js";

// Contains what the main page can do, including rendering itself
export default class Mainpage extends AbstractView {
  constructor() {
    super();
    this.setTitle("Mainpage");
  }

  async renderHTML() {
    const nav = new Nav(); // Create an instance of the Nav class
    const navHTML = await nav.renderHTML(); // Get the HTML content for the navigation
    const postForm = getPostFormHTML();
    return `
      ${navHTML}
	  ${postForm}
      <div class="contentContainer">
        <div id="userContainer" class="contentContainer-user">user container</div>
        <div id="postContainer" class="contentContainer-post"></div>
        <div id="rightContainer" class="contentContainer-right">right container, probably chat</div>
      </div>
    `;
  }

  async Logout(){
    await clearCookie()
  }

  // The event listener for the post form
  async attachPostSubmitForm() {
    await postSubmitForm()
  }

  async displayUserContainer() {
    await userList()
  }

  async displayPostContainer() {
    await handlePostContainer();
  }
  
  // Adds reactions to db
  async reactions() {
    handleReactions();
  }
}


import AbstractView from "./AbstractView.js";
import { userIDFromSessionID } from "../utils/utils.js";
import Posts from "./Post.js";

export default class PostSubmitForm extends AbstractView{

  async renderHTML() {
    return `<div class="post-form">
    <form id="post-form" method="POST">
      <p>Kindly fill in this form to post.</p>
     
        <div class="post-form-input-field">
          <label for="postText"><b>Post</b></label>
          <input
            type="text"
            placeholder="Enter Message"
            name="postText"
            id="postText"
            required
          />
        </div>

        <label for="categories"><b>Categories</b></label>
        <ul>
       <label><input type="checkbox" name="Category" value="Dogs"> Dogs </label>
        <label><input type="checkbox" name="Category" value="Sausages"> Sausages</label>
        <label><input type="checkbox" name="Category" value="Cats"> Cats</label>
        <label><input type="checkbox" name="Category" value="Meows"> Meows </label>
       </ul>
        <div class="post-form-input-field">
          <label for="image"><b>Image</b></label>
          <input
            type="text"
            placeholder="Enter Image String"
            name="image"
            id="image"
            required
          />
        </div>
      
      <button class="postSubmitButton" id="submit">Submit Post</button>
    </form>
  </div>`;
  }

  async handlePostSubmission() {
    const postForm = document.getElementById("post-form");

    postForm.addEventListener(
      "submit",
      async function (event) {
        event.preventDefault();
        const currentUserID = await userIDFromSessionID()
        const postText = document.getElementById("postText").value;
        const categoriesCheckboxes = document.querySelectorAll('input[name="Category"]:checked');
        const categories = Array.from(categoriesCheckboxes).map(categoriesCheckboxes => categoriesCheckboxes.value);

        const image = document.getElementById("image").value;
        console.log("submitted post:", postText, categories, image);

        try {
          const response = await fetch("https://localhost:8080/api/addposts", {
            method: "POST",
            headers: {
              Accept: "application/json",
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              userID: currentUserID,
              body: postText,
              categories: categories,
              img: image,
            }),
            credentials: "include",
          });

          if (response.ok) {
            // clears the submitted form values, unsure if this helps but apparently it's good practice
            document.getElementById("postText").value = "";
            const checkboxes = document.querySelectorAll('input[type="checkbox"]');
            checkboxes.forEach(checkbox => { checkbox.checked = false; });
            document.getElementById("image").value = "";
            const post = new Posts();
            // Call displayPostContainer to refresh the post container
            await post.renderHTML()
          }
        } catch (error) {
          console.log(error);
        }
      }.bind(this)
    );
  }
}



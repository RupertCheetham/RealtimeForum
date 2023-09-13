/**
 * The function `postsToHTML` fetches posts from a server and dynamically creates HTML elements to
 * display the post information on a webpage.
 */
async function postsToHTML() {
  const response = await fetch('http://localhost:8080/posts');

  const postContainer = document.getElementById('postContainer');
  response.json().then(posts => {
    posts.forEach(post => {
      const postElement = document.createElement('div');
      postElement.setAttribute("id", "Post" + post.id);
      postElement.textContent = `
            Id: ${post.id},
            Username: ${post.username},
            Img: ${post.img},
            Body: ${post.body},
            Categories: ${post.categories}
            Likes: ${post.likes},
            Dislikes: ${post.dislikes},
            WhoLiked: ${post.whoLiked},
            WhoDisliked: ${post.whoDisliked}
        `;
      postContainer.appendChild(postElement);
    });
  });
}

postsToHTML();

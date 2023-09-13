async function postsToHTML() {
  const response = await fetch('http://localhost:8080/posts');
  const postContainer = document.getElementById('postContainer');
  const posts = await response.json();

  for (const post of posts) {
    const postElement = document.createElement('div');
    postElement.id = 'Post' + post.id;
    postElement.classList.add('post');

    const comments = await fetchComments(post.id); // Wait for the comments to be fetched
    console.log(comments);

    postElement.textContent = `
      Id: ${post.id},
      Username: ${post.username},
      Img: ${post.img},
      Body: ${post.body},
      Categories: ${post.categories},
      Likes: ${post.likes},
      Dislikes: ${post.dislikes},
      WhoLiked: ${post.whoLiked},
      WhoDisliked: ${post.whoDisliked},
    `;

    if (comments.length > 0) {
      const commentsContainer = document.createElement('div');
      commentsContainer.id = 'commentContainer';
      let commentsNum = 1
      comments.forEach(comment => {
        const commentElement = document.createElement('div');
        commentElement.className = 'comment' + commentsNum++;
        commentElement.textContent = `Comment: ${comment.body}`;
        commentsContainer.appendChild(commentElement);
      });

      postElement.appendChild(commentsContainer);
    }

    postContainer.appendChild(postElement);
  }
}

async function fetchComments(parentPostID) {
  const response = await fetch('http://localhost:8080/comments');
  const comments = await response.json();
  console.log("comments:", comments)
  return comments.filter(comment => comment.parentPostId == parentPostID);
}

postsToHTML();

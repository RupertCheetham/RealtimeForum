export function handleReactions() {
    const reactionButtons = document.querySelectorAll('.reaction-button');
    reactionButtons.forEach((reactButton) => {
        reactButton.addEventListener('click', async (event) => {
            event.preventDefault();
            const Action = reactButton.getAttribute('reaction-action');
            const Type = reactButton.getAttribute('reaction-type');
            const ParentID = reactButton.getAttribute('reaction-parent-id');
            const ReactionID = reactButton.getAttribute('reaction-id');
            // Placeholder UserID
            const UserID = 1;
            // Placeholder UserID

            console.log(`Reacted to ${Type} ${ParentID} with action: ${Action}, whilst reactionID is ${ReactionID}`);

            fetch("http://localhost:8080/reaction", {
                method: "POST",
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    userID: UserID,
                    type: Type,
                    parentID: parseInt(ParentID),
                    action: Action,
                    reactionID: parseInt(ReactionID),
                }),
            })
                .then(async (response) => {
                    if (response.ok) {
                        // If the POST request was successful, make a GET request for reactionID
                        const reactionData = await fetch(`http://localhost:8080/reaction?parentID=${parseInt(ParentID)}&rowID=${parseInt(ReactionID)}&reactionTable=${Type}`);

                        if (reactionData.ok) {
                            const data = await reactionData.json();
                            console.log("Reaction data:", data);

                            // Update the like and dislike buttons in the DOM
                            const likeButton = document.querySelector(`.reaction-button[reaction-id="${ReactionID}"][reaction-action="like"]`);
                            console.log("likeButton", likeButton)
                            const dislikeButton = document.querySelector(`.reaction-button[reaction-id="${ReactionID}"][reaction-action="dislike"]`);
                            if (likeButton && dislikeButton) {
                                likeButton.innerText = `👍 ${data.Likes}`;
                                dislikeButton.innerText = `👎 ${data.Dislikes}`;
                                console.log("made it")
                            }
                        } else {
                            console.log("Error fetching reaction data:", reactionData.statusText);
                        }
                    } else {
                        console.log("Error with the POST request:", response.statusText);
                    }
                })
                .catch((error) => {
                    console.log("Error with the POST request:", error);
                });
        });
    });
}

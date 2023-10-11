export function handleReactions() {
    
    const postContainer = document.getElementById("postContainer");
    postContainer.addEventListener('click', async (event) => {
        const target = event.target;

        // Check if the clicked element has the class 'reaction-button'
        if (target.classList.contains('reaction-button')) {
            event.preventDefault();
            const Action = target.getAttribute('reaction-action');
            const Type = target.getAttribute('reaction-parent-class');
            const ParentID = target.getAttribute('reaction-parent-id');
            const postElement = target.closest('.' + Type);
            const ReactionID = postElement.getAttribute('reactionID');
            // Placeholder UserID
            const UserID = 1;
            // Placeholder UserID

            console.log(`Reacted to ${Type} ${ParentID} with action: ${Action}, whilst reactionID is ${ReactionID}`);

            fetch("https://localhost:8080/reaction", {
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
                credentials: "include",
            })
                .then(async (response) => {
                    if (response.ok) {
                        // If the POST request was successful, make a GET request for reactionID
                        const reactionData = await fetch(`https://localhost:8080/reaction?rowID=${parseInt(ReactionID)}&reactionParentClass=${Type}`, {
                            credentials: "include", // Ensure cookies are included in the request
                        });

                        if (reactionData.ok) {
                            const data = await reactionData.json();

                            const otherButtonAction = Action === 'like' ? 'dislike' : 'like';
                            const otherButton = postElement.querySelector(`.reaction-button[reaction-action="${otherButtonAction}"]`);
                            if (Action == "like") {
                                target.innerText = `👍 ${data.Likes}`;
                                otherButton.innerText = `👎 ${data.Dislikes}`;
                            } else {
                                otherButton.innerText = `👍 ${data.Likes}`;
                                target.innerText = `👎 ${data.Dislikes}`;
                            }

                            // Update the reactionID attribute
                            postElement.setAttribute('reactionID', data.ReactionID);
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
        }
    });

}



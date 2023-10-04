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
            }).catch((error) => {
                console.log(error)
            })
            // window.location.reload();
        });
        
    });

}
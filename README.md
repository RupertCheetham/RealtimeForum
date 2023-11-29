# REALTIME FORUM

Welcome to the Realtime Forum, a project designed to provide an interactive and dynamic forum experience using a combination of Golang for the backend, JavaScript for frontend interactions, and the power of WebSockets for real-time communication. This project builds upon the foundations of a previous forum implementation, introducing enhanced features such as private messaging, real-time actions, and improved post interactions.

## Technologies and Learning Focus

The Realtime Forum serves as a practical learning experience in various web technologies, covering the basics of web development, Golang features like goroutines and channels, WebSockets implementation on both the backend and frontend, SQL database manipulation, and the secure handling of user data. The project emphasizes the importance of understanding the interaction between frontend and backend components, fostering skills in communication between different server-side and client-side technologies.

Note: The project prohibits the use of frontend libraries or frameworks, challenging developers to delve into the fundamentals of web development and gain a deeper understanding of the technologies involved.

## Dependencies

    Go Packages:
       - go-sqlite3: Lightweight SQL database for Go.
       - google/uuid: Package for generating UUIDs.
       - gorilla/websocket: Enables WebSocket functionality.
        -x/crypto: Used for encrypting data, such as user passwords.

        These are downloaded automatically when you run the backend server, instructions below

## How to Run

Our forum uses security certificates made by OpenSSL to make it more secure from outside threats, this does add an extra step to getting the forum up and running the first time

1. Download the repository and a code editor of your choice
2. Open the terminal/code editor and navigate to the root directory of the repo.
3. Run the following commands:
   "cd backend"
   "go run ."
4. Wait for the program to download the required packages. This might take a few minutes based on your internet speed. 
(Packages:
go-sqlite3 - a light version of SQL database
google/uuid - a google package for making unique numbers
gorilla/websocket - used to easily enable websocket functionality
x/crypto - used to encrypt data, in this case the users password
)
5. Without closing the original terminal, open a new terminal and run the following commands:
   "cd frontend"
   "node server.mjs"
6. The message "Server is running on port:3000" should appear in the terminal, this should happen near immediately.
7. Open your browser and navigate to [https://localhost:8080/](https://localhost:8080/) (the backend server) and accept the security certificates, now do the same for [https://localhost:3000/](https://localhost:3000/), this will take you to our forum.



## Account Information for Testing:

You can make a new account, or if you'd prefer we made a couple of easy to log in accounts to test the functionality
Username: t
Password: t
and
Username: r
Password: r

## Features

- Registration (mandatory)
- View posts, likes/dislikes, categories, and comments.
- Register an account to:
  - Like/dislike posts and comments.
  - Access your posts and liked posts via your username in the top left of the screen.
  - Realtime Chat by clicking on a user from the list on the left (need to log in to 2 different users, 1 per browser)


## Skills Demonstrated

- JS
- Node.JS
- HTML
- CSS
- Golang
- Websockets
- Sessions and cookies
- Communication between a frontend and a backend server
- DOM
- HTTPS security certificates and keys

## Team Members

- Martin F. - [GitHub](https://github.com/m-fenton)
- Michael A. - [GitHub](https://github.com/11ma)
- Nikoi K. - [GitHub](https://github.com/kn1ko1)
- Captain Rupert C. - [GitHub](https://github.com/Cheethamthing)



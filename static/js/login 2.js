// Get references to HTML elements
const usernameOrEmailInput = document.getElementById('usernameOrEmail');
const passwordInput = document.getElementById('password');
const outputElement = document.getElementById('output');

  document.getElementById('log-in-form').addEventListener('submit', (e) => {
    e.preventDefault();

  const usernameOrEmail = usernameOrEmailInput.value;
  const password = passwordInput.value;

  // Make an HTTP GET request to your database API
  fetch(`http://localhost:8080/registrations?search=${usernameOrEmail}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then((users) => {
      const user = users.find((user) => {
        return user.nickname === usernameOrEmail || user.email === usernameOrEmail;
      });

      if (user) {
        // User found, check if the provided password matches
        if (user.password === password) {
          outputElement.textContent = `Login successful. Welcome, ${user.nickname}`;
          console.log("Login successful. Welcome.")
        } else {
          outputElement.textContent = 'Incorrect password. Login failed.';
          console.log("Incorrect password. Login failed.")
        }
      } else {
        outputElement.textContent = 'User not found. Please check your credentials.';
        console.log("User not found. Please check your credentials.")
      }
    })
    .catch((error) => {
      console.error('Error:', error);
      outputElement.textContent = 'An error occurred while fetching user data.';
      console.log("An error occurred while fetching user data.")
    });
});


    
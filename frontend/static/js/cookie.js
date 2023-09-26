export function generateSessionId() {
    let n = 1000
    return n++
} 

// Function to set a session cookie with an expiration time
export function setSessionCookie() {

    const sessionId = generateSessionId();

    const currentTime = new Date();
    console.log('Current Time:', currentTime);

    // Calculate the expiration time (e.g., 1 hour from now)
    const expirationDate = new Date();
    expirationDate.setTime(expirationDate.getTime() + 60 * 60 * 1000); // 1 hour in milliseconds

    // Create the cookie string
    const cookieString = `sessionID=${sessionId}; expires=${expirationDate.toUTCString()}; path=/`;

    // Set the cookie
    document.cookie = cookieString;
    console.log("cookie set")
}

// Call the function to set the session cookie
// setSessionCookie();


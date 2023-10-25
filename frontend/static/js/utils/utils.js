// Function to get a cookie by name
export function getCookie(name) {
	const cookieString = document.cookie
	const cookies = cookieString.split("; ")
	for (const cookie of cookies) {
		const [cookieName, cookieValue] = cookie.split("=")
		if (cookieName === name) {
			return decodeURIComponent(cookieValue)
		}
	}
	return ""
}

// export async function userNameFromSessionID(sessionID) {
// 	try {
// 	  const response = await fetch(`https://localhost:8080/api/getUsername?sessionID=${sessionID}`, {
// 		credentials: "include",
// 	  });

// 	  if (response.ok) {
// 		const username = await response.json();

// 		return username;
// 	  } else {
// 		// Handle the error or return a default value in case of an error
// 		console.error("Failed to fetch username:", response.status, response.statusText);
// 		return null; // or return a default value, or throw an error
// 	  }
// 	} catch (error) {
// 	  console.error("An error occurred while fetching username:", error);
// 	  throw error;
// 	}
//   }

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

export function checkSessionTimeout(response) {
	if (response.status == 408) {
		window.location.href = "/"
	}
	return
}

// A throttle function to limit the frequency of calling another function
export function throttle(func, delay) {
	let lastCall = 0
	return function (...args) {
		const now = new Date().getTime()
		if (now - lastCall >= delay) {
			func(...args)
			lastCall = now
		}
	}
}

export async function usernameFromUserID(userID) {

	try {
		const response = await fetch(`https://localhost:8080/api/getUsernameFromUserID?userID=${userID}`, { //
			credentials: "include",
		});

		if (response.ok) {
			const username = await response.json();

			return username;
		} else {
			// Handle the error or return a default value in case of an error
			console.error("Failed to fetch username:", response.status, response.statusText);
			return null; // or return a default value, or throw an error
		}
	} catch (error) {
		console.error("An error occurred while fetching username:", error);
		throw error;
	}
}
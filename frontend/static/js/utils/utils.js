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
import * as fs from "fs"
import https from "https"
import path from "path"

const hostname = "localhost"
const port = 3000
const __dirname = path.dirname(new URL(import.meta.url).pathname)
let urlReg = new RegExp(`\/.`)

// Load your SSL certificate and key
const privateKey = fs.readFileSync("client.key", "utf8")
const certificate = fs.readFileSync("client.crt", "utf8")
const credentials = { key: privateKey, cert: certificate }

const handler = (req, res) => {
	// Set the response header
	res.setHeader("Content-Type", "text/html")

	// Check if the request URL starts with "/static/"
	if (req.url.startsWith("/static/")) {
		// Serve static files from the "static" directory
		const staticFilePath = path.join(
			__dirname,
			"static",
			req.url.replace("/static/", "")
		)

		fs.readFile(staticFilePath, (err, data) => {
			if (err) {
				res.statusCode = 404
				res.end("Not Found\n")
			} else {
				// Determine the content type based on the file extension
				const ext = path.extname(staticFilePath).toLowerCase()
				const contentType =
					{
						".html": "text/html",
						".css": "text/css",
						".js": "application/javascript",
					}[ext] || "application/octet-stream" // Default to binary

				res.setHeader("Content-Type", contentType)
				res.statusCode = 200
				res.end(data)
			}
		})
	} else if (urlReg.test(req.url) || req.url === "/") {
		// Respond with index for the root URL
		fs.readFile(path.join(__dirname, "index.html"), "utf8", (err, data) => {
			if (err) {
				res.statusCode = 500
				res.end("Internal Server Error\n")
				return
			}
			res.statusCode = 200
			res.end(data)
		})
	} else {
		// Respond with a 404 Not Found for other URLs
		res.statusCode = 404
		res.end("Not Found\n")
	}
}

// Create an HTTP server
// const server = http.createServer(handler)
const server = https.createServer(credentials, handler)

// Start the server
server.listen(port, hostname, () => {
	console.log(`Server is running on port:${port}`)
})

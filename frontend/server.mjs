import * as fs from "fs"
import path, { dirname } from "path"
import https from "https"

const hostname = "localhost"
//const port = 3000
const httpsPort = 3001 // Use a different port to HTTPS

const __dirname = path.dirname(new URL(import.meta.url).pathname)

const options = {
	key: fs.readFileSync(path.join(__dirname, "../backend", "server.key")),
	cert: fs.readFileSync(path.join(__dirname, "../backend", "server.crt"))
}

let urlReg = new RegExp(`\/.`)

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
		// Respond with "Hello, World!" for the root URL
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
		res.end("Not Found\n bitat")
	}
}

// Create an HTTPS server
const server = https.createServer(options, handler)

// Start the server
server.listen(httpsPort, hostname, () => {
	console.log(`HTTPS Server is running on port:${httpsPort}`)
})

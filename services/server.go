package forum

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	//shall we use a servemux?
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	fmt.Println("HTTPS(!)Server Starting on port 7000...")

	// interestingly this version of listen and serve needs to check the .pem files; this version of the site needs to be
	// accessed through HTTPS://localhost:7000
	// The browser won't trust it readily, but it's a first step in dealing with security certificates
	err := http.ListenAndServeTLS(":7000", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("HTTPS server failed to start: ", err, ".  If unsure of why then you may need to renew the cert.pem/key.pem")
	}
}

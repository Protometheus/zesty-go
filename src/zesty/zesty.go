package main

import (
	"os"
	// "log"
    "fmt"
    "net/http"
	
	// "github.com/gin-gonic/gin"
)

func base_handler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/code/challenge", 301)
}

func endpoint_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	
	port = ":" + port
	
	// gin.New()
	
	fmt.Println(port)

    http.HandleFunc("/", base_handler)
	http.HandleFunc("/code/challenge", endpoint_handler)
    http.ListenAndServe(port, nil)
}

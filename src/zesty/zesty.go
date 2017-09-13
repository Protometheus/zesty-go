package main

import (
	"os"
	// "log"
    "fmt"
    "net/http"
	
	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	
	port = ":" + port
	
	gin.New()
	
	fmt.Println(port)

    http.HandleFunc("/", handler)
    http.ListenAndServe(port, nil)
}

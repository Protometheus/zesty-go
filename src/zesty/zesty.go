// package main

// import (
//     "fmt"
//     "net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
//     http.HandleFunc("/", handler)
//     http.ListenAndServe(":5000", nil)
// }
package main

import (

	"log"
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	//router.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.tmpl.html", nil)
        c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port)
}
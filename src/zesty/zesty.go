package main

import (
	"os"
    "fmt"
    "net/http"
	"encoding/json"
)

// Struct to represent name; there might be a more Go-way to do this
type Name struct {
	First 	string	`json:"first"`
	Last 	string	`json:"last"`
}

// Struct to represent return JSON
type PersonalInfo struct {
	Name 				Name	`json:"name"`
	Github_repo_link 	string	`json:"github_repo_link"`
	Website 			string	`json:"website"`
	Email				string	`json:"email"`
}

// Handle hits to '/'
func base_handler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/code/challenge", 301)
}

// Handle hits to '/code/challenge'
func endpoint_handler(w http.ResponseWriter, r *http.Request) {
	personal_info := PersonalInfo{Name{"Mohammad", "Oweis"}, "https://github.com/Protometheus/zesty-go",  "https://github.com/Protometheus/zesty-go", "oweismo.applicant@gmail.com"}

	js, _ := json.Marshal(personal_info)
	
    // this should be handled by framework
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}

func main() {
    // Get our heroku port
	port := os.Getenv("PORT")

    // Or set it ourselves
	if port == "" {
		port = "5000"
	}
	
	port = ":" + port
	
    // set routes
	http.HandleFunc("/", base_handler)
	http.HandleFunc("/code/challenge", endpoint_handler)
	
    // start server
	fmt.Println("Starting Server... on port: " + port)
    http.ListenAndServe(port, nil)
}

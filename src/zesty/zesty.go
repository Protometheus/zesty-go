package main

import (
	"os"
    "fmt"
    "net/http"
	"json"
)

type Name struct {
	first 	string
	last 	string
}

type PersonalInfo struct {
	name 				Name
	github_repo_link 	string
	website 			string
	email				string
}

func base_handler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/code/challenge", 301)
}

func endpoint_handler(w http.ResponseWriter, r *http.Request) {
	personal_info := PersonalInfo{Name{"Mohammad", "Oweis"}, "https://github.com/Protometheus/zesty-go",  "https://github.com/Protometheus/zesty-go", "oweismo.applicant@gmail.com"}
	
	js, err := json.Marshal(personal_info)
	
	w.WriteHeader(201)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	
	port = ":" + port
	
	fmt.Println(port)

    http.HandleFunc("/", base_handler)
	http.HandleFunc("/code/challenge", endpoint_handler)
	
    http.ListenAndServe(port, nil)
}

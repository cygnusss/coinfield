package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type userProfile struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../../client/index.html")
}

func handleSignup(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		handleAllRoutes(w, r)

	case "POST":
		var profile userProfile

		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&profile)

	}
}

func main() {
	r := mux.NewRouter()

	// http.Post("")
	r.PathPrefix("/signup").HandlerFunc(handleSignup)

	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("../../../client/dist/"))))
	r.PathPrefix("/").HandlerFunc(handleAllRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4200"
	}

	log.Println("Running go app on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

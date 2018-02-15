package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../../client/index.html")
}

func handleSignup(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		handleAllRoutes(w, r)

	case "POST":
		var profile UserProfile

		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&profile)

	default:
		handleAllRoutes(w, r)
	}

}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/signup").HandlerFunc(handleSignup)

	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("../../../client/dist/"))))
	r.PathPrefix("/").HandlerFunc(handleAllRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4200"
	}

	NewClient()

	log.Println("Running go app on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

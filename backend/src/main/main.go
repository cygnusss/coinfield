package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handleAllRoutes(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../../client/index.html")
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("../../../client/dist/"))))
	r.PathPrefix("/").HandlerFunc(handleAllRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4200"
	}

	log.Println("Running go app on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

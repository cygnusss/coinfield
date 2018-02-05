package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")	
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client/dist")))
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "4200"
	}
	
	fmt.Println("listening on port :"+port)
	http.ListenAndServe(":"+port, nil)
}
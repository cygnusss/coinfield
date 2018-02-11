package main

import (
	"fmt"
	"net/http"
	// no need for html/template since serving the dist folder
	// "html/template"
	"os"
)

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("../templates/index.html")
// 	if err != nil {
// 		fmt.Fprintf(w, err.Error())
// 	}
// 	t.ExecuteTemplate(w, "index", nil)
// }

func main() {
	// http.Handle("/dist/", http.FileServer(http.Dir("../../../client/dist")))
	http.Handle("/dist", http.FileServer(http.Dir("../../../client/dist/")))
	http.HandleFunc("/", HandleAll)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4200"
	}

	fmt.Println("listening on port :" + port)
	http.ListenAndServe(":"+port, nil)
}

// HandleAll handles all routes and sends the html file
func HandleAll(w http.ResponseWriter, r *http.Request) {
	path := "../../../client/dist/index.html"
	http.ServeFile(w, r, path)
}

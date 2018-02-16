package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// HandleAllRoutes handles all routes
func HandleAllRoutes(w http.ResponseWriter, r *http.Request) {

	var storage SessionDB

	cookie, err := r.Cookie("sessionID")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "sessionID",
			Value: "0",
		}
	}

	if r.Method == "POST" && r.URL.Path == "/signup" {
		var profile UserProfile

		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&profile)

		pw := profile.Password
		un := profile.Email

		u, _ := uuid.NewV4()

		storage = SessionDB{u.String(), un}

		if pw == "password" {
			cookie = &http.Cookie{
				Name:  "sessionID",
				Value: u.String(),
			}
			w.WriteHeader(http.StatusCreated)

		} else {
			http.Redirect(w, r, "/signup", 301)
		}
	}

	if cookie.Value != "0" && r.URL.Path == "/logout" {
		cookie = &http.Cookie{
			Name:   "sessionID",
			Value:  "0",
			MaxAge: -1,
		}

		http.Redirect(w, r, "/", 200)
	}

	log.Println(storage)

	http.SetCookie(w, cookie)

	http.ServeFile(w, r, "../../../client/index.html")

}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("../../../client/dist/"))))
	r.PathPrefix("/").HandlerFunc(HandleAllRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":10443"
	}

	log.Println("Running go app on port", port)

	go http.ListenAndServe(port, http.RedirectHandler("https://localhost:"+port, 301))
	err := http.ListenAndServeTLS(port, "cert.pem", "key.pem", r)

	if err != nil {
		log.Panic(err)
	}
}

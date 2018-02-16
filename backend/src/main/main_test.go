package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	// time.Sleep(100 * time.Millisecond)
	log.Println("Should response with 200 on the home route")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{Transport: tr}
	rp, err := c.Get("https://localhost:10443/")

	if err != nil {
		t.Fatal(err)
	}

	defer rp.Body.Close()

	if rp.StatusCode != http.StatusOK {
		t.Errorf("Response code was %v; want 200", rp.StatusCode)
	}

	// body, err := ioutil.ReadAll(rp.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// expected := []byte("Hello World")

	// if bytes.Compare(expected, body) != 0 {
	// 	t.Errorf("Response body was '%v'; want '%v'", expected, body)
	// }
}

func TestServer(t *testing.T) {

	data := []byte(`{"firstname":"David","lastname":"Kumarbai","password":"password","email":"email"}`)

	r, _ := http.NewRequest("POST", "https://localhost:10443/signup", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{Transport: tr}

	rp, err := c.Do(r)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("Creates users on the signup route")

	if rp.StatusCode != http.StatusCreated {
		t.Errorf("Response code was %v; want 201", rp.StatusCode)
	}

	defer rp.Body.Close()
}

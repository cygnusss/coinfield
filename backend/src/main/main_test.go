package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	// time.Sleep(100 * time.Millisecond)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{Transport: tr}
	res, err := c.Get("https://localhost:10443/")

	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Response code was %v; want 200", res.StatusCode)
	}

	// body, err := ioutil.ReadAll(res.Body)
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
		fmt.Println(err)
	}

	defer rp.Body.Close()

	fmt.Println(rp)
}

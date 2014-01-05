// ipify-api
//
// This software implements a basic REST API that provides users with a simple
// way to query their public IP address (IPv4 or IPv6).  This code assumes that
// you are running it on Heroku's platform (https://www.heroku.com/).

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

type IPAddress struct {
	IP string `json:"ip"`
}

func jsonip(w http.ResponseWriter, r *http.Request) {
	host := net.ParseIP(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1]).String()
	jsonStr, _ := json.Marshal(IPAddress{host})

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonStr))
}

func textip(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	for k, _ := range r.Form {
		fmt.Println(k, r.Form[k][0])
	}
	fmt.Println(r.Form["hiya"][0])

	host := net.ParseIP(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1]).String()

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, host)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func main() {
	http.HandleFunc("/json", jsonip)
	http.HandleFunc("/text", textip)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
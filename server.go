package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var happy bool = true

func hostname() string {
	return os.Getenv("HOSTNAME")
}

func serverDescription() string {
	if happy {
		return "Very Happy server on host " + hostname()
	} else {
		return "Very Sad server on host " + hostname()
	}
}

func MakeSad(w http.ResponseWriter, r *http.Request) {
	response := serverDescription() + " is now sad\n"
	happy = false
	fmt.Println(response)
	io.WriteString(w, response)
	return
}

func MakeHappy(w http.ResponseWriter, r *http.Request) {
	response := serverDescription() + " is now happy\n"
	happy = true
	fmt.Println(response)
	io.WriteString(w, response)
	return
}

func Something(w http.ResponseWriter, r *http.Request) {
	if !happy {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response := serverDescription() + " handling request: " + r.URL.String() + "\n"
	fmt.Println(response)
	io.WriteString(w, response)
}

func main() {
	http.HandleFunc("/something", Something)
	http.HandleFunc("/sad", MakeSad)
	http.HandleFunc("/happy", MakeHappy)

	serverAddr := os.Getenv("SERVER_ADDR")
	fmt.Printf("Starting server on address %s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

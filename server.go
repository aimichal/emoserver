package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var happy bool = true

var hostname string = os.Getenv("HOSTNAME")
var serverAddr string = os.Getenv("SERVER_ADDR")

func serverDescription() string {
	if happy {
		return "Happy server on host " + hostname
	} else {
		return "Sad server on host " + hostname
	}
}

func getsad(w http.ResponseWriter, r *http.Request) {
	response := serverDescription() + " is now sad\n"
	happy = false
	fmt.Println(response)
	io.WriteString(w, response)
	return
}
func gethappy(w http.ResponseWriter, r *http.Request) {
	response := serverDescription() + " is now happy\n"
	happy = true
	fmt.Println(response)
	io.WriteString(w, response)
	return
}

func something(w http.ResponseWriter, r *http.Request) {
	if !happy {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response := serverDescription() + " handling request: " + r.URL.String() + "\n"
	fmt.Println(response)
	io.WriteString(w, response)
}

func main() {
	http.HandleFunc("/something", something)
	http.HandleFunc("/sad", getsad)
	http.HandleFunc("/happy", gethappy)
	fmt.Printf("Starting server on address %s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

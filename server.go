package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "Hello World!")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "OK")

}

func writeResponse(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	_, err := writer.Write(response)
	if err != nil {
		log.Println(err)
	}
}

// comment new

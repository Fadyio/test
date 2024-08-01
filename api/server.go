package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloworld)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloworld(writer http.ResponseWriter, requst *http.Request) {
	if requst.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "hello world")
}
func handleHealth(writer http.ResponseWriter, requst *http.Request) {
	if requst.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	fmt.Println("endpoint health tirggerdd")
	writeResponse(writer, "Health is ok")
}

func writeResponse(w http.ResponseWriter, r string) {
	response := []byte(r)
	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

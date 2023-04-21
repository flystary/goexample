package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//fmt.Printf("Type: %T\n", myFunc)

	//coolFunc(myFunc)
	handleRequests()
	//handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage")
}

func newEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("My New Endpoint")
	fmt.Fprintf(w, "My second endpoint")
}

func handleRequests() {
	//http.HandleFunc("/", homePage)
	http.Handle("/", isAuthorized(homePage))
	http.Handle("/new/", isAuthorized(newEndpoint))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

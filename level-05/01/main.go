/*
 * 		- Statement:
 * 			Create a simple endpoint to just return a message with the method of the Request.
 * 			The function should handle at least GET and POST.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s request handled.", strings.ToUpper(r.Method))))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRequest)
	println("HTTP server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

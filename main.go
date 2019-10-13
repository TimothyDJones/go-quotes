package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", quotes)

	fmt.Printf("Server listening on port 8080...")
	log.Panic(http.ListenAndServe(":8080", nil))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	writeResponseOrPanic(writer, "Welcome to quote API home page!")
}

func quotes(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		newQuote(writer, request)
	} else {
		writeResponseOrPanic(writer, "Invalid request method!")
	}
}

func newQuote(writer http.ResponseWriter, request *http.Request) {
	quote, err := NewQuoteFromRequest(request)
	if err != nil {
		writeResponseOrPanic(writer, fmt.Sprintf("Error: Unable to create quote from request!\nMessage: %s\n", err.Error()))
	}

	err = quote.storeInDatabase()
	if err != nil {
		writeResponseOrPanic(writer, fmt.Sprintf("Error: Unable to store quote in database!\nMessage: %s\n", err.Error()))
	}

	writeResponseOrPanic(writer, fmt.Sprintf("Quote added: \"%s\" --%s\n", quote.Quote, quote.Author))
}

// Utility function to write response using http.ResponseWriter.
// If it fails, it will panic.
func writeResponseOrPanic(writer http.ResponseWriter, message string) {
	_, err := fmt.Fprintf(writer, message)
	if err != nil {
		log.Panic(err)
	}
}

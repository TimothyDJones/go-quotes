package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type QuoteStruct struct {
	Quote string `json:"quote"`
	Author string `json:"author"`
}

// Create new quote from HTTP Request.
func NewQuoteFromRequest(request *http.Request) (*QuoteStruct, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var quote QuoteStruct
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	if quote.Quote == "" {
		return nil, errors.New("Error: Quote cannot be empty.")
	}

	return &quote, nil
}

func (quote *QuoteStruct) storeInDatabase() error {
	query := "INSERT INTO quotes (id, quote, author) VALUES (?, ?, ?)"
	_, err := ExecDB(query, nil, quote.Quote, quote.Author)

	return err
}

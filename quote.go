package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

func RandomQuoteFromDatabase() (*QuoteStruct, error) {
	fmt.Println("In 'RandomQuoteFromDatabase'...")
	query := "SELECT quote, author FROM quotes ORDER BY RANDOM() LIMIT 1"
	row, err := QueryDB(query)
	if err != nil {
		return nil, err
	}

	fmt.Println("In 'RandomQuoteFromDatabase' after query...")

	if !row.Next() {
		return nil, errors.New("Error: No quote found in database!")
	}

	fmt.Println("In 'RandomQuoteFromDatabase' after reading row...")

	var (
		quote string
		author string
	)
	err = row.Scan(&quote, &author)
	if err != nil {
		return nil, err
	}
	log.Println(quote, author)

	fmt.Println("In 'RandomQuoteFromDatabase' after scanning 'quote' and ''author'...")

	quoteStruct := &QuoteStruct{
		Quote: quote,
		Author: author,	
	}

	return quoteStruct, nil
}

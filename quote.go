package main

import (
	"encoding/json"
	"errors"
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

	query := "SELECT quote, author FROM quotes ORDER BY RANDOM() LIMIT 1"
	row, err := QueryDB(query)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, errors.New("Error: No quote found in database!")
	}

	var (
		quote string
		author string
	)
	err = row.Scan(&quote, &author)
	if err != nil {
		return nil, err
	}
	log.Println(quote, author)

	quoteStruct := &QuoteStruct{
		Quote: quote,
		Author: author,	
	}

	return quoteStruct, nil
}

func GetQuoteCountFromDatabase() (int, error) {
	query := "SELECT COUNT(1) FROM quotes"
	row, err := QueryDB(query)
	if err != nil {
		return 0, err
	}

	if !row.Next() {
		return 0, errors.New("Error: Unable to get count of quotes from database!")
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	log.Println(count)

	return count, nil
}

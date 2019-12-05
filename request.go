package crawler

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Request custom errors
var (
	ErrNotOKResponse    = errors.New("response status is not OK")
	ErrUnableToReadBody = errors.New("could not read the response body")
)

// Request ...
type Request struct {
	URL string
}

// Call makes request call returns goquery.Document
func (r Request) Call() (*goquery.Document, error) {
	res, err := http.Get(r.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, ErrNotOKResponse
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, ErrUnableToReadBody
	}
	return doc, nil
}

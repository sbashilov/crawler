package test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sbashilov/crawler"
)

func TestCrawler_Crawl(t *testing.T) {
	ts := runTestServer()
	c, err := crawler.New(ts.URL)
	require.Nil(t, err)
	p, err := c.Crawl()
	require.Nil(t, err)
	require.NotNil(t, p.CrossPages)
	require.Equal(t, ts.URL, p.Link)
	require.Equal(t, 2, len(p.CrossPages))
	ts.Close()
}

func runTestServer() *httptest.Server {
	tmplMain := template.Must(template.ParseFiles("layout/index.html"))
	tmplContacts := template.Must(template.ParseFiles("layout/contacts.html"))
	tmplFaq := template.Must(template.ParseFiles("layout/faq.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplMain.Execute(w, nil)
	})
	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		tmplContacts.Execute(w, nil)
	})
	mux.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		tmplFaq.Execute(w, nil)
	})

	ts := httptest.NewServer(mux)
	return ts
}

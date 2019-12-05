package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmplMain := template.Must(template.ParseFiles("index.html"))
	tmplArticles := template.Must(template.ParseFiles("articles.html"))
	tmplContacts := template.Must(template.ParseFiles("contacts.html"))
	tmplFaq := template.Must(template.ParseFiles("faq.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplMain.Execute(w, nil)
	})

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		tmplArticles.Execute(w, nil)
	})

	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		tmplContacts.Execute(w, nil)
	})

	http.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		tmplFaq.Execute(w, nil)
	})

	http.ListenAndServe(":80", nil)
}

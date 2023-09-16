package main

import (
	"log"
	"net/http"
	"text/template"
)

const Portnumber = ":8001"

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")

}

func main() {

	http.HandleFunc("/", homePage)

	fs := http.FileServer(http.Dir("img/"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	log.Println("Starting WebServer on Port" + Portnumber)
	http.ListenAndServe(Portnumber, nil)

}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}

}

package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type Page struct {
	Name string
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
func index(writer http.ResponseWriter, request *http.Request) {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	p := Page{"GopherBrasil"}

	if name := request.FormValue("name"); name != "" {
		p.Name = name
	}

	if err := templates.ExecuteTemplate(writer, "index.html", p); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

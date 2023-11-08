package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	switch r.Method {
	case "post", "POST":
		fmt.Println("r:", r.Method)
	}
	tmpl.Execute(w, nil)
}

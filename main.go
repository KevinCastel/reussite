package main

import (
	"fmt"
	"net/http"
	reussite "reussite/utils"
)

const PORT = "8020"

func main() {
	fmt.Println("Access to the web server by this one: http://localhost:" + PORT + "/home")
	fileServerCss := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServerCss))

	http.HandleFunc("/home", reussite.MainPageHandler)
	http.ListenAndServe("localhost:"+PORT, nil)
}

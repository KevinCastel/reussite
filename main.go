package main

import (
	"fmt"
	"net/http"
	Reussite "reussite/utils"
)

const PORT = "8020"

func main() {
	fmt.Println("Access to the web server by this one: http://localhost:" + PORT + "/home")

	http.HandleFunc("/home", Reussite.mainPage)
	http.ListenAndServe("localhost:"+PORT, nil)
}

package main

import (
	"fmt"
	"net/http"
	reussite "reussite/utils"
)

const PORT = "8020"

func init() {
	go reussite.GetExercices()
}

func main() {

	fmt.Println("Access to the web server by this one: http://localhost:" + PORT + "/home")
	fileServerCss := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServerCss))

	http.HandleFunc("/home", reussite.MainPageHandler)
	http.ListenAndServe("localhost:"+PORT, nil)
}

/*
   exercises := map[int][]string{
       1: []string{"displayfirstparam", "displaylastparam", "displayz", "displaya", "hello", "onlya", "onlyz", "printdigits", "strlen", "paramcount", "displayalpham", "displayalrevm", "countdown"},
       2: []string{"nrune", "lastrune", "printstr", "strrev", "firstrune", "printreversealphabet", "printalphabet", "wdmatch", "ispowerof2", "rot13", "lastword", "rot14", "max", "reduceint"},
       3: []string{"switchcase", "swap", "compare", "expandstr", "tabmult", "searchreplace", "alphamirror", "doop", "findprevprime", "reversebits", "chunk", "foldint"},
       4: []string{"swapbits", "capitalize", "repeatalpha", "atoi", "reversestrcap", "printbits", "inter", "union", "piglatin", "romannumbers"},
       5: []string{"firstword", "sortwordarr", "cleanstr", "addprimesum", "printhex", "gcd"},
       6: []string{"printcomb", "split", "hiddenp", "rostring", "revwstr", "reverserange", "range", "slice"},
       7: []string{"itoa", "atoibase", "foreach", "fprime", "printrevcomb"},
       8: []string{"listsize", "rpncalc", "brackets", "options", "grouping", "printmemory"},
       9: []string{"listremoveif", "itoabase", "brainfuck"},
   }*/

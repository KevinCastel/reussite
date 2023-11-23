package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const URL string = "https://github.com/01-edu/public/projects?type=classic"

func GetExercices() {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Le lien n'a pas réussis à être ouvert :", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Une erreur est survenue:", err)
	}

	arrayLink := GetExercicesLink(body)

	for _, link := range arrayLink {
		resp, err := http.Get(link)

		if err == nil {

			defer resp.Body.Close()

			fmt.Println("link :", link)

			if strings.HasSuffix(link, "displayfirstparam") {
				_, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Impossible de lire le contenu de la page :", err)
					os.Exit(1)
				}
			}
		}

	}
}

/*
Iterate body and get all sub link for accessing to the exercices
*/
func GetExercicesLink(bodyBuff []byte) []string {
	arrayLinks := make([]string, 0)

	WriteLog(string(bodyBuff))
	body := strings.Split(string(bodyBuff), "\n")
	var lineFormatted string
	for _, line := range body {
		lineFormatted = strings.TrimSpace(line)
		if strings.HasPrefix(lineFormatted, "<li>") {

			regexObj, _ := regexp.Compile(`<li>\(\d*\)\s*<a\s*href="(?P<link>[^"]*)">.*<\/a>`)
			match := regexObj.FindStringSubmatch(lineFormatted)
			indexMatch := regexObj.SubexpIndex("link")
			if len(match) != 0 {
				arrayLinks = append(arrayLinks, match[indexMatch])
			}

			arrayLinks = append(arrayLinks, lineFormatted)
		}
	}

	return arrayLinks
}

/*
Called for getting childrens elements as string
"Box-sc-g0xbh4-0 bJMeLZ js-snippet-clipboard-copy-unpositioned"
*/
func GetChildrensElementsByClass(parentElementClassName, body string) []string {
	for _, line := range strings.Split(body, "\n") {
		regexObj, _ := regexp.Compile(`(.*)(div class=)(?P<class_name>\"Box-sc-g0xbh4-0 bJMeLZ js-snippet-clipboard-copy-unpositioned")(.*)`)
		match := regexObj.FindStringSubmatch(line)
		index := regexObj.SubexpIndex("class_name")
		result := match[index]
		fmt.Println("result:", result)

	}

	return make([]string, 0)
}

/*
Called for writing logs
*/
func WriteLog(content string) {

	filePath, errPath := os.Getwd()
	if errPath != nil {
		fmt.Println("Une erreur est survenue pour la récupération du chemin :", errPath)
	}

	err := os.WriteFile(filePath+"/webpage.html", []byte(content), 0777)
	if err != nil {
		fmt.Println("L'écriture est impossible :", err)
	}
	fmt.Println("Les logs ont bien été ajouter")
}

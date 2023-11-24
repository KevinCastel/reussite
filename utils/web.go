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

// Called for getting the exercices datas
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
		link = strings.ReplaceAll(strings.TrimSpace(link), "tree/", "")
		link = strings.ReplaceAll(link, "public/", "public/blob/")
		resp, err := http.Get(link + "/README.md?plain=1")
		if err == nil {

			defer resp.Body.Close()

			if strings.HasSuffix(link, "displayfirstparam") {
				fmt.Println("ouverture du lien :\"", link+"/README.md?plain=1\"")
				bodyExercice, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Impossible de lire le contenu de la page :", err)
					os.Exit(1)
				}
				reactAppNode := GetChildrensByTagName("react-app", string(bodyExercice))
				scriptNode := GetChildrensByTagName("script", reactAppNode)

				arrayTxt := GetExercicesData(scriptNode)
				FormatExercices(arrayTxt)

			}
		}
	}
}

// Called for getting exercices informations in an array of strings so get
// rule and output
func GetExercicesData(content string) []string {
	pattern := `(\"blob\"\:)\{(\"[\w|]*\":)\[(?P<content>[^]]*)`
	regexObj, _ := regexp.Compile(pattern)
	match := regexObj.FindStringSubmatch(content)
	mIndex := regexObj.SubexpIndex("content")
	result := make([]string, 0)

	for _, e := range strings.Split(match[mIndex], ",") {
		if len(strings.TrimSpace(e)) > 2 && !strings.HasPrefix(e, "\"```") {
			result = append(result, e)
		}
	}
	return result
}

/*
Iterate body and get all sub link for accessing to the exercices
*/
func GetExercicesLink(bodyBuff []byte) []string {
	arrayLinks := make([]string, 0)

	//WriteLog(string(bodyBuff))
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

		}
	}

	return arrayLinks
}

/*
Called for getting all childrens of an parent by his id
*/
func GetChildrendsById(id, body string) string {
	body = strings.TrimSpace(strings.ReplaceAll(body, "\n", ""))

	var result string

	//regexObj, _ := regexp.Compile(`(.*)(\<react-app)(?P<tag_childrens>.*)(\</react-app>)(.*)`)
	regexObj, _ := regexp.Compile(`(.*)(\<` + id + `)(?P<tag_childrens>.*)(\</` + id + `>)(.*)`)
	match := regexObj.FindStringSubmatch(body)
	index := regexObj.SubexpIndex("tag_childrens")
	if len(match) != 0 {
		result = match[index]
	}

	return result

}

/*
Called for getting childrens elements as string by the tag name
Take args as:

	tagName string is the tagname
	body string is the web page
*/
func GetChildrensByTagName(tagName string, body string) string {
	body = strings.TrimSpace(strings.ReplaceAll(body, "\n", ""))

	var result string

	//regexObj, _ := regexp.Compile(`(.*)(\<react-app)(?P<tag_childrens>.*)(\</react-app>)(.*)`)
	regexObj, _ := regexp.Compile(`(.*)(\<` + tagName + `)(?P<tag_childrens>.*)(\</` + tagName + `>)(.*)`)
	match := regexObj.FindStringSubmatch(body)
	index := regexObj.SubexpIndex("tag_childrens")
	if len(match) != 0 {
		result = match[index]
	}

	return result
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

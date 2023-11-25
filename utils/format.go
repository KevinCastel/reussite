package utils

import (
	"fmt"
	"regexp"
)

// Called for formatting each string elements of an array,
// The array passed as argument have to contains "rules" and "output"
// So return-it as an array of string
func FormatExercices(exercicesContents StringSlice, link string) []string {
	var typeLine string
	var actualType string

	mapPattern := map[string]string{
		"(.*)(Instruction)(.*)": "instruction",
		"(.*)(Usage)(.*)":       "usage",
	}

	mapResultArray := make(map[string][]string, 0)

	for _, line := range exercicesContents {
		for pattern, patternType := range mapPattern {
			actualType = GetLineType(line, pattern, patternType)
			if actualType != "" {
				typeLine = actualType
			}
		}
		if typeLine == "instruction" || typeLine == "usage" {
			mapResultArray[typeLine] = append(mapResultArray[typeLine], line)
		}
	}

	exercicesContents.Clear()

	exercicesContents = append(exercicesContents, GetExerciceName(link))

	var result string
	var lastType string

	for currentType, arrayContent := range mapResultArray {
		for indexLine, line := range arrayContent {
			if indexLine > 0 && IsStringMatched(line, "(.*)(Instructions|Usage)(.*)") {
				result += line + "\n"
			}
		}

		if lastType != currentType {
			lastType = currentType
			exercicesContents = append(exercicesContents, result)
			result = ""
		}
	}

	return exercicesContents
}

// Called for getting if the string matched or not
func IsStringMatched(s, p string) bool {
	regexpObj, _ := regexp.Compile("")
	match := regexpObj.FindStringSubmatch(s)

	return (len(match) > 0)
}

// Called for getting an line type(using regex)
func GetLineType(s, pattern, typeLine string) string {
	regexObj, _ := regexp.Compile(pattern)
	match := regexObj.FindStringSubmatch(s)
	if len(match) > 0 {
		return typeLine
	} else {
		return ""
	}

}

// Called for getting an exercice name from an link by using the regex tools
func GetExerciceName(link string) string {
	var pattern string = "(https://github.com/01-edu/public/blob/master/subjects/)(?P<exerciceName>.*)((/README.md?plain=1)*)"
	fmt.Println("link:", link)
	regexObj, _ := regexp.Compile(pattern)
	match := regexObj.FindStringSubmatch(link)
	if len(match) > 0 {

		matchIndex := regexObj.SubexpIndex("exerciceName")
		return match[matchIndex]
	}

	return ""
}

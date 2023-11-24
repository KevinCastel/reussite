package utils

import (
	"fmt"
	"regexp"
)

// Called for formatting each string elements of an array,
// The array passed as argument have to contains "rules" and "output"
func FormatExercices(a []string) {
	var typeLine string
	var actualType string

	mapPattern := map[string]string{
		"(.*)(Instruction)(.*)": "instruction",
		"(.*)(Usage)(.*)":       "usage",
	}

	for _, line := range a {
		for pattern, patternType := range mapPattern {
			actualType = GetLineType(line, pattern, patternType)
			if actualType != "" {
				typeLine = actualType
			}
		}

		if typeLine == "instruction" {

		} else if typeLine == "usage" {

		}

	}
	fmt.Println(a)
}

// Called for regex
func GetLineType(s, pattern, typeLine string) string {
	regexObj, _ := regexp.Compile(pattern)
	match := regexObj.FindStringSubmatch(s)
	if len(match) > 0 {
		return typeLine
	} else {
		return ""
	}

}

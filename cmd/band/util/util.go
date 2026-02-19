package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetGoModule(dir string) string {
	file, err := os.Open(dir + "go.mod")
	if err != nil {
		log.Printf("Error when opening file: %s", err)
		return ""
	}
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		fmt.Println()
		arr := strings.Split(text, " ")
		return arr[1]
	}
	return ""
}

// FirstUpper
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func UnderscoreToCamelCase(input string) string {

	parts := strings.Split(input, "_")

	var result strings.Builder

	for _, part := range parts {

		if len(part) > 0 {

			result.WriteString(strings.ToUpper(string(part[0])) + strings.ToLower(part[1:]))

		}

	}

	return result.String()
}

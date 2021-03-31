package utils

import (
	"os"
	"strings"
)

// CheckFile checks if file exist
// , if not exist it creates the file
func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReplaceURLSpaces removes any space from string
// and add %20
func ReplaceURLSpaces(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

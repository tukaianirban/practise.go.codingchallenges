package grepfilecontent

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// returns the line number at which the string is matched
func findStringInFile(filename string, searchString string) []int {

	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	foundLines := make([]int, 0)

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {

		line := scanner.Text()
		// log.Printf("searching for string in line=%s", line)
		if strings.Contains(line, searchString) {
			foundLines = append(foundLines, lineNumber)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error reading from file=%s", err.Error())
	}

	return foundLines
}

// traverse throguh the directoru structure and run find func in each file
func traverseFolders(pathName string, searchString string) {

	fileItems, err := ioutil.ReadDir(pathName)
	if err != nil {
		return
	}

	for _, item := range fileItems {

		// find in files
		if !item.IsDir() {
			foundLines := findStringInFile(pathName+item.Name(), searchString)
			log.Printf("File: %s found matches in lines=%#v", item.Name(), foundLines)
		} else {
			traverseFolders(pathName+item.Name()+"/", searchString)
		}
	}
}

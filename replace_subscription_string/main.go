package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"regexp"
)

const (
	folderPath = ""

	// Logic app ID (under construction)
	oldLa = `(?P<first>LogicApp_+[a-zA-Z]+_*[a-zA-z]+_)+(?P<last>[a-zA-Z0-9])+\b`
	newLa = `\1`

	// Subscription id i.e xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	oldSub = `([^/])+([0-9])+(-)+([a-zA-Z0-9])+(-)+([a-zA-Z0-9])+(-)+([a-zA-Z0-9])+(-)+([a-zA-Z0-9])+([$/])`
	newSub = `` + "/"
)

func main() {
	fmt.Print("Starting application...\n")

	re := regexp.MustCompile(oldSub)

	files, err := checkDir(folderPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		hasMatch, err := checkForStringMatches(file.Name())

		if err != nil {
			log.Fatal(err)
		}
		
		if hasMatch == nil {
			continue
		}

		err = replaceString(hasMatch, re, newSub)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("finished file: ", file.Name())
	}
}

func checkDir(folderPath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(folderPath)
	
	// Check if directory exists
	if err != nil {
		return nil, errors.New("Unable to find directory. Does not exist or was denied access")
	}

	// Error if there are no files in dir
	if len(files) == 0 {
		return nil, errors.New("there are no files in this directory")
	}

	return files, nil
}

func checkForStringMatches(file string) (*string, error) {

	content, err := ioutil.ReadFile(path.Join(folderPath, file))

	if err != nil {
		return nil, errors.New("Error reading content")
	}

	matched, err := regexp.Match(oldSub, content)

	if err != nil {
		return nil, errors.New("Error matching regexp with file content")
	}
	
	if matched != true {
		return nil, nil
	}

	return &file, nil
}

func replaceString(file *string, re *regexp.Regexp, repl string) error {

	fullPath := path.Join(folderPath, *file)
	content, err := ioutil.ReadFile(fullPath)

	if err != nil {
		return errors.New("Error reading content")
	}

	newContent := re.ReplaceAll(content, []byte(repl))

	err = ioutil.WriteFile(fullPath, newContent, 0777)

	return err
}
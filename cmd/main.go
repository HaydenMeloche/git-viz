package main

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]

	filePath, email, e := validateArguments(args)

	if e != nil {
		print(e.Error())
		return
	}

	folders := findGitFolders(filePath)
	stats(folders, email)
}

func findGitFolders(info string) []string {
	var gitLocations []string
	items, _ := ioutil.ReadDir(info)
	for _, item := range items {
		if item.IsDir() && item.Name() == ".git" {
			gitLocations = append(gitLocations, info + "/" + item.Name())
			continue
		}
		if item.IsDir() {
			gitLocations = append(gitLocations, findGitFolders(info + "/" + item.Name())...)
		}
	}
	return gitLocations
}

func validateArguments(args []string) (string, string, error) {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(args) != 2 {
		return "", "", errors.New("exactly two arguments must be provided")
	}

	_, e := os.Stat(args[0])

	if e != nil {
		return "", "", e
	}

	if !rxEmail.MatchString(args[1]) {
		return "", "", errors.New("2nd argument must be a valid email address")
	}

	return args[0], args[1], nil
}
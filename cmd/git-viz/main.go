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

	printFolders(filePath)
	print(email)

}

func printFolders(info string) () {
	items, _ := ioutil.ReadDir(info)
	for _, item := range items {
		print(item.Name())
	}
}

func validateArguments(args []string) (string, string, error) {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(args) != 2 {
		return "", "", errors.New("exactly two arguments must be provided")
	}

	_, error := os.Stat(args[0])

	if error != nil {
		return "", "", error
	}

	if !rxEmail.MatchString(args[1]) {
		return "", "", errors.New("2nd argument must be a valid email address")
	}

	return args[0], args[1], nil
}
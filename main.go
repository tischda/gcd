package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println(injectSlashDIfneeded(os.Args[1], cwd))
}

func injectSlashDIfneeded(path, cwd string) string {
	matched, err := regexp.MatchString("^[a-zA-Z]:", path)
	if err != nil {
		log.Fatalln(err)
	}
	if matched {
		a := strings.ToLower(cwd)[0]
		b := strings.ToLower(path)[0]
		if a != b {
			return ("/d " + path)
		}
	}
	return path
}

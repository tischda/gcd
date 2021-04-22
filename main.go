package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// no path given, cd alias should execute the command 'cd' which prints cwd
		// the for loop in the calling script must return a token to execute 'cd <space>'
		fmt.Println(" ")
	} else {
		cwd, _ := os.Getwd()
		path := strings.Join(os.Args[1:], " ")
		fmt.Println(addSwitchIfNeeded(path, cwd))
	}
}

// Use the /D switch to change current drive in addition to changing current
// directory for a drive.
func addSwitchIfNeeded(path, cwd string) string {
	matched, _ := regexp.MatchString("^[a-zA-Z]:", path)
	if matched {
		a := strings.ToLower(cwd)[0]
		b := strings.ToLower(path)[0]
		if a != b {
			return ("/d " + path)
		}
	}
	return path
}

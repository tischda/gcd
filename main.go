package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// no path given, cd alias should execute the command [cd] which prints cwd
		fmt.Println(" ") // so that for loop returns a token (cf. README.md) and executes cd <space>
	} else {
		cwd, _ := os.Getwd()
		fmt.Println(addSwitchIfNeeded(os.Args[1], cwd))
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

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	if len(os.Args) < 2 {
		// no path given, cd alias should execute the command 'cd' which prints cwd
		// the for loop in the calling script must return a token to execute 'cd <space>'
		fmt.Println(" ")
	} else {
		cwd, _ := os.Getwd()
		path := strings.Join(os.Args[1:], " ")
		param := addSwitchIfNeeded(path, cwd)
		printEncodeParameter(param)
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

// golang outputs UTF-8. This will not work on a standard Windows console.
func printEncodeParameter(param string) {
	var t *charmap.Charmap
	switch cp := GetConsoleCP(); cp {
	case uint32(437):
		t = charmap.CodePage437
	case uint32(850):
		t = charmap.CodePage850
	case uint32(1252):
		t = charmap.Windows1252
	default:
		fmt.Println(param)
		return
	}
	output, _, err := transform.String(t.NewEncoder(), param)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

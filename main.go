package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

var version string
var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Printf("gcd %s - serves arguments to a \"cd\" command alias\n", version)
	} else if len(os.Args) < 2 {
		// no argument given, cd alias should execute the command 'cd' which prints cwd
		// the for loop in the calling script must return a token to execute 'cd <space>'
		fmt.Println(" ")
	} else {
		out := processArgs(os.Args[1:])
		fmt.Println(out)
	}
}

func processArgs(args []string) string {
	cwd, _ := os.Getwd()
	path := strings.Join(args, " ")

	// remove quotes: "Program Files (x86)" --> Program Files (x86)
	// quotes will be re-added later so we have better control of what is quoted.
	path = strings.ReplaceAll(path, "\"", "")

	cmd := addSwitchIfNeeded(path, cwd)

	return translateUTF8ToCodepage(cmd)
}

// Use the /D switch to change current drive in addition to changing current
// directory for a drive.
func addSwitchIfNeeded(path, cwd string) string {
	quotedPath := fmt.Sprintf("\"%s\"", path)
	matched, _ := regexp.MatchString("^[a-zA-Z]:", path)
	if matched {
		a := strings.ToLower(cwd)[0]
		b := strings.ToLower(path)[0]
		if a != b {
			return ("/d " + quotedPath)
		}
	}
	return quotedPath
}

// Translate UTF-8 to console codepage (golang outputs UTF-8, this will not work on a standard Windows console).
func translateUTF8ToCodepage(param string) string {
	var t *charmap.Charmap
	switch cp := GetConsoleCP(); cp {
	case uint32(437):
		t = charmap.CodePage437
	case uint32(850):
		t = charmap.CodePage850
	case uint32(1252):
		t = charmap.Windows1252
	default:
		return param

	}
	out, _, err := transform.String(t.NewEncoder(), param)
	if err != nil {
		panic(err)
	}
	return out
}

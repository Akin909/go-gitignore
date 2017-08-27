// Package main provides a gitignore file prepopulated with files to be ignored
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func exists(filePath string) (exists bool) {
	exists = true

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exists = false
		return exists
	}
	return exists
}

func main() {
	fg := flag.String("ft", "node", "Set the desired file type to generate a git ignore for")
	rm := flag.Bool("rm", false, "Empty the gitignore file")
	flag.Parse()
	args := flag.Args()

	if *rm == true {
		err := ioutil.WriteFile(".gitignore", []byte(""), 0)
		check(err)
	}

	bool := exists("./.gitignore")
	if bool == false {

		f, err := os.Create(".gitignore")
		check(err)
		writeGitignore(f, fg, args)

	} else {
		f, err := os.OpenFile("./.gitignore", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		check(err)
		writeGitignore(f, fg, args)
		defer f.Close()
	}

}

var elmGitIgnore = `
### Elm template
# elm-package generated files
elm-stuff/
# elm-repl generated files
repl-temp-*
### Node template
# Logs
logs
*.log
npm-debug.log*

# Runtime data
pids
*.pid
*.seed`

func writeGitignore(file *os.File, flag *string, args []string) {

	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("> Selected file type is: ->", *flag)
	c.Println("> Files to be added to gitignore: ->", args)

	switch *flag {
	case "node":
		git, err := ioutil.ReadFile("./node_gitignore")
		check(err)
		s := string(git)
		_, err := file.WriteString(s)
		check(err)
	case "go":
		_, err := file.WriteString("gin_bin\n")
		check(err)
	case "elm":
		_, err := file.WriteString(elmGitIgnore + "\n")
		check(err)
	}

	for _, word := range args {
		_, err := file.WriteString("\n" + word)
		check(err)
	}

}

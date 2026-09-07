package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"flag"
	"strings"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content string
}

func main() {
	// Read the command line and define a flag
	filePtr := flag.String("file", "first-post.txt", "text file to render")
	flag.Parse()

	// Read file contents 
	fileContents, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		panic(err)
	}

	// Create the page 
	page := Page{
		TextFilePath: *filePtr,
		TextFileName: *filePtr,
		HTMLPagePath: strings.TrimSuffix(*filePtr, ".txt") + ".html",
		Content: string(fileContents),
	}

	// Save a new template to memory by parsing the given template
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Print the result to stdout
	if err := t.Execute(os.Stdout, page); err != nil {
		panic(err)
	}
	
	// Create a new file for the template the write to
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	// Write the template to the new file
	if err := t.Execute(newFile, page); err != nil {
		panic(err)
	}
}

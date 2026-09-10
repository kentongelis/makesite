package main

import (
	"github.com/gomarkdown/markdown"
	"html/template"
	"io/ioutil"
	"os"
	"flag"
	"strings"
	"log"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content template.HTML
}

func main() {
	// Read the command line and define a flag
	// filePtr := flag.String("file", "first-post.txt", "text file to render")
	// flag.Parse()

	// Read the command line and define a directory flag
	dirPtr := flag.String("dir", ".", "directory to render files")
	flag.Parse()

	// Get all files in given directory
	files, err := ioutil.ReadDir(*dirPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Loop over all files and skip over any file that doesn't end in .txt
	for _, file := range files {
		filename := file.Name()
		if file.IsDir() || !strings.HasSuffix(filename, ".txt") {
			continue
		}

		// Read file contents
		fileContents, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		// Parse the Markdown content into HTML
		htmlContent := markdown.ToHTML(fileContents, nil, nil)

		// Create the page
		page := Page{
			TextFilePath: filename,
			TextFileName: filename,
			HTMLPagePath: strings.TrimSuffix(filename, ".txt") + ".html",
			Content: template.HTML(htmlContent),
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

		// Write the template to the new file
		if err := t.Execute(newFile, page); err != nil {
			panic(err)
		}
		newFile.Close()
	}
}

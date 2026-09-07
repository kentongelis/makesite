package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content string
}

func main() {
	// Read file contents 
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	page := Page{
		TextFilePath: "first-post.txt",
		TextFileName: "first-post.txt",
		HTMLPagePath: "first-post.html",
		Content: string(fileContents),
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	if err := t.Execute(os.Stdout, page); err != nil {
		panic(err)
	}
	
	newFile, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	if err := t.Execute(newFile, page); err != nil {
		panic(err)
	}
}

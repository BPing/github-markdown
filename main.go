package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

func main() {

	var input []byte
	var output []byte
	var err error

	var filePath = "README.md"
	var fileOutPath = "output.html"

	commonHtmlFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES|
		blackfriday.HTML_COMPLETE_PAGE

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	renderer := blackfriday.HtmlRenderer(commonHtmlFlags, "cbping", "github-markdown.css")

	if input, err = ioutil.ReadFile(filePath); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from", filePath, ":", err)
		os.Exit(-1)
	}

	output = blackfriday.Markdown(input, renderer, extensions)
	//output=github_flavored_markdown.Markdown(output)


	// output the result
	var out *os.File

	if out, err = os.Create(fileOutPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating %s: %v", fileOutPath, err)
		os.Exit(-1)
	}
	defer out.Close()

	if _, err = out.Write(output); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing output:", err)
		os.Exit(-1)
	}
}
package main

import (
	"fmt"
	"log"
	"os"
)

func saveFileToGorMD(title string, markdown string, slug string, date string, tags []string) {
	var format string
	format = "---\ndate: %s\nlayout: post\ntitle: %s\npermalink: '%s'\ncategories:\n%s\ntags:\n%s\n---\n\n"
	var filePath string
	filePath = "dist/" + slug + ".md"
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var tagg string
	for _, t := range tags {
		tagg = tagg + "- " + t + "\n"
	}
	f.WriteString(fmt.Sprintf(format, date, title, slug, "- 随记", tagg))
	f.WriteString(markdown)

}

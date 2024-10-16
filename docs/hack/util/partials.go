package util

import (
	"fmt"
	"path/filepath"
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var partialNameRegex = regexp.MustCompile(`(\..*$)|[^a-zA-Z]`)

func GetPartialImportName(partialImport string) string {
	basename := filepath.Base(partialImport)
	basename = partialNameRegex.ReplaceAllString(basename, "")

	caser := cases.Title(language.English, cases.NoLower)
	return fmt.Sprintf("Partial%s", caser.String(basename))
}

func GetPartialImport(partialFile, importingFile string) string {
	partialImportPath, err := filepath.Rel(filepath.Dir(importingFile), partialFile)
	if err != nil {
		panic(err)
	}

	if partialImportPath[0:1] != "." {
		partialImportPath = "./" + partialImportPath
	}

	return fmt.Sprintf(TemplatePartialImport, GetPartialImportName(partialFile), partialImportPath)
}

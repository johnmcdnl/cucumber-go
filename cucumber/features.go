package cucumber

import (
	"github.com/cucumber/gherkin-go"
	"path/filepath"
	"os"
	"strings"
)

const (
	featureExtension = ".feature"
)

func Features(path string) ([]*gherkin.GherkinDocument, error) {
	var gherkinDocuments []*gherkin.GherkinDocument

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, featureExtension) {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			gd, err := gherkin.ParseGherkinDocument(f)
			if err != nil {
				return err
			}
			gherkinDocuments = append(gherkinDocuments, gd)
		}
		return err
	})

	return gherkinDocuments, err
}

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
	var documents []*gherkin.GherkinDocument

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, featureExtension) {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			document, err := gherkin.ParseGherkinDocument(f)
			if err != nil {
				return err
			}
			documents = append(documents, document)
		}
		return err
	})

	return documents, err
}

func RunFeature(document *gherkin.GherkinDocument) {

}

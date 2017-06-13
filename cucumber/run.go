package cucumber

import (
	"fmt"
)

func Run(path string) error {
	gherkinDocuments, err := Features(path)
	if err != nil {
		return err
	}

	for _, gherkinDocument := range gherkinDocuments {
		fmt.Println(gherkinDocument.Feature.Name)
	}
	return nil

}

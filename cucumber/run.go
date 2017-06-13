package cucumber

func Run(path string) error {
	documents, err := Features(path)

	if err != nil {
		return err
	}

	for _, document := range documents {
		RunFeature(document)
	}

	return nil

}

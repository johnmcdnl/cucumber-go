package cucumber

import (
	"context"
	"sync"
)

func Run(ctx context.Context, path string) error {
	documents, err := Features(path)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, document := range documents {
		wg.Add(1)
		go RunFeature(ctx, &wg, document)
	}
	wg.Wait()

	return nil

}

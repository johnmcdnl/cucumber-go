package cucumber

import (
	"github.com/cucumber/gherkin-go"
	"fmt"
	"context"
	"sync"
)

func RunPickle(ctx context.Context, pickleWg *sync.WaitGroup, pickle *gherkin.Pickle) {
	defer pickleWg.Done()
	fmt.Println(pickle.Name)

}

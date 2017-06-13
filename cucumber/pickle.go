package cucumber

import (
	"github.com/cucumber/gherkin-go"
	"fmt"
	"context"
	"sync"
	"time"
	"math/rand"
)

func RunPickle(ctx context.Context, pickleWg *sync.WaitGroup, pickle *gherkin.Pickle) {
	defer pickleWg.Done()

	for _, step :=range pickle.Steps{
		ctx = context.WithValue(ctx, "scenarioName", pickle.Name)
		RunPickleStep(ctx, step)
	}
}

func RunPickleStep(ctx context.Context, step *gherkin.PickleStep) {
	scenarioName := ctx.Value("scenarioName").(string)
	fmt.Println(scenarioName, step.Text)
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
}


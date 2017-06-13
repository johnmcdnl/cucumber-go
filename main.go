package main

import (
	"github.com/johnmcdnl/cucumber-go/cucumber"
	"context"
)

func main() {
	if err := cucumber.Run(context.TODO(), "./internal"); err != nil {
		panic(err)
	}
}

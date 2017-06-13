package main

import (
	"github.com/johnmcdnl/cucumber-go/cucumber"
)

func main() {
	if err := cucumber.Run("./"); err != nil {
		panic(err)
	}
}

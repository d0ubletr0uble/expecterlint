package main

import (
	"github.com/d0ubletr0uble/expecterlint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(expecterlint.Analyzer)
}

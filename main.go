package main

import (
	"fmt"
)

func main() {

	var options Options = readArguments()
	var extraSteps int = 1
	if options.genIndexPage { extraSteps +=1 }

	for i, e := range options.paths {
		fmt.Printf("[%v/%v] Processing file %v", (i+1), len(options.paths)+extraSteps, e)
		fmt.Printf(" (%.1fms)\n", timing(func(){genFile(e, options)}))
	}

	fmt.Printf("[%v/%v] Generating styles ", len(options.paths)+1, len(options.paths)+extraSteps)
	fmt.Printf(" (%.1fms)\n", timing(func(){genCss(options)}))
	if options.genIndexPage {
		fmt.Printf("[%v/%v] Generating styles ", len(options.paths)+2, len(options.paths)+extraSteps)
		fmt.Printf(" (%.1fms)\n", timing(func(){genIndexPage(options)}))
	}
	fmt.Println("[SUCCESS] Processing has ended successfully.")
}


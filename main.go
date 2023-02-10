package main

import (
	"fmt"
	"workspace/gocolor"
	"workspace/goprocessing"
)

func main() {
	fmt.Printf(gocolor.Green+"[SUCCESS]"+gocolor.Reset+" All files has been processed "+gocolor.Yellow+"(Total: %.1f ms)\n"+gocolor.Reset, goprocessing.Timing(start))
}

func start() {
	var options goprocessing.Options = goprocessing.ReadArguments()
	var extraSteps int = 1
	if options.GenIndexPage { extraSteps +=1 }

	for i, e := range options.Paths {
		fmt.Printf(gocolor.Blue+"[%v/%v]"+gocolor.Reset+" Processing file %v", (i+1), len(options.Paths)+extraSteps, e)
		fmt.Printf(gocolor.Yellow+" (%.1f ms)\n"+gocolor.Reset, goprocessing.Timing(func(){goprocessing.GenFile(e, options)}))
	}

	fmt.Printf(gocolor.Blue+"[%v/%v]"+gocolor.Reset+" Generating styles", len(options.Paths)+1, len(options.Paths)+extraSteps)
	fmt.Printf(gocolor.Yellow+" (%.1f ms)\n"+gocolor.Reset, goprocessing.Timing(func(){goprocessing.GenCss(options)}))
	if options.GenIndexPage {
		fmt.Printf(gocolor.Blue+"[%v/%v]"+gocolor.Reset+" Generating index", len(options.Paths)+2, len(options.Paths)+extraSteps)
		fmt.Printf(gocolor.Yellow+" (%.1f ms)\n"+gocolor.Reset, goprocessing.Timing(func(){goprocessing.GenIndexPage(options)}))
	}
}


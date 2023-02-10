package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
	"workspace/gocolor"
	"workspace/goprocessing"
)

var lock sync.Mutex

func main() {
	goprocessing.CallClear()
	fmt.Println("Watching files for changes...")
	fmt.Printf(gocolor.Green+"[SUCCESS]"+gocolor.Reset+" All files has been processed "+gocolor.Yellow+"(Total: %.1f ms)\n"+gocolor.Reset, goprocessing.Timing(start))
}

func start() {
	/// TODO: PERMITIR RECIBIR UN LAYOUT COMO PARAMETRO CON ALGUNA ETIQUETA @MP_INCLUDE o @MP_INCLUDE_PART
	var options goprocessing.Options = goprocessing.ReadArguments()

	infoTasks := make(map[int]string)

	var wg sync.WaitGroup

	modifiedFiles := goprocessing.GetModifiedFiles(options)

	for i, e := range modifiedFiles {
		if !goprocessing.IsFileUpdated(e, options){
			wg.Add(1)
			go processAsyncFile(&wg, i, e, options, len(modifiedFiles), infoTasks)
			i++
		}
	}

	wg.Wait()

	if !goprocessing.IsFileUpdated("mpstyles.css", options){
		generateCss(options)
	}

	if (options.GenIndexPage && len(modifiedFiles) > 0) || options.ForceGeneration {
		generateIndex(options)
	}


	if options.WatchMode {
		if len(modifiedFiles) > 0 {
			fmt.Printf("Last update at %v\n\n", time.Now())
			fmt.Println("Watching files for changes...")
		}
		time.Sleep(1 * time.Second)
		start()
	}
}

func processAsyncFile(wg *sync.WaitGroup, i int, e string, options goprocessing.Options, totalFiles int, infoTasks map[int]string) {
	defer wg.Done()
	updateInfoTask(infoTasks, i, fmt.Sprintf(gocolor.Blue+"[%v/%v]"+gocolor.Reset+" Processing file %v", (i+1), totalFiles, e), false)
	print(infoTasks)
	timingResult := goprocessing.Timing(func(){goprocessing.GenFile(e, options)})
	updateInfoTask(infoTasks, i, fmt.Sprintf(gocolor.Yellow+" (%.1f ms)"+gocolor.Reset, timingResult), true)
	print(infoTasks)
}

func updateInfoTask (infoTasks map[int]string, item int, value string, append bool) {
	lock.Lock()
	defer lock.Unlock()
	if append {
		infoTasks[item] += value
	} else {
		infoTasks[item] = value
	}
}

func generateCss(options goprocessing.Options) {
	var pathToCss string = filepath.Join(options.Outdir, "mpstyles.css")
	fmt.Printf(gocolor.Blue+"[CSS] "+gocolor.Reset+"Generating styles file")
	timingResult := goprocessing.Timing(func(){goprocessing.GenCss(pathToCss, options)})
	fmt.Printf(gocolor.Yellow+" (%.1f ms)\n"+gocolor.Reset, timingResult)
}

func generateIndex(options goprocessing.Options) {
	var pathToIndex string = filepath.Join(options.Outdir, "index.html")
	fmt.Printf(gocolor.Blue+"[INDEX] "+gocolor.Reset+"Generating index file")
	timingResult := goprocessing.Timing(func(){goprocessing.GenIndexPage(pathToIndex, options)})
	fmt.Printf(gocolor.Yellow+" (%.1f ms)\n"+gocolor.Reset, timingResult)
}

func print(infoTasks map[int]string) {
    lock.Lock()
    defer lock.Unlock()
	
	goprocessing.CallClear()
	
	for i:=0; i<len(infoTasks); i++ {
		fmt.Println(infoTasks[i])
	}
}
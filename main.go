package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
	"workspace/goconsole"
	"workspace/goio"
	"workspace/gopreprocessing"
	"workspace/goutils"
)

var lock sync.Mutex

func main() {
	goconsole.CallClear()
	fmt.Println("Watching files for changes...")
	fmt.Printf(goconsole.Green+"[SUCCESS]"+goconsole.Reset+" All files has been processed "+goconsole.Yellow+"(Total: %.1f ms)\n"+goconsole.Reset, goutils.Timing(start))
}

func start() {
	/// TODO: PERMITIR RECIBIR UN LAYOUT COMO PARAMETRO CON ALGUNA ETIQUETA @MP_INCLUDE o @MP_INCLUDE_PART
	var options goconsole.Options = goconsole.ReadArguments()

	infoTasks := make(map[int]string)

	var wg sync.WaitGroup

	fmt.Println(goio.IsDestUpdated(options.Path, options.Outdir))
	if !goio.IsDestUpdated(options.Path, options.Outdir){
		restart(false)
	}

	modifiedFiles := goio.GetModifiedFiles(options.Paths, options.Outdir)

	for i, e := range modifiedFiles {
		if !goio.IsFileUpdated(e, options.Outdir) || options.ForceGeneration {
			wg.Add(1)
			go processAsyncFile(&wg, i, e, options, len(modifiedFiles), infoTasks)
			i++
		}
	}

	wg.Wait()

	if !goio.IsFileUpdated("mpstyles.css", options.Outdir) || options.ForceGeneration {
		generateCss(options)
	}

	if (options.GenIndexPage && len(modifiedFiles) > 0) || options.ForceGeneration {
		generateIndex(options)
	}


	if options.WatchMode {
		restart(len(modifiedFiles) > 0)
	}
}

func restart (modifiedFiles bool) {
	if modifiedFiles {
		fmt.Printf("Last update at %v\n\n", time.Now())
		fmt.Println("Watching files for changes...")
	}
	time.Sleep(1 * time.Second)
	start()
}

func processAsyncFile(wg *sync.WaitGroup, i int, e string, options goconsole.Options, totalFiles int, infoTasks map[int]string) {
	defer wg.Done()
	updateInfoTask(infoTasks, i, fmt.Sprintf(goconsole.Blue+"[%v/%v]"+goconsole.Reset+" Processing file %v", (i+1), totalFiles, e), false)
	print(infoTasks)
	timingResult := goutils.Timing(func(){gopreprocessing.GenFile(e, options)})
	updateInfoTask(infoTasks, i, fmt.Sprintf(goconsole.Yellow+" (%.1f ms)"+goconsole.Reset, timingResult), true)
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

func generateCss(options goconsole.Options) {
	var pathToCss string = filepath.Join(options.Outdir, "mpstyles.css")
	fmt.Printf(goconsole.Blue+"[CSS] "+goconsole.Reset+"Generating styles file")
	timingResult := goutils.Timing(func(){gopreprocessing.GenCss(pathToCss, options)})
	fmt.Printf(goconsole.Yellow+" (%.1f ms)\n"+goconsole.Reset, timingResult)
}

func generateIndex(options goconsole.Options) {
	var pathToIndex string = filepath.Join(options.Outdir, "index.html")
	fmt.Printf(goconsole.Blue+"[INDEX] "+goconsole.Reset+"Generating index file")
	timingResult := goutils.Timing(func(){gopreprocessing.GenIndexPage(pathToIndex, options)})
	fmt.Printf(goconsole.Yellow+" (%.1f ms)\n"+goconsole.Reset, timingResult)
}

func print(infoTasks map[int]string) {
    lock.Lock()
    defer lock.Unlock()
	
	goconsole.CallClear()
	
	for i:=0; i<len(infoTasks); i++ {
		fmt.Println(infoTasks[i])
	}
}
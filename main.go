package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
	"workspace/goconsole"
	"workspace/goio"
	"workspace/gopreprocessing"
	"workspace/gostyles"
	"workspace/goutils"
)

var lock sync.Mutex

// / TODO: PERMITIR RECIBIR UN LAYOUT COMO PARAMETRO CON ALGUNA ETIQUETA @MP_INCLUDE o @MP_INCLUDE_PART
func main() {
	goconsole.CallClear()
	fmt.Println("Watching files for changes...")
	fmt.Printf(goconsole.Green+"[SUCCESS]"+goconsole.Reset+" All files has been processed "+goconsole.Yellow+"(Total: %.1f ms)\n"+goconsole.Reset, goutils.Timing(start))
}

func start() {
	var options goconsole.Options = goconsole.ReadArguments()
	isNewContent := goio.IsNewContent(options.Path, options.Outdir)

	infoTasks := make(map[int]string)

	// Parse md files and generate HTML
	isSomeFileModified := genFiles(options, infoTasks)

	// FIX: Esto no actualiza el CSS en caso de que se inserte un CSS con -s y se modifique
	shouldGenCSS := !goio.AlreadyExistsCSS(options.Outdir, gostyles.FILENAME) || options.ForceGeneration
	shouldGenIndex := ((isNewContent || !goio.AlreadyExistsIndex(options.Outdir)) && options.GenIndexPage) || options.ForceGeneration

	// Repaint output if none
	if (shouldGenCSS || shouldGenIndex) && !isSomeFileModified {
		goconsole.CallClear()
		print(infoTasks)
	}

	// Generates styles
	if shouldGenCSS {
		generateCSS(options)
	}

	// Generates index.html if specified in input arguments and there are new content
	if shouldGenIndex {
		generateIndex(options)
	}

	// If any update, forces mtime of directory (this performs looking for changed files)
	anyUpdate := isSomeFileModified || shouldGenCSS || shouldGenIndex
	if anyUpdate {
		goio.UpdateOutdirInfo(options.Outdir)
	}

	// If watch mode activated, keep watching for changes
	if options.WatchMode {
		restart(anyUpdate, options.Outdir)
	}
}

func genFiles(options goconsole.Options, infoTasks map[int]string) bool {
	var filesToParse []string

	if options.ForceGeneration {
		filesToParse = options.Paths
	} else {
		// Check if any change in input files
		filesToParse = goio.GetModifiedFiles(options.Paths, options.Outdir)
	}

	var wg sync.WaitGroup
	// For every modified file, it starts an async process
	for i, e := range filesToParse {
		wg.Add(1)
		go processAsyncFile(&wg, i, e, options, len(filesToParse), infoTasks)
		i++
	}

	// Wait for all async file processing for finishing
	wg.Wait()

	return len(filesToParse) > 0
}

func restart(modifiedFiles bool, outdir string) {
	defer start()
	if modifiedFiles {
		goio.UpdateOutdirInfo(outdir)
		fmt.Printf("Last update at %v\n\n", time.Now())
		fmt.Println("Watching files for changes...")
	}
	time.Sleep(1 * time.Second)
}

func processAsyncFile(wg *sync.WaitGroup, i int, e string, options goconsole.Options, totalFiles int, infoTasks map[int]string) {
	defer wg.Done()
	updateInfoTask(infoTasks, i, fmt.Sprintf(goconsole.Blue+"[%v/%v]"+goconsole.Reset+" Processing file %v", (i+1), totalFiles, e), false)
	print(infoTasks)
	timingResult := goutils.Timing(func() { gopreprocessing.GenFile(e, options) })
	updateInfoTask(infoTasks, i, fmt.Sprintf(goconsole.Yellow+" (%.1f ms)"+goconsole.Reset, timingResult), true)
	print(infoTasks)
}

// Uptades infoTasks struct without collisions among process
func updateInfoTask(infoTasks map[int]string, item int, value string, append bool) {
	lock.Lock()
	defer lock.Unlock()
	if append {
		infoTasks[item] += value
	} else {
		infoTasks[item] = value
	}
}

func generateCSS(options goconsole.Options) {
	var pathToCSS string = filepath.Join(options.Outdir, gostyles.FILENAME)
	fmt.Printf(goconsole.Blue + "[CSS] " + goconsole.Reset + "Generating styles file")
	timingResult := goutils.Timing(func() { gopreprocessing.GenCSS(pathToCSS, options) })
	fmt.Printf(goconsole.Yellow+" (%.1f ms)\n"+goconsole.Reset, timingResult)
}

func generateIndex(options goconsole.Options) {
	var pathToIndex string = filepath.Join(options.Outdir, "index.html")
	fmt.Printf(goconsole.Blue + "[INDEX] " + goconsole.Reset + "Generating index file")
	timingResult := goutils.Timing(func() { gopreprocessing.GenIndexPage(pathToIndex, options) })
	fmt.Printf(goconsole.Yellow+" (%.1f ms)\n"+goconsole.Reset, timingResult)
}

// Prints output without collisions among process
func print(infoTasks map[int]string) {
	lock.Lock()
	defer lock.Unlock()

	goconsole.CallClear()

	for i := 0; i < len(infoTasks); i++ {
		fmt.Println(infoTasks[i])
	}
}

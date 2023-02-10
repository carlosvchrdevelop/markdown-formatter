package goprocessing

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func processLines (path string, f func(line string)) {
    readFile, err := os.Open(path)
  
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
		f(fileScanner.Text())
    }
  
    readFile.Close()
}

func writeFile (path string, text string) {
    var dirpath string = filepath.Dir(path)

    err := os.MkdirAll(dirpath, 0777)
	if err != nil {
		log.Println(err)
	}

    data := []byte(text)
    err = ioutil.WriteFile(path, data, 0777)

    if err != nil {
        log.Fatal(err)
    }
}

func readFile (path string) string {
    content, err := ioutil.ReadFile(path)

    if err != nil {
        log.Fatal(err)
    }
   
    return string(content)
}

func findFiles (root, ext string) []string {
    var a []string
    filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
        if e != nil { return e }
        if filepath.Ext(d.Name()) == ext {
            a = append(a, s)
        }
        return nil
    })
    return a
}

func GetModifiedFiles (options Options) []string {
    var modifiedFiles []string
    for _, e := range options.Paths {
        if !IsFileUpdated(e, options) {
            modifiedFiles = append(modifiedFiles, e)
        }
    }
    return modifiedFiles
}

func IsFileUpdated (path string, options Options) bool {

    if options.ForceGeneration {
        return false
    }

    var cleanPathParts []string = strings.Split(path, ".")
    var generatedFilePath string = strings.Join(cleanPathParts[0 : len(cleanPathParts)-1], "")
    var extension string = cleanPathParts[len(cleanPathParts)-1]

    if extension == "md" {
        generatedFilePath += ".html"
    } else if extension == "css" {
        generatedFilePath += ".css"
    } else if extension == "html" {
        generatedFilePath += ".html"
    }
        
    mdinfo, err := os.Stat(path)
    
    if err != nil  {
        log.Fatal(err)
    }
    
    generatedFilePath = filepath.Join(options.Outdir, generatedFilePath)
    genfileinfo, errgen := os.Stat(generatedFilePath)

    if errgen != nil  {
		return false
    }

    return mdinfo.ModTime().Before(genfileinfo.ModTime())
}
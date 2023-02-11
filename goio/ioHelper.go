package goio

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ProcessLines (path string, f func(line string)) {
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

func WriteFile (path string, text string) {
    var dirpath string = filepath.Dir(path)

    err := os.MkdirAll(dirpath, 0777)
	if err != nil {
		log.Println(err)
	}

    data := []byte(text)
    err = os.WriteFile(path, data, 0777)

    if err != nil {
        log.Fatal(err)
    }
}

func ReadFile (path string) string {
    content, err := os.ReadFile(path)

    if err != nil {
        log.Fatal(err)
    }
   
    return string(content)
}

func FindFiles (root, ext string) []string {
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

func GetModifiedFiles (paths []string, outdir string) []string {
    var modifiedFiles []string
    for _, e := range paths {
        if !IsFileUpdated(e, outdir) {
            modifiedFiles = append(modifiedFiles, e)
        }
    }
    return modifiedFiles
}

func AlreadyExistsIndex (outdir string) bool {
    _, err := os.Stat(filepath.Join(outdir, "index.html"))
    return err == nil
}

func AlreadyExistsCSS(outdir string, cssfilename string) bool {
    _, err := os.Stat(filepath.Join(outdir, cssfilename))
    return err == nil
}

func UpdateOutdirInfo (outdir string) {
    fileName := outdir
    currentTime := time.Now().Local()

    //Set both access time and modified time of the file to the current time
    err := os.Chtimes(fileName, currentTime, currentTime)
    if err != nil {
        log.Panic(err)
    }
}

func IsNewContent (path string, path2 string) bool {
    originInfo, err := os.Stat(path)

    if err != nil  {
		log.Panic(err)
    }

    destInfo, err := os.Stat(path2)

    if err != nil {
        return true
    }

    return originInfo.ModTime().After(destInfo.ModTime())
}

func IsFileUpdated (path string, outdir string) bool {

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
    
    generatedFilePath = filepath.Join(outdir, generatedFilePath)
    genfileinfo, errgen := os.Stat(generatedFilePath)

    if errgen != nil  {
		return false
    }

    return mdinfo.ModTime().Before(genfileinfo.ModTime())
}
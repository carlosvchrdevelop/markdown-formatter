package goprocessing

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func findFiles(root, ext string) []string {
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
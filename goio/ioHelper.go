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

// ProcessLines is a function that opens a file an execute callback for each line
func ProcessLines(path string, f func(line string)) {
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

// WriteFile is a function for writting files given the path and value to write
func WriteFile(path string, text string) {
	dirpath := filepath.Dir(path)
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

// ReadFile is a function that retrieves the content of a file in a string variable
func ReadFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

// FindFiles is a function that retrieves all files with an specific extension
func FindFiles(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

// GetModifiedFiles is a function that retrieves only the files that has been modified
func GetModifiedFiles(paths []string, outdir string) []string {
	var modifiedFiles []string
	for _, e := range paths {
		if !IsFileUpdated(e, outdir) {
			modifiedFiles = append(modifiedFiles, e)
		}
	}
	return modifiedFiles
}

// AlreadyExistsIndex determines whether index.html has been generated or not
func AlreadyExistsIndex(outdir string) bool {
	_, err := os.Stat(filepath.Join(outdir, "index.html"))
	return err == nil
}

// AlreadyExistsCSS determines whether mpstyles.css has been generated or not
func AlreadyExistsCSS(outdir string, cssfilename string) bool {
	_, err := os.Stat(filepath.Join(outdir, cssfilename))
	return err == nil
}

// UpdateOutdirInfo updates the modification time to current time
func UpdateOutdirInfo(outdir string) {
	fileName := outdir
	currentTime := time.Now().Local()

	//Set both access time and modified time of the file to the current time
	err := os.Chtimes(fileName, currentTime, currentTime)
	if err != nil {
		log.Panic(err)
	}
}

// IsNewContent verifies if new content has been added and not generated yet
func IsNewContent(path string, path2 string) bool {
	originInfo, err := os.Stat(path)

	if err != nil {
		log.Panic(err)
	}

	destInfo, err := os.Stat(path2)

	if err != nil {
		return true
	}

	return originInfo.ModTime().After(destInfo.ModTime())
}

// IsFileUpdated retrieves whether a file has been generated after its modification or not
func IsFileUpdated(path string, outdir string) bool {

	cleanPathParts := strings.Split(path, ".")
	generatedFilePath := strings.Join(cleanPathParts[0:len(cleanPathParts)-1], "")
	extension := cleanPathParts[len(cleanPathParts)-1]

	if extension == "md" {
		generatedFilePath += ".html"
	} else if extension == "css" {
		generatedFilePath += ".css"
	} else if extension == "html" {
		generatedFilePath += ".html"
	}

	mdinfo, err := os.Stat(path)

	if err != nil {
		log.Fatal(err)
	}

	generatedFilePath = filepath.Join(outdir, generatedFilePath)
	genfileinfo, errgen := os.Stat(generatedFilePath)

	if errgen != nil {
		return false
	}

	return mdinfo.ModTime().Before(genfileinfo.ModTime())
}

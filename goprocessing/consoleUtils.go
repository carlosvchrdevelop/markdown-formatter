package goprocessing

import (
	"fmt"
	"log"
	"os"
	"strings"
	"workspace/gocolor"
)

const help string = `
Usage:

$ mp [options] <dir>

Options:

	-p: generates a partial html, without the head and body tags.

	-i: generates an index.html file with references to all pages.

	-s,--css <file.css>: specifies a css file that overrides the default styles.

	-o,--outdir <dir>: specifies the directory in which the output is generated.

`

type Options struct {
	ExternalCss string
	GenFullPage bool
	Outdir string
	GenIndexPage bool
	Paths []string
}

func ReadArguments () Options {

	var options Options = Options {
		ExternalCss: "",
		GenFullPage: true,
		GenIndexPage: false,
		Outdir: "./",
	}

	if len(os.Args) < 2 {
		fmt.Print(gocolor.Red+"[ERROR]"+gocolor.Reset+" ")
		log.Fatal("Unexpected number of parameters. Run 'mf --help' for more information.")
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Print(help)
		os.Exit(0)
	}

	options.Paths = findFiles(os.Args[len(os.Args)-1], ".md")

	if len(options.Paths) == 0 {
		fmt.Print(gocolor.Red+"[ERROR]"+gocolor.Reset+" ")
		log.Fatal("No .md format file could be found.")
	}

	for i:=1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") && strings.Contains(os.Args[i], "p") {
			options.GenFullPage = false
		}

		if strings.HasPrefix(os.Args[i], "-") && strings.Contains(os.Args[i], "i") {
			options.GenIndexPage = true
		}

		if os.Args[i] == "-s" || os.Args[i] == "--css"{
			i += 1
			options.ExternalCss = readFile(os.Args[i])
		}

		if os.Args[i] == "-o" || os.Args[i] == "--outdir" {
			i += 1
			options.Outdir = os.Args[i]
		}
		

	}
	return options
}
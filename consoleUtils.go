package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	externalCss string
	genFullPage bool
	outdir string
	genIndexPage bool
	paths []string
}

func readArguments () Options {

	var options Options = Options {
		externalCss: "",
		genFullPage: true,
		genIndexPage: false,
		outdir: "./",
	}

	if len(os.Args) < 2 {
		log.Fatal("[ERROR] Unexpected number of parameters. Run 'mf --help' for more information.")
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Print(help)
		os.Exit(0)
	}

	options.paths = findFiles(os.Args[len(os.Args)-1], ".md")

	if len(options.paths) == 0 {
		log.Fatal("[ERROR] No .md format file could be found.")
	}

	for i:=1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") && strings.Contains(os.Args[i], "p") {
			options.genFullPage = false
		}

		if strings.HasPrefix(os.Args[i], "-") && strings.Contains(os.Args[i], "i") {
			options.genIndexPage = true
		}

		if os.Args[i] == "-s" || os.Args[i] == "--css"{
			i += 1
			options.externalCss = readFile(os.Args[i])
		}

		if os.Args[i] == "-o" || os.Args[i] == "--outdir" {
			i += 1
			options.outdir = os.Args[i]
		}
		

	}
	return options
}
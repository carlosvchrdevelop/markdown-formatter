package main

import (
	"log"
	"path/filepath"
	"strings"
)

func genFile (path string, options Options) {

	var lines = []string{}
 
    processLines(path, func(line string){
		processAll(&line)
        lines = append(lines, line)
	})

	var file string = strings.Join(lines, "\n")
	var relpath, err = filepath.Rel(filepath.Join(options.outdir, path),options.outdir)

	if err != nil {
		log.Println(err)
	}

	if options.genFullPage {
		wrapFullPage(filepath.Join(relpath,"mpstyles.css"), &file)
	}

	writeFile(strings.Replace(filepath.Join(options.outdir, path), ".md", ".html", 1), file)
}

func genCss (options Options) {
	writeFile(filepath.Join(options.outdir, "mpstyles.css"), readFile("./mpstyles.css") + "\n" + options.externalCss)
}

func genIndexPage (options Options) {
	writeFile(filepath.Join(options.outdir, "index.html"), getIndexPage(options.paths))
}

func getIndexPage (references []string) string {
	var page string = ""
	for _, e := range references {
		var curfname = strings.Replace(filepath.Base(e), filepath.Ext(e), "", 1)
		var curfpath = strings.Replace(e, filepath.Ext(e), ".html", 1)
		page += "<a href='"+curfpath+"'>"+curfname+"</a>"
	}
	wrapFullPage("mpstyles.css", &page)
	return page
}

func wrapFullPage (path string, file *string) {
	var pre = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
	<link rel="stylesheet" href="`+path+`">
</head>
<body>
`

	const after = `
</body>
</html>
`
	*file = pre + *file + after;
}
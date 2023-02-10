package goprocessing

import (
	"log"
	"path/filepath"
	"strings"
)

func GenFile (path string, options Options) {

	var lines = []string{}
 
    processLines(path, func(line string){
		processAll(&line)
        lines = append(lines, line)
	})

	var file string = strings.Join(lines, "\n")
	var relpath, err = filepath.Rel(filepath.Join(options.Outdir, path),options.Outdir)

	if err != nil {
		log.Println(err)
	}

	if options.GenFullPage {
		WrapFullPage(filepath.Join(relpath,"mpstyles.css"), &file)
	}

	writeFile(strings.Replace(filepath.Join(options.Outdir, path), ".md", ".html", 1), file)
}

func GenCss (options Options) {
	writeFile(filepath.Join(options.Outdir, "mpstyles.css"), readFile("./mpstyles.css") + "\n" + options.ExternalCss)
}

func GenIndexPage (options Options) {
	writeFile(filepath.Join(options.Outdir, "index.html"), getIndexPage(options.Paths))
}

func getIndexPage (references []string) string {
	var page string = ""
	for _, e := range references {
		var curfname = strings.Replace(filepath.Base(e), filepath.Ext(e), "", 1)
		var curfpath = strings.Replace(e, filepath.Ext(e), ".html", 1)
		page += "<a href='"+curfpath+"'>"+curfname+"</a>"
	}
	WrapFullPage("mpstyles.css", &page)
	return page
}

func WrapFullPage (path string, file *string) {
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
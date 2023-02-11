package gopreprocessing

import (
	"regexp"
	"strings"
)

var openParagraph bool

func init () {
	openParagraph = false
}

func processAll(line *string) {
	blockCode := processCode(line)
	if !blockCode {
		processHr(line)
		processStyles(line)
		processParagraph(line)
		processHeaders(line)
		processMedia(line)
		processLinks(line)
		processOrderedLists(line)
	}
}

func processParagraph (line *string) {
	copyLine := strings.TrimSpace(*line)
	if len(copyLine) != 0 {
		return
	}
	openParagraph = !openParagraph
	if openParagraph {
		*line = "<p>"
	} else {
		*line = "</p>"
	}
}

func processHr (line *string) {
	if strings.HasPrefix(strings.TrimSpace(*line), "---") {
		*line = "<hr>"
	}
}

func processStyles (line *string) {
	var regexBold = regexp.MustCompile(`(\*\*[^\*]*\*\*)|(__[^_]*__)`)
	var regexItalic = regexp.MustCompile(`(\*[^\*]*\*)|(_[^_]*_)`)
	var regexStrike = regexp.MustCompile(`(\~~[^\~]*\~~)`)
	var regexCite = regexp.MustCompile(`(\""[^\"]*\"")`)
	
	var ocur string = regexBold.FindString(*line)
	for len(ocur) > 0 {
		parsed := "<b>"+ocur[2:len(ocur)-2]+"</b>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexBold.FindString(*line)
	}
	
	ocur = regexItalic.FindString(*line)
	for len(ocur) > 0 {
		parsed:= "<i>"+ocur[1:len(ocur)-1]+"</i>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexItalic.FindString(*line)
	}

	ocur = regexStrike.FindString(*line)
	for len(ocur) > 0 {
		parsed:= "<strike>"+ocur[2:len(ocur)-2]+"</strike>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexStrike.FindString(*line)
	}

	ocur = regexCite.FindString(*line)
	for len(ocur) > 0 {
		parsed:= "<cite>"+ocur[2:len(ocur)-2]+"</cite>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexCite.FindString(*line)
	}
}

func processHeaders (line *string) {
	if strings.HasPrefix(*line, "##### "){
		*line = strings.Replace(*line, "##### ", "<h5>", 1) + "</h5>"
	}
	if strings.HasPrefix(*line, "#### "){
		*line = strings.Replace(*line, "#### ", "<h4>", 1) + "</h4>"
	}
	if strings.HasPrefix(*line, "### "){
		*line = strings.Replace(*line, "### ", "<h3>", 1) + "</h3>"
	}
	if strings.HasPrefix(*line, "## "){
		*line = strings.Replace(*line, "## ", "<h2>", 1) + "</h2>"
	}
	if strings.HasPrefix(*line, "# "){
		*line = strings.Replace(*line, "# ", "<h1>", 1) + "</h1>"
	}
}

// TODO: terminar esto
func processOrderedLists (line *string) {
	// if strings.HasPrefix(*line, "-"){
	// 	*line = strings.Replace(*line, "-", "<li>", 1) + "</li>"
	// }
}

func processMedia (line *string) {
	var regex = regexp.MustCompile(`\!\[[^\[]*\]\([^(]*\)`)
	var ocur string = regex.FindString(*line)

	for len(ocur) > 0 {
		var parts []string = strings.Split(ocur, "](")
		parts[0] = strings.Replace(parts[0],"![","", 1)
		parts[1] = strings.Replace(parts[1],")","", 1)
		urlname := strings.Split(parts[1], ".")
		extension := urlname[len(urlname)-1]

		if _, ok := MediaFormats[extension]; ok {
			var parsed string
			switch MediaFormats[extension] {
			case IMAGE:
				parsed = "<img src='"+parts[1]+"' alt='"+parts[0]+"' />"
			case VIDEO:
				parsed = "<video src='"+parts[1]+"' alt='"+parts[0]+"' controls ></video>"
			case AUDIO:
				parsed = "<audio src='"+parts[1]+"' alt='"+parts[0]+"' controls ></video>"
			}
			*line = strings.Replace(*line, ocur, parsed, 1)
		} else {
			break
		}

		ocur = regex.FindString(*line)
	}
}


func processLinks (line *string) {
	var regex = regexp.MustCompile(`\[[^\[]*\]\([^(]*\)`)
	var ocur string = regex.FindString(*line)
	for len(ocur) > 0 {
		var parts []string = strings.Split(ocur, "](")
		parts[0] = strings.Replace(parts[0],"[","", 1)
		parts[1] = strings.Replace(parts[1],")","", 1)
		var parsedLink string = "<a href='"+parts[1]+"'>"+parts[0]+"</a>"
		*line = strings.Replace(*line, ocur, parsedLink, 1)
		ocur = regex.FindString(*line)
	}
}
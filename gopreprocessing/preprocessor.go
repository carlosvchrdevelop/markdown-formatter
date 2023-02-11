package gopreprocessing

import (
	"regexp"
	"strings"
)

var intoOrderedList = false

func processAll(line *string) {
	if lang == NONE {
		processHeaders(line)
		processLinks(line)
	}
	processCode(line)
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

func processOrderedLists (line *string) {
	if strings.HasPrefix(*line, "-"){
		*line = strings.Replace(*line, "-", "<li>", 1) + "</li>"
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
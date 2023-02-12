package gopreprocessing

import (
	"regexp"
	"strings"
)

var (
	openParagraph bool
	ids           string
	classes       string
	jsblocks      string
)

func init() {
	openParagraph = false
}

func processAll(line *string) {
	blockCode := processCode(line)
	if !blockCode {
		processHr(line)
		processDecorators(line)
		processStyles(line)
		processParagraph(line)
		processHeaders(line)
		processMedia(line)
		processLinks(line)
		processOrderedLists(line)
	}
}

// When finishes all processing, this closes all opened tags
func processClosing() string {
	var closingItems string
	if openParagraph {
		closingItems += "</p>"
	}
	return closingItems
}

// processDecorators processes tags like @ID, @CLASS, etc.
func processDecorators(line *string) {
	copyLine := strings.TrimSpace(*line)

	regexClassContent := regexp.MustCompile(SeparatedCommaIdentifiers)
	regexIDContent := regexp.MustCompile(IdentifierPattern)

	if len(regexp.MustCompile(RegexID).FindString(*line)) > 0 {
		ids = strings.TrimSpace(regexIDContent.FindString(strings.Split(copyLine, "(")[1]))
		*line = ""
	} else if len(regexp.MustCompile(RegexClass).FindString(*line)) > 0 {
		matches := strings.Split(regexClassContent.FindString(strings.Split(copyLine, "(")[1]), ",")
		for i := 0; i < len(matches); i++ {
			matches[i] = strings.TrimSpace(matches[i])
		}
		classes = strings.Join(matches, " ")
		*line = ""
	} else if len(regexp.MustCompile(RegexJS).FindString(*line)) > 0 {
		jsblocks = regexClassContent.FindString(strings.Split(copyLine, "(")[1])
		*line = ""
	}
}

// addIdentifiers generate tags for ids, classes and js calls
func addIdentifiers(tag string) string {
	var identifiersUnified string

	if len(ids) > 0 {
		identifiersUnified += " id='" + ids + "'"
		ids = ""
	}

	if len(classes) > 0 {
		identifiersUnified += " class='" + classes + "'"
		classes = ""
	}

	// TODO: implementar JS

	return strings.Split(tag, ">")[0] + identifiersUnified + ">"
}

func processParagraph(line *string) {
	copyLine := strings.TrimSpace(*line)
	if len(copyLine) != 0 {
		return
	}
	openParagraph = !openParagraph
	if openParagraph {
		*line = addIdentifiers("<p>")
	} else {
		*line = "</p>"
	}
}

func processHr(line *string) {
	if strings.HasPrefix(strings.TrimSpace(*line), "---") {
		*line = addIdentifiers("<hr>")
	}
}

func processStyles(line *string) {
	var regexBold = regexp.MustCompile(`(\*\*[^\*]*\*\*)|(__[^_]*__)`)
	var regexItalic = regexp.MustCompile(`(\*[^\*]*\*)|(_[^_]*_)`)
	var regexStrike = regexp.MustCompile(`(\~~[^\~]*\~~)`)
	var regexCite = regexp.MustCompile(`(\""[^\"]*\"")`)

	var ocur string = regexBold.FindString(*line)
	for len(ocur) > 0 {
		parsed := "<b>" + ocur[2:len(ocur)-2] + "</b>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexBold.FindString(*line)
	}

	ocur = regexItalic.FindString(*line)
	for len(ocur) > 0 {
		parsed := "<i>" + ocur[1:len(ocur)-1] + "</i>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexItalic.FindString(*line)
	}

	ocur = regexStrike.FindString(*line)
	for len(ocur) > 0 {
		parsed := "<strike>" + ocur[2:len(ocur)-2] + "</strike>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexStrike.FindString(*line)
	}

	ocur = regexCite.FindString(*line)
	for len(ocur) > 0 {
		parsed := "<cite>" + ocur[2:len(ocur)-2] + "</cite>"
		*line = strings.Replace(*line, ocur, parsed, 1)
		ocur = regexCite.FindString(*line)
	}
}

func processHeaders(line *string) {
	if strings.HasPrefix(*line, "##### ") {
		*line = strings.Replace(*line, "##### ", addIdentifiers("<h5>"), 1) + "</h5>"
	}
	if strings.HasPrefix(*line, "#### ") {
		*line = strings.Replace(*line, "#### ", addIdentifiers("<h4>"), 1) + "</h4>"
	}
	if strings.HasPrefix(*line, "### ") {
		*line = strings.Replace(*line, "### ", addIdentifiers("<h3>"), 1) + "</h3>"
	}
	if strings.HasPrefix(*line, "## ") {
		*line = strings.Replace(*line, "## ", addIdentifiers("<h2>"), 1) + "</h2>"
	}
	if strings.HasPrefix(*line, "# ") {
		*line = strings.Replace(*line, "# ", addIdentifiers("<h1>"), 1) + "</h1>"
	}
}

// TODO: terminar esto
func processOrderedLists(line *string) {
	// if strings.HasPrefix(*line, "-"){
	// 	*line = strings.Replace(*line, "-", "<li>", 1) + "</li>"
	// }
}

func processMedia(line *string) {
	var regex = regexp.MustCompile(`\!\[[^\[]*\]\([^(]*\)`)
	var ocur string = regex.FindString(*line)

	for len(ocur) > 0 {
		var parts []string = strings.Split(ocur, "](")
		parts[0] = strings.Replace(parts[0], "![", "", 1)
		parts[1] = strings.Replace(parts[1], ")", "", 1)
		urlname := strings.Split(parts[1], ".")
		extension := urlname[len(urlname)-1]

		if _, ok := mediaFormats[extension]; ok {
			var parsed string
			switch mediaFormats[extension] {
			case IMAGE:
				parsed = addIdentifiers("<img src='" + parts[1] + "' alt='" + parts[0] + "' />")
			case VIDEO:
				parsed = addIdentifiers("<video src='"+parts[1]+"' alt='"+parts[0]+"' controls >") + "</video>"
			case AUDIO:
				parsed = addIdentifiers("<audio src='"+parts[1]+"' alt='"+parts[0]+"' controls >") + "</audio>"
			}
			*line = strings.Replace(*line, ocur, parsed, 1)
		} else {
			break
		}

		ocur = regex.FindString(*line)
	}
}

func processLinks(line *string) {
	var regex = regexp.MustCompile(`\[[^\[]*\]\([^(]*\)`)
	var ocur string = regex.FindString(*line)
	for len(ocur) > 0 {
		var parts []string = strings.Split(ocur, "](")
		parts[0] = strings.Replace(parts[0], "[", "", 1)
		parts[1] = strings.Replace(parts[1], ")", "", 1)
		var parsedLink string = addIdentifiers("<a href='"+parts[1]+"'>") + parts[0] + "</a>"
		*line = strings.Replace(*line, ocur, parsedLink, 1)
		ocur = regex.FindString(*line)
	}
}

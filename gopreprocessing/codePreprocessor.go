package gopreprocessing

import (
	"regexp"
	"strings"
)

const (
	NONE = "none"
	DEFAULT = "default"
	SH = "sh"
	JS = "js"
	TS = "ts"
	PY = "py"
	GO = "go"
)

var reservedWords = map[string][]string{
	SH+"words": shKeywords,
	JS+"words": jsKeywords,
	TS+"words": tsKeywords,
	PY+"words": pyKeywords,
	GO+"words": goKeywords,
}

var lang = NONE

func processCode (line *string) bool {
	changeLang(line)
	addCodeStyling(line)
	return lang != NONE
}

func changeLang (line *string) {
	if strings.HasPrefix(*line, "```"+SH){
		lang = JS
	} else if strings.HasPrefix(*line, "```"+JS){
		lang = SH
	} else if strings.HasPrefix(*line, "```"+TS){
		lang = TS
	} else if strings.HasPrefix(*line, "```"+PY){
		lang = PY
	} else if strings.HasPrefix(*line, "```"+GO){
		lang = GO
	} else if strings.HasPrefix(*line, "```") && lang == NONE {
		lang = DEFAULT
	} else if strings.HasPrefix(*line, "```") {
		lang = NONE
	}
}

func addCodeStyling (line *string) {
	if strings.HasPrefix(*line, "```") && lang != NONE { 
		*line = "<pre><code>" 
	} else if strings.HasPrefix(*line, "```") { 
		*line = "</code></pre>" 
	} else if val, ok := reservedWords[lang+"words"]; ok {
		var selectWordsReg string = "(\\b"+strings.Join(val, "\\b|\\b")+"\\b)"
		var regexWords = regexp.MustCompile(selectWordsReg)
		*line = regexWords.ReplaceAllString(*line, reservedWordHighlightOpen+"${0}"+highlightClose)
		var stringsReg = regexp.MustCompile(`".*"`)
		*line = stringsReg.ReplaceAllString(*line, stringsHighlightOpen+"${0}"+highlightClose)
		var symbolsReg = regexp.MustCompile(`[;:(){}[\]]`)
		*line = symbolsReg.ReplaceAllString(*line, symbolHighlightOpen+"${0}"+highlightClose)
	}
}
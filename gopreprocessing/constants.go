package gopreprocessing

const (
	prefix                    = "mp"
	highlightClose            = "</span>"
	reservedWordHighlightOpen = "<span class='" + prefix + "-reserved-word'>"
	stringsHighlightOpen      = "<span class='" + prefix + "-string'>"
	symbolHighlightOpen       = "<span class='" + prefix + "-symbol'>"
)

const (
	// AUDIO media format
	AUDIO int = 0
	// VIDEO media format
	VIDEO int = 1
	// IMAGE media format
	IMAGE int = 2
)

var mediaFormats = map[string]int{
	"aac": AUDIO, "m4a": AUDIO, "mp3": AUDIO, "wav": AUDIO, "wave": AUDIO,
	"wma": AUDIO, "oga": AUDIO, "ogg": AUDIO, "bmp": IMAGE, "gif": IMAGE,
	"jpg": IMAGE, "jpeg": IMAGE, "png": IMAGE, "svg": IMAGE, "webp": IMAGE,
	"avi": VIDEO, "m4v": VIDEO, "mkv": VIDEO, "mp4": VIDEO, "mov": VIDEO,
	"ogv": VIDEO, "mpeg": VIDEO, "qt": VIDEO, "webm": VIDEO, "wmv": VIDEO,
}

var shKeywords = []string{
	"case", "coproc", "do", "done", "elif", "else",
	"esac", "fi", "for", "function", "if", "in",
	"select", "then", "until", "while",
	"CASE", "COPROC", "DO", "DONE", "ELIF", "ELSE",
	"ESAC", "FI", "FOR", "FUNCTION", "IF", "IN",
	"SELECT", "THEN", "UNTIL", "WHILE"}

var goKeywords = []string{
	"break", "default", "func", "interface", "select", "var",
	"case", "defer", "go", "map", "struct", "chan", "else",
	"goto", "package", "switch", "const", "fallthrough", "if",
	"range", "type", "continue", "for", "import", "return",
}

var jsKeywords = []string{
	"await", "break", "case", "catch", "class", "switch",
	"const", "continue", "debugger", "default", "delete",
	"do", "else", "enum", "export", "extends", "false",
	"finally", "for", "function", "if", "interface",
	"implements", "import", "in", "instanceof", "this",
	"let", "new", "null", "package", "private", "throw",
	"protected", "public", "return", "super", "true", "try",
	"typeof", "var", "void", "while", "with", "static", "yield",
}

var tsKeywords = []string{
	"break", "as", "any", "switch", "case", "if", "yield",
	"var", "number", "string", "get", "throw", "else",
	"module", "type", "instanceof", "typeof", "void", "new",
	"public", "private", "enum", "export", "null", "const",
	"finally", "for", "while", "super", "this", "new",
	"in", "return", "true", "false", "static", "let", "try",
	"package", "implements", "interface", "function",
	"continue", "do", "catch", "any", "extends",
}

var pyKeywords = []string{
	"and", "as", "assert", "break", "class", "continue",
	"def", "del", "elif", "else", "except", "False", "yield",
	"finally", "for", "from", "global", "if", "import",
	"in", "is", "lambda", "None", "nonlocal", "not", "or",
	"pass", "raise", "return", "True", "try", "while", "with",
}

package gopreprocessing

const (
	AUDIO int = 0
	VIDEO int = 1
	IMAGE int = 2
)

var MediaFormats = map[string]int{
	"aac": AUDIO,
	"m4a": AUDIO,
	"mp3": AUDIO,
	"wav": AUDIO,
	"wave": AUDIO,
	"wma": AUDIO,
	"oga": AUDIO,
	"ogg": AUDIO,
	"bmp": IMAGE,
	"gif": IMAGE,
	"jpg": IMAGE,
	"jpeg": IMAGE,
	"png": IMAGE,
	"svg": IMAGE,
	"webp": IMAGE,
	"avi": VIDEO,
	"m4v": VIDEO,
	"mkv": VIDEO,
	"mp4": VIDEO,
	"mov": VIDEO,
	"ogv": VIDEO,
	"mpeg": VIDEO,
	"qt": VIDEO,
	"webm": VIDEO,
	"wmv": VIDEO,
}
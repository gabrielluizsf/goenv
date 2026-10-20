package lang

type lang int

const (
	langEN lang = iota
	langPTBR
)

var current = detectLang()

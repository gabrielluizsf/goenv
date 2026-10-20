package lang

// Translate devolve o texto no idioma detectado.
func Translate(en, pt string) string {
	if current == langPTBR {
		return pt
	}
	return en
}

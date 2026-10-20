package lang

import (
	"os"
	"strings"

	"github.com/gabrielluizsf/goenv/tz"
)

// detectLang: GOENV_LANG (pt-br | en) força o idioma; senão usa a timezone.
// Timezone brasileira -> pt-BR, qualquer outra -> inglês.
func detectLang() lang {
	switch strings.ToLower(strings.ReplaceAll(os.Getenv("GOENV_LANG"), "_", "-")) {
	case "pt-br", "pt":
		return langPTBR
	case "en":
		return langEN
	}
	if tz.IsBrazil(tz.Timezone.Name()) {
		return langPTBR
	}
	return langEN
}

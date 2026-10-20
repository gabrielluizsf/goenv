package env

import (
	"strings"

	"github.com/gabrielluizsf/goenv/cmdx"
	"github.com/gabrielluizsf/goenv/lang"
)

func Merge(base, extra []string) []string {
	for _, kv := range extra {
		k, _, ok := strings.Cut(kv, "=")
		if !ok {
			cmdx.Fatal(lang.Translate("invalid variable (use KEY=VALUE): %q", "variável inválida (use KEY=VALUE): %q"), kv)
		}
		replaced := false
		for i, b := range base {
			bk, _, _ := strings.Cut(b, "=")
			if strings.EqualFold(bk, k) { // Windows é case-insensitive
				base[i], replaced = kv, true
				break
			}
		}
		if !replaced {
			base = append(base, kv)
		}
	}
	return base
}

package tz

import "strings"

// Nomes IANA (Linux/macOS) e IDs do Windows (tzutil /g) do Brasil.
var brazilTZ = map[string]bool{
	// IANA
	"america/sao_paulo": true, "america/bahia": true, "america/fortaleza": true,
	"america/recife": true, "america/belem": true, "america/maceio": true,
	"america/araguaina": true, "america/santarem": true, "america/manaus": true,
	"america/porto_velho": true, "america/boa_vista": true, "america/cuiaba": true,
	"america/campo_grande": true, "america/rio_branco": true, "america/eirunepe": true,
	"america/noronha": true,
	// Windows
	"e. south america standard time": true, "central brazil standard time": true,
	"tocantins standard time": true, "bahia standard time": true,
	"sa eastern standard time": true, "sa western standard time": true,
}

func IsBrazil(name string) bool {
	n := strings.ToLower(strings.TrimSpace(name))
	return brazilTZ[n] || strings.HasPrefix(n, "brazil/")
}

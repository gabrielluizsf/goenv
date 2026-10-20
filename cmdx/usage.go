package cmdx

import (
	"fmt"
	"os"

	"github.com/gabrielluizsf/goenv/lang"
)


func Usage() { fmt.Fprint(os.Stderr, lang.Translate(usageEN, usagePT)) }


const usageEN = `goenv - run any app with environment variables

Usage:
  goenv [flags] -- <command> [args...]

Flags:
  -e KEY=VALUE     set a variable (repeatable)
  --file PATH      load variables from a .env file (alias: --env-file)
  -C DIR           working directory
  -log PATH        log file (default: goenv.log)
  -fg              run in foreground

Examples:
  goenv -e PORT=8080 -e DB_URL=postgres://localhost/app -- go run ./cmd/api
  goenv --file='.env' -- ./api
  goenv --file=.env -e PORT=9000 -C ./backend -- go run .
`

const usagePT = `goenv - sobe qualquer app com variáveis de ambiente

Uso:
  goenv [flags] -- <comando> [args...]

Flags:
  -e KEY=VALUE     define variável (repetível)
  --file PATH      carrega variáveis de um .env (alias: --env-file)
  -C DIR           diretório de trabalho
  -log PATH        arquivo de log (padrão: goenv.log)
  -fg              roda em foreground

Exemplos:
  goenv -e PORT=8080 -e DB_URL=postgres://localhost/app -- go run ./cmd/api
  goenv --file='.env' -- ./api
  goenv --file=.env -e PORT=9000 -C ./backend -- go run .
`
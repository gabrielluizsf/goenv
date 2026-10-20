package cmdx

import (
	"fmt"
	"os"
)

func Fatal(f string, a ...any) {
	fmt.Fprintf(os.Stderr, "goenv: "+f+"\n", a...)
	os.Exit(1)
}

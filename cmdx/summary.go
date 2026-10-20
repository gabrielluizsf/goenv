package cmdx

import (
	"fmt"
	"runtime"

	"github.com/gabrielluizsf/goenv/lang"
)

func PrintSummary(pid int, logPath string, detached bool) {
	fmt.Printf("\n✔ %s (%s)\n", lang.Translate("App started", "App iniciado"), runtime.GOOS)
	fmt.Printf("  PID: %d\n", pid)
	if logPath != "" {
		fmt.Printf("  Log: %s\n", logPath)
	}
	fmt.Printf("\n%s\n", lang.Translate("To stop (copy and paste):", "Para parar (copie e cole):"))
	fmt.Printf("  %s\n\n", Cmd.KillCommand(pid, detached))
}

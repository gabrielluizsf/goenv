package cmdx

import (
	"os"
	"os/exec"
)

type CMD interface {
	Detach(cmd *exec.Cmd)
	PrepareChild(cmd *exec.Cmd)
	Terminate(p *os.Process)
	KillCommand(pid int, detached bool) string
}

var Cmd CMD = cmd{}

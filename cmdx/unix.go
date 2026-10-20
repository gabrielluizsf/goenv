//go:build !windows

package cmdx

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type cmd struct {}

func (cmd) Detach(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
}

func (cmd) PrepareChild(cmd *exec.Cmd) {}

func (cmd) Terminate(p *os.Process) { _ = p.Signal(syscall.SIGTERM) }

func (cmd) KillCommand(pid int, detached bool) string {
	if detached {
		// "-PID" mata o grupo inteiro (supervisor + app + filhos)
		return fmt.Sprintf("kill -TERM -- -%d", pid)
	}
	return fmt.Sprintf("kill %d", pid)
}

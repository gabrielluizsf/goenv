//go:build windows

package cmdx

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const (
	detachedProcess = 0x00000008
	createNoWindow  = 0x08000000
)

type cmd struct {}

func (cmd) Detach(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP | detachedProcess,
		HideWindow:    true,
	}
}

// O supervisor não tem console; sem isso o app abriria uma janela.
func (cmd) PrepareChild(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: createNoWindow,
		HideWindow:    true,
	}
}

func (cmd) Terminate(p *os.Process) { _ = p.Kill() }

func (cmd) KillCommand(pid int, _ bool) string {
	// /T mata a árvore (supervisor + app + filhos), /F força
	return fmt.Sprintf("taskkill /PID %d /T /F", pid)
}

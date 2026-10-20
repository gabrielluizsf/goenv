//go:build windows

package tz

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type timezone struct{}

func (timezone) Name() string {
	if tz := os.Getenv("TZ"); tz != "" {
		return tz
	}
	cmd := exec.Command("tzutil", "/g")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

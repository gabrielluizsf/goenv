//go:build !windows

package tz

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type timezone struct{}

func (timezone) Name() string {
	if tz := strings.TrimPrefix(os.Getenv("TZ"), ":"); tz != "" {
		return tz
	}
	if b, err := os.ReadFile("/etc/timezone"); err == nil { // Debian/Ubuntu
		if s := strings.TrimSpace(string(b)); s != "" {
			return s
		}
	}
	if p, err := filepath.EvalSymlinks("/etc/localtime"); err == nil { // Linux/macOS
		if _, after, ok := strings.Cut(p, "zoneinfo/"); ok {
			return after
		}
	}
	return time.Local.String()
}

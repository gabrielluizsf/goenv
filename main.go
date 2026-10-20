package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"

	"github.com/gabrielluizsf/goenv/cmdx"
	"github.com/gabrielluizsf/goenv/env"
	"github.com/gabrielluizsf/goenv/lang"
	"github.com/gabrielluizsf/goenv/supervisor"
)

// Config holds the parsed command-line options and arguments.
type Config struct {
	Envs    cmdx.MultiFlag
	EnvFile string
	LogPath string
	Dir     string
	Fg      bool
	Super   bool
	Args    []string
}

func main() {
	cfg := parseFlags()

	if cfg.Super {
		handleSupervisor(cfg)
		return
	}

	e := buildEnvironment(cfg)
	cmd := prepareCommand(cfg, e)

	if cfg.Fg {
		runForeground(cmd)
		return
	}

	runBackground(cfg, e)
}

// parseFlags parses command-line flags and returns the configuration structure.
func parseFlags() Config {
	var cfg Config

	flag.Var(&cfg.Envs, "e", "KEY=VALUE")
	flag.StringVar(&cfg.EnvFile, "file", "", ".env file")
	flag.StringVar(&cfg.EnvFile, "env-file", "", "alias of --file")
	flag.StringVar(&cfg.LogPath, "log", "goenv.log", "log file")
	flag.StringVar(&cfg.Dir, "C", "", "working dir")
	flag.BoolVar(&cfg.Fg, "fg", false, "foreground")
	flag.BoolVar(&cfg.Super, "supervise", false, "internal")
	flag.Usage = cmdx.Usage
	flag.Parse()

	// Windows cmd.exe single quotes literals fix: --file='.env' -> '.env'
	cfg.EnvFile = strings.Trim(cfg.EnvFile, `'"`)

	cfg.Args = flag.Args()
	if len(cfg.Args) == 0 {
		cmdx.Usage()
		os.Exit(2)
	}

	return cfg
}

// handleSupervisor executes the internal supervisor mode for detached logging.
func handleSupervisor(cfg Config) {
	supervisor.Run(cfg.Args, cfg.LogPath, cfg.Dir)
}

// buildEnvironment compiles the environment variables from the system, .env file, and flags.
func buildEnvironment(cfg Config) []string {
	e := os.Environ()
	if cfg.EnvFile != "" {
		kv, err := env.ParseFile(cfg.EnvFile)
		if err != nil {
			cmdx.Fatal(lang.Translate("error reading %s: %v", "erro lendo %s: %v"), cfg.EnvFile, err)
		}
		e = env.Merge(e, kv)
	}
	return env.Merge(e, cfg.Envs)
}

// prepareCommand sets up the command to be executed with its environment and working directory.
func prepareCommand(cfg Config, e []string) *exec.Cmd {
	cmd := exec.Command(cfg.Args[0], cfg.Args[1:]...)
	cmd.Env = e
	cmd.Dir = cfg.Dir
	return cmd
}

// runForeground executes the command in the foreground, attaching standard input/output/error.
func runForeground(cmd *exec.Cmd) {
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		cmdx.Fatal(lang.Translate("failed to start: %v", "falha ao iniciar: %v"), err)
	}
	cmdx.PrintSummary(cmd.Process.Pid, "", false)
	_ = cmd.Wait()
}

// runBackground validates the log file and re-executes the binary as a detached supervisor.
func runBackground(cfg Config, e []string) {
	lf, err := os.OpenFile(cfg.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		cmdx.Fatal(lang.Translate("error opening log: %v", "erro abrindo log: %v"), err)
	}
	lf.Close()

	exe, err := os.Executable()
	if err != nil {
		cmdx.Fatal(lang.Translate("failed to start: %v", "falha ao iniciar: %v"), err)
	}

	supArgs := append([]string{"--supervise", "--log=" + cfg.LogPath, "-C=" + cfg.Dir, "--"}, cfg.Args...)
	sup := exec.Command(exe, supArgs...)
	sup.Env = e
	cmdx.Cmd.Detach(sup)

	if err := sup.Start(); err != nil {
		cmdx.Fatal(lang.Translate("failed to start: %v", "falha ao iniciar: %v"), err)
	}

	pid := sup.Process.Pid
	_ = sup.Process.Release()
	cmdx.PrintSummary(pid, cfg.LogPath, true)
}

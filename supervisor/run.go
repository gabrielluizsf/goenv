package supervisor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gabrielluizsf/goenv/cmdx"
	"github.com/gabrielluizsf/goenv/lang"
)

const tsLayout = "2006-01-02 15:04:05.000"

// Run roda desanexado: executa o app, captura stdout/stderr
// e grava cada linha no log com data/hora. O PID impresso pelo goenv é o dele.
func Run(args []string, logPath, dir string) {
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		os.Exit(1)
	}
	defer lf.Close()

	var mu sync.Mutex
	write := func(tag, msg string) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintf(lf, "%s [%s] %s\n", time.Now().Format(tsLayout), tag, msg)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ() // já vem montado pelo goenv (.env + -e)
	cmd.Dir = dir
	cmdx.Cmd.PrepareChild(cmd)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		write("goenv", fmt.Sprintf(lang.Translate("failed to start: %v", "falha ao iniciar: %v"), err))
		os.Exit(1)
	}
	write("goenv", fmt.Sprintf(lang.Translate("started: %s (pid %d)", "iniciado: %s (pid %d)"),
		strings.Join(args, " "), cmd.Process.Pid))

	// Se alguém matar só o supervisor, repassa o sinal para o app.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-sig
		write("goenv", fmt.Sprintf(lang.Translate("received %v, stopping app", "recebi %v, parando o app"), s))
		cmdx.Cmd.Terminate(cmd.Process)
	}()

	var wg sync.WaitGroup
	pump := func(r io.Reader, tag string) {
		defer wg.Done()
		br := bufio.NewReader(r)
		for {
			line, err := br.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			if line != "" || err == nil {
				write(tag, line)
			}
			if err != nil {
				return
			}
		}
	}
	wg.Add(2)
	go pump(stdout, "out")
	go pump(stderr, "err")
	wg.Wait()

	code := 0
	if err := cmd.Wait(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			code = ee.ExitCode()
		} else {
			code = 1
		}
	}
	write("goenv", fmt.Sprintf(lang.Translate("app exited (code %d)", "app encerrou (código %d)"), code))
	os.Exit(code)
}

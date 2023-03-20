package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"log"
)

var (
	cmd *exec.Cmd
)

func main() {

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		for {
			sig := <-s
			cmd.Process.Signal(sig)
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)

			cmd.Process.Signal(syscall.SIGTERM)
		}
	}()

	for {
		log.Printf("I will start sub progress")
		cmd = exec.Command("./internal/internal")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			log.Printf("Run with error %s", err.Error())
		}

		exitCode := cmd.ProcessState.ExitCode()
		if exitCode == -1 {
			log.Printf("sub exit with code %v", exitCode)
		}
	}
}

package cmd

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
)

type Command struct {
	command []string
	Stdout  io.Reader
	Stderr  *strings.Builder
}

func NewCommand(command []string) *Command {
	return &Command{command: command, Stderr: new(strings.Builder)}
}

// Execute runs the command locally
func (c *Command) Execute(outchan, errchan *chan string) error {
	command := exec.Command(c.command[0], c.command[1:]...)
	stdout, _ := command.StdoutPipe()
	stderr, _ := command.StderrPipe()
	go redirectStandard(stdout, outchan)
	go redirectStandard(stderr, errchan)
	return command.Run()
}

func redirectStandard(std io.Reader, to *chan string) {
	if to != nil {
		scanner := bufio.NewScanner(std)
		for scanner.Scan() {
			*to <- scanner.Text()
		}
	}
}

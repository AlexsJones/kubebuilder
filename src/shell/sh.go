package sh

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/AlexsJones/kubebuilder/src/log"
)

func RunCommand(path string, commands string) error {
	//Run build commands
	cmd := exec.Command("/bin/bash", "-c", commands)
	cmd.Dir = path
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			logger.GetInstance().Log(fmt.Sprintf(scanner.Text()))
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}
	return nil
}

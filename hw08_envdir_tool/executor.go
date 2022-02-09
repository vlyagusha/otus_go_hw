package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 1
	}
	command := exec.Command(cmd[0], cmd[1:]...) // nolint:gosec

	for key, value := range env {
		if value.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				continue
			}
		} else {
			err := os.Setenv(key, value.Value)
			if err != nil {
				continue
			}
		}
	}

	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok { // nolint:errorlint
			return exitError.ExitCode()
		}
	}

	return 0
}

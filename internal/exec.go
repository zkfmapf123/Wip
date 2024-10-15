package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func MustRunCommandInPath(path string, command string, args ...string) (string, error) {

	if err := os.Chdir(path); err != nil {
		return "", fmt.Errorf("failed chdir %s", path)
	}

	fmt.Println(path, command, args)

	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

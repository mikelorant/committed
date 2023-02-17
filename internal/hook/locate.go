package hook

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	gitCommand        = "git"
	gitGlobalArgs     = []string{"config", "--get", "core.hooksPath"}
	gitRepositoryArgs = []string{"rev-parse", "--absolute-git-dir"}
)

var ErrLocation = errors.New("no hook location found")

func Locate(run Runner) (string, error) {
	glob, _ := runCmd(run, gitCommand, gitGlobalArgs)

	if glob != "" {
		return glob, nil
	}

	repo, _ := runCmd(run, gitCommand, gitRepositoryArgs)

	if repo != "" {
		return repo, nil
	}

	return "", ErrLocation
}

func runCmd(run Runner, cmd string, args []string) (string, error) {
	var buf bytes.Buffer

	if err := run(&buf, cmd, args); err != nil {
		return "", fmt.Errorf("unable to run command: %w", err)
	}

	out, err := io.ReadAll(&buf)
	if err != nil {
		return "", fmt.Errorf("unable to read buffer: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}

package shell

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os/exec"

	"github.com/creack/pty"
)

func Run(w io.Writer, command string, args []string) error {
	cmd := exec.Command(command, args...)
	fh, err := pty.Start(cmd)
	if err != nil {
		var execError *exec.Error

		if errors.As(err, &execError) {
			return fmt.Errorf("unable to exec command: %v: %w", execError.Name, execError.Err)
		}

		return fmt.Errorf("unable to exec command: %w", err)
	}
	defer fh.Close()

	if _, err = io.Copy(w, fh); err != nil {
		var pathError *fs.PathError

		if !errors.As(err, &pathError) {
			return fmt.Errorf("unable to copy commit output: %w", err)
		}

		if pathError.Path != "/dev/ptmx" {
			return fmt.Errorf("unable to copy commit output: %v: %w", pathError.Path, pathError.Err)
		}
	}

	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError

		if errors.As(err, &exitErr) {
			return fmt.Errorf("non-zero exit code returned: %w", err)
		}

		return fmt.Errorf("unable to exec command: %w", err)
	}

	return nil
}

package hook

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func (h *Hook) Install() error {
	loc, err := h.Locater(h.Runner)
	if err != nil {
		return fmt.Errorf("unable to determine hook location: %w", err)
	}

	if loc == "" {
		return ErrLocation
	}

	h.Location = path.Join(loc, GitHook)

	manage, err := h.manage()
	if err != nil {
		return fmt.Errorf("unable to determine managed state: %w", err)
	}

	if !manage {
		return ErrUnmanaged
	}

	_, err = h.file.WriteString(PrepareGitMessage)
	if err != nil {
		return fmt.Errorf("unable to write message: %w", err)
	}

	return nil
}

func (h *Hook) manage() (bool, error) {
	managed, err := h.isManaged()
	if err != nil {
		return false, err
	}

	if !managed {
		return false, nil
	}

	fh, err := h.Creator(h.Location, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return false, fmt.Errorf("unable to create file: %w", err)
	}

	h.file = fh

	return true, nil
}

func (h *Hook) isManaged() (bool, error) {
	if !h.exists() {
		return true, nil
	}

	fh, err := h.Opener(h.Location)
	if err != nil {
		return false, fmt.Errorf("unable to open file: %w", err)
	}

	return checkSignature(fh)
}

func (h *Hook) exists() bool {
	_, err := h.Stater(h.Location)

	return err == nil
}

func checkSignature(fh io.ReadWriter) (bool, error) {
	var line string

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line = scanner.Text()
		break
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("unable to scan file: %w", err)
	}

	if strings.Contains(line, Marker) {
		return true, nil
	}

	return false, nil
}

package hook

import (
	"fmt"
	"path"
)

func (h *Hook) Uninstall() error {
	loc, err := h.Locater(h.Runner)
	if err != nil {
		return fmt.Errorf("unable to determine hook location: %w", err)
	}

	if loc == "" {
		return ErrLocation
	}

	h.Location = path.Join(loc, GitHook)

	manage, err := h.unmanage()
	if err != nil {
		return fmt.Errorf("unable to determine managed state: %w", err)
	}

	if !manage {
		return ErrUnmanaged
	}

	return nil
}

func (h *Hook) unmanage() (bool, error) {
	managed, err := h.isManaged()
	if err != nil {
		return false, err
	}

	if !managed {
		return false, nil
	}

	if err := h.Deleter(h.Location); err != nil {
		return false, fmt.Errorf("unable to delete file: %w", err)
	}

	return true, nil
}

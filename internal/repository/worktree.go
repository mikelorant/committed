package repository

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
)

type Worktree struct {
	Status git.Status
}

func (r *Repository) Worktree() (Worktree, error) {
	var wt Worktree

	w, err := r.Worktreer.Worktree()
	if err != nil {
		return Worktree{}, fmt.Errorf("unable to get worktree: %w", err)
	}

	// Performance of internal status was poor. Replace with external call
	// which is significantly faster.
	// s, err := w.Status()
	s, err := status(w)
	if err != nil {
		return Worktree{}, fmt.Errorf("unable to get status of worktree: %w", err)
	}
	wt.Status = s

	return wt, nil
}

// Alternative method to determine file status. Modified from original
// version which was part of the following pull request.
// https://github.com/zricethezav/gitleaks/pull/463
func status(wt *git.Worktree) (git.Status, error) {
	c := exec.Command("git", "status", "--porcelain", "-z")
	c.Dir = wt.Filesystem.Root()

	out, err := c.Output()
	if err != nil {
		return wt.Status()
	}

	lines := strings.Split(string(out), "\000")
	status := make(map[string]*git.FileStatus, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		ltrim := strings.TrimLeft(line, " ")

		pathStatusCode := strings.SplitN(ltrim, " ", 2)
		if len(pathStatusCode) != 2 {
			continue
		}

		statusCode := []byte(pathStatusCode[0])[0]
		path := strings.Trim(pathStatusCode[1], " ")

		status[path] = &git.FileStatus{
			Staging: git.StatusCode(statusCode),
		}
	}

	return status, err
}

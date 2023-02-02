package commit

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

func FileOpen() func(string) (io.Reader, error) {
	return func(file string) (io.Reader, error) {
		var pathError *fs.PathError

		fh, err := os.Open(os.ExpandEnv(file))
		switch {
		case err == nil:
		case errors.As(err, &pathError):
			return strings.NewReader(""), nil
		default:
			return nil, fmt.Errorf("unable to open file: %w", err)
		}

		return fh, nil
	}
}

func FileCreate() func(string) (io.WriteCloser, error) {
	return func(file string) (io.WriteCloser, error) {
		err := os.MkdirAll(path.Dir(os.ExpandEnv(file)), 0o755)
		switch {
		case err == nil:
		case os.IsExist(err):
		default:
			return nil, err
		}

		fh, err := os.Create(os.ExpandEnv(file))
		if err != nil {
			return nil, fmt.Errorf("unable to create file: %w", err)
		}

		return fh, nil
	}
}

func FileExists(file string) bool {
	_, err := os.Stat(os.ExpandEnv(file))

	return !errors.Is(err, os.ErrNotExist)
}

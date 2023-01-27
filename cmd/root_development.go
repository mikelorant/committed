//go:build !release

package cmd

func isDryRun() bool {
	return true
}

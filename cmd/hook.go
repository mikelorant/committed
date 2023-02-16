package cmd

import (
	"github.com/mikelorant/committed/internal/hook"
	"github.com/spf13/cobra"
)

func NewHookCmd(a App) *cobra.Command {
	var hookOptions hook.Options

	cmd := &cobra.Command{
		Use:   "hook",
		Short: "Install and uninstall Git hook",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if ok := help(cmd, hookOptions); ok {
				return
			}

			if err := a.Hooker.Do(hookOptions); err != nil {
				a.Logger.Fatalf("Unable to install or uninstall hook.")
			}
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().BoolVar(&hookOptions.Install, "install", false, "Install Git hook")
	cmd.Flags().BoolVar(&hookOptions.Uninstall, "uninstall", false, "Uninstall Git hook")
	cmd.Flags().Lookup("install").NoOptDefVal = "true"
	cmd.Flags().Lookup("uninstall").NoOptDefVal = "true"

	return cmd
}

func help(cmd *cobra.Command, opts hook.Options) bool {
	if !(opts.Install || opts.Uninstall) {
		cmd.Help()

		return true
	}

	return false
}

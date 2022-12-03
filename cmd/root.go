package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mikelorant/committed/internal/ui"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committed",
		Short: "A brief description of your application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ui.New(); err != nil {
				log.Fatalf("unable to init ui: %v", err)
			}

			return nil
		},
	}

	cc.Init(&cc.Config{
		RootCmd:         cmd,
		Headings:        cc.HiGreen + cc.Bold,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	return cmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

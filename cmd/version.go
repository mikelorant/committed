package cmd

import (
	_ "embed"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	ReleaseVersion = "snapshot"
	ReleaseDate    = "unknown"
	ReleaseCommit  = "none"
)

//go:embed version.gotmpl
var verTmpl string

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:         "version",
		Short:       "Print the version information",
		Annotations: annotations(),
		Run: func(cmd *cobra.Command, args []string) {
			tmpl, err := template.New("version").Parse(verTmpl)
			if err != nil {
				log.Fatal("Unable to parse version template.")
			}
			if err = tmpl.Execute(os.Stdout, cmd); err != nil {
				log.Fatal("Unable to show version.")
			}
		},
	}
}

func annotations() map[string]string {
	return map[string]string{
		"version": ReleaseVersion,
		"date":    ReleaseDate,
		"commit":  ReleaseCommit,
	}
}

package cmd

import (
	_ "embed"
	"log"
	"runtime/debug"
	"text/template"

	"github.com/spf13/cobra"
)

var version = "snapshot"

//go:embed version.gotmpl
var verTmpl string

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:         "version",
		Short:       "Print the version information",
		Annotations: annotations(),
		Run: func(cmd *cobra.Command, args []string) {
			tmpl := template.Must(template.New("version").Parse(verTmpl))
			if err := tmpl.Execute(cmd.OutOrStdout(), cmd); err != nil {
				log.Fatal("Unable to show version.")
			}
		},
	}
}

func annotations() map[string]string {
	var (
		date   = "unknown"
		commit = "none"
	)

	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalf("unable to read build info")
	}

	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.time":
			date = s.Value
		case "vcs.revision":
			commit = s.Value
		}
	}

	return map[string]string{
		"version": version,
		"date":    date,
		"commit":  commit,
	}
}

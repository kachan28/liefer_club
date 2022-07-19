package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	pdf  = "pdf"
	json = "json"
)

func Export() *cobra.Command {
	protocolCommand := &cobra.Command{
		Use:   "export",
		Short: "export all changes in database",
		Run: func(cmd *cobra.Command, args []string) {
			runExport(args)
		},
	}

	return protocolCommand
}

func runExport(args []string) {
	for _, arg := range args {
		if arg == pdf || arg == json {
			fmt.Printf("export run to %s \n", arg)
		} else {
			fmt.Println("unknown method for export, skip")
		}
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	return &cobra.Command{
		Use:   "protocoller",
		Short: "protocoller is a tool, which allows to protocol and export database changes",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
}

func Execute() {
	rootCmd := Root()
	rootCmd.AddCommand(Protocol())
	rootCmd.AddCommand(Export())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

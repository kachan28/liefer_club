package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Protocol() *cobra.Command {
	protocolCommand := &cobra.Command{
		Use:   "protocol",
		Short: "protocol all changes in database",
		Run: func(cmd *cobra.Command, args []string) {
			runProtocol()
		},
	}

	return protocolCommand
}

func runProtocol() {
	fmt.Println("protocol run")
}

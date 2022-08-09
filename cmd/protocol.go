package cmd

import (
	"log"

	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/services/protocol"
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
	conf := new(app.Conf)
	conf.GetConf()
	protocolService := protocol.MakeProtocolService()
	err := protocolService.MakeProtocol(conf)
	if err != nil {
		log.Fatal("can't make protocol - ", err)
	}
}

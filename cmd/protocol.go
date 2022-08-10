package cmd

import (
	"fmt"
	"log"

	"github.com/kachan28/liefer_club/app"
	fileService "github.com/kachan28/liefer_club/internal/services/file"
	"github.com/kachan28/liefer_club/internal/services/protocol"
	"github.com/spf13/cobra"
)

func Protocol() *cobra.Command {
	protocolCommand := &cobra.Command{
		Use:   "protocol",
		Short: "protocol all changes in database",
		Run: func(cmd *cobra.Command, args []string) {
			runProtocol(args)
		},
	}

	return protocolCommand
}

func runProtocol(args []string) {
	conf := new(app.Conf)
	conf.GetConf()
	for _, arg := range args {
		if arg == "list" {
			dts, err := fileService.FileService{}.GetFileCreateDts()
			if err != nil {
				log.Fatal("can't get dumps dts - ", err)
			}
			for _, dt := range dts {
				fmt.Println(dt)
			}
			return
		}
	}
	protocolService := protocol.MakeProtocolService()
	err := protocolService.MakeProtocol(conf)
	if err != nil {
		log.Fatal("can't make protocol - ", err)
	}
}

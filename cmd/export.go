package cmd

import (
	"fmt"
	"log"

	"github.com/kachan28/liefer_club/internal/services/export"
	"github.com/spf13/cobra"
)

const (
	pdf  = "pdf"
	json = "json"

	langFlag     = "l"
	dumpDateFlag = "date"
	dumpTimeFlag = "time"
)

func Export() *cobra.Command {
	protocolCommand := &cobra.Command{
		Use:   "export",
		Short: "export all changes in database",
		Run: func(cmd *cobra.Command, args []string) {
			exportLang, _ := cmd.Flags().GetString(langFlag)
			if exportLang == "" {
				log.Fatal("please specify export language")
			}
			dumpDate, _ := cmd.Flags().GetString(dumpDateFlag)
			dumpTime, _ := cmd.Flags().GetString(dumpTimeFlag)
			runExport(args, exportLang, dumpDate, dumpTime)
		},
	}

	protocolCommand.PersistentFlags().String(langFlag, "", "select lang")
	protocolCommand.PersistentFlags().String(dumpDateFlag, "", "select date of dump")
	protocolCommand.PersistentFlags().String(dumpTimeFlag, "", "select time of dump")

	return protocolCommand
}

func runExport(args []string, exportLang string, date string, time string) {
	if len(args) == 0 {
		log.Fatal("please enter formats for export")
	}
	for _, arg := range args {
		if arg == pdf {
			fmt.Printf("export run to %s \n", arg)
			err := export.ExportService{}.ExportToPdf(exportLang, date, time)
			if err != nil {
				log.Fatalf("can't export to pdf with err: %v", err)
			}
		} else {
			fmt.Println("unknown method for export, skip")
		}
	}
}

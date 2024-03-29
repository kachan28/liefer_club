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
)

func Export() *cobra.Command {
	protocolCommand := &cobra.Command{
		Use:   "export",
		Short: "export all changes in database",
		Run: func(cmd *cobra.Command, args []string) {
			exportLang, _ := cmd.Flags().GetString("l")
			if exportLang == "" {
				log.Fatal("please specify export language")
			}
			runExport(args, exportLang)
		},
	}

	protocolCommand.PersistentFlags().String("l", "", "select lang")

	return protocolCommand
}

func runExport(args []string, exportLang string) {
	if len(args) == 0 {
		log.Fatal("please enter formats for export")
	}
	for _, arg := range args {
		if arg == pdf {
			fmt.Printf("export run to %s \n", arg)
			err := export.ExportService{}.ExportToPdf(exportLang)
			if err != nil {
				log.Fatalf("can't export to pdf with err: %v", err)
			}
		} else {
			fmt.Println("unknown method for export, skip")
		}
	}
}

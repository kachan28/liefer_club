package export

import (
	"github.com/kachan28/liefer_club/internal/services/file"
)

type ExportService struct{}

func (e ExportService) ExportToPdf() error {
	exportConfig, err := GetExportConfigService{}.GetConfig()
	if err != nil {
		return err
	}
	protocolTitle, err := file.FileService{}.GetLastProtocol()
	if err != nil {
		return err
	}
	result, err := file.FileService{}.ReadProtocol(protocolTitle)
	if err != nil {
		return err
	}
	err = CreatePdfProtocol{}.CreatePdfFile(*result, *exportConfig, "pdfs/example.pdf")
	if err != nil {
		return err
	}

	return nil
}

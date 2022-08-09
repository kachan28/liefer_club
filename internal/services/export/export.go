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
	protocolTitle, _, err := file.FileService{}.GetLastProtocol()
	if err != nil {
		return err
	}
	result, err := file.FileService{}.ReadProtocol(protocolTitle)
	if err != nil {
		return err
	}
	exportDir, err := file.FileService{}.GetExportDirectory()
	if err != nil {
		return nil
	}
	err = CreatePdfProtocol{}.CreatePdfFile(*result, *exportConfig, file.FileService{}.SetFullExportFilename(exportDir, protocolTitle))
	if err != nil {
		return err
	}

	return nil
}

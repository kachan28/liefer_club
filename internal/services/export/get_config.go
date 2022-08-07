package export

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kachan28/liefer_club/internal/models"
)

const exportConfigPath = "export_config.json"

type GetExportConfigService struct{}

func (g GetExportConfigService) GetConfig() (*models.ExportConfig, error) {
	exportConfig := new(models.ExportConfig)
	currPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(filepath.Join(currPath, exportConfigPath))
	err = json.Unmarshal(content, exportConfig)
	exportConfig.PrepareLabels()
	return exportConfig, nil
}

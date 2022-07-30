package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	protokolsFolder    = "LC-Kasse-Programmierprotokolle"
	protokolMenuFolder = "JSON-Format"
)

var (
	protokolPath = filepath.Join(protokolsFolder, protokolMenuFolder)
)

type FileService struct{}

func (f FileService) WriteProtokol(result models.ResultModel) error {
	currPath, err := os.Getwd()
	if err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(currPath, protokolPath)); os.IsNotExist(err) {
		err := os.Mkdir(protokolPath, 0755)
		if err != nil {
			return err
		}
	}
	jsonResult, err := json.Marshal(&result)
	if err != nil {
		return err
	}
	protocolFileCreationTime := time.Now()
	protocolFileName := fmt.Sprintf("%s_%s", result.Company.Name, protocolFileCreationTime.Format("2006.01.02 15:04:05"))
	err = os.WriteFile(filepath.Join(currPath, protokolPath, protocolFileName)+".json", jsonResult, 0644)
	if err != nil {
		return err
	}

	return err
}

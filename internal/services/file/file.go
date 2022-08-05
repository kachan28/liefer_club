package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	programmFilesFolder = "LC-Kasse-Programmierprotokolle"
	protokolMenuFolder  = "JSON-Format"
)

var (
	protokolPath = filepath.Join(programmFilesFolder, protokolMenuFolder)
)

type FileService struct{}

func (f FileService) WriteProtokol(result models.ResultModel) error {
	currPath, err := os.Getwd()
	if err != nil {
		return err
	}
	absoluteProgramPath := filepath.Join(currPath, protokolPath)

	if _, err := os.Stat(absoluteProgramPath); os.IsNotExist(err) {
		err := os.MkdirAll(absoluteProgramPath, 0755)
		if err != nil {
			return err
		}
	}
	jsonResult, err := json.Marshal(&result)
	if err != nil {
		return err
	}

	dt, err := result.GetCreationTimeForFile()
	if err != nil {
		return err
	}

	protocolFileName := fmt.Sprintf("%s_%s", strings.ReplaceAll(result.Company.Name, " ", "_"), dt)
	err = os.WriteFile(filepath.Join(absoluteProgramPath, protocolFileName)+".json", jsonResult, 0644)
	if err != nil {
		return err
	}

	return err
}

package file

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kachan28/liefer_club/internal/models"
)

const (
	protokolFolder = "LC-Kasse-Programmierprotokolle/JSON-Format"
)

type FileService struct{}

func (f FileService) WriteProtokol(result models.ResultModel) error {
	if _, err := os.Stat(fmt.Sprintf("./%s", protokolFolder)); os.IsNotExist(err) {
		err := os.Mkdir(fmt.Sprintf("%s", protokolFolder), 0755)
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
	err = os.WriteFile(fmt.Sprintf("%s/%s.json", protokolFolder, protocolFileName), jsonResult, 0644)
	if err != nil {
		return err
	}

	return err
}

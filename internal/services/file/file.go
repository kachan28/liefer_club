package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kachan28/liefer_club/internal/models"
	timeService "github.com/kachan28/liefer_club/internal/services/time"
)

const (
	programmFilesFolder = "LC-Kasse-Programmierprotokolle"
	protokolMenuFolder  = "JSON-Format"
	jsonFormat          = ".json"
	pdfFormat           = ".pdf"
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

	protocolFileName := fmt.Sprintf("%s,%s", strings.ReplaceAll(result.Company.Name, " ", "_"), dt)
	err = os.WriteFile(filepath.Join(absoluteProgramPath, protocolFileName)+jsonFormat, jsonResult, 0644)
	if err != nil {
		return err
	}

	return err
}

func (f FileService) GetLastProtocol() (string, time.Time, error) {
	currPath, err := os.Getwd()
	if err != nil {
		return "", time.Time{}, err
	}
	files, err := ioutil.ReadDir(filepath.Join(currPath, protokolPath))
	if err != nil {
		return "", time.Time{}, err
	}
	if len(files) == 0 {
		return "", time.Time{}, fmt.Errorf("no files in json protocols folder")
	}
	var creationDt time.Time
	var lastCreatedProtocolName string
	for _, file := range files {
		fileName := file.Name()
		dateTimeAndFormat := strings.Split(fileName, ",")
		dateTime, err := timeService.TimeService{}.GetTimeFromFileTitle(strings.Split(dateTimeAndFormat[1], jsonFormat)[0])
		if err != nil {
			return "", time.Time{}, err
		}
		if dateTime.After(creationDt) {
			creationDt = dateTime
			lastCreatedProtocolName = fileName
		}
	}
	return lastCreatedProtocolName, creationDt, nil
}

func (f FileService) ReadProtocol(protocolFileName string) (*models.ResultModel, error) {
	currPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	absoluteProgramPath := filepath.Join(currPath, protokolPath)
	content, err := ioutil.ReadFile(filepath.Join(absoluteProgramPath, protocolFileName))
	if err != nil {
		return nil, err
	}
	result := new(models.ResultModel)
	err = json.Unmarshal(content, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (f FileService) GetExportDirectory() (string, error) {
	currPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(currPath, programmFilesFolder), nil
}

func (f FileService) SetFullExportFilename(folderPath, fileName string) string {
	return filepath.Join(folderPath, fileName) + pdfFormat
}

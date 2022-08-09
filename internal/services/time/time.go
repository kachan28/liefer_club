package time

import (
	"time"
)

const (
	formatForProtocol = "2006-01-02 15:04:05"
	formatForFile     = "2006.01.02_15-04-05"
	formatForExport   = "15:04:05 02.01.2006"
	formatForCsv      = "2006.01.02_15-04-05"
)

type TimeService struct{}

func (t TimeService) GetTimeForProtocol() string {
	return time.Now().Format(formatForProtocol)
}

func (t TimeService) GetCreationTimeForFile(creationDate string) (string, error) {
	protocolFileCreationTime, err := time.Parse(formatForProtocol, creationDate)
	if err != nil {
		return "", err
	}
	return protocolFileCreationTime.Format(formatForFile), nil
}

func (t TimeService) GetTimeFromFileTitle(creationTime string) (time.Time, error) {
	protocolFileCreationTime, err := time.Parse(formatForFile, creationTime)
	if err != nil {
		return time.Now(), err
	}
	return protocolFileCreationTime, nil
}

func (t TimeService) GetDateStringForExport(creationTime string) (string, error) {
	protocolFileCreationTime, err := time.Parse(formatForProtocol, creationTime)
	if err != nil {
		return "", err
	}
	return protocolFileCreationTime.Format(formatForExport), nil
}

func (t TimeService) GetUpdateDatabaseDtFromCsv(creationTime string) (time.Time, error) {
	dbUpdateDt, err := time.Parse(formatForCsv, creationTime)
	if err != nil {
		return time.Time{}, err
	}
	return dbUpdateDt, nil
}

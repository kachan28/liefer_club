package time

import (
	"time"
)

const (
	formatForProtocol      = "2006-01-02 15:04:05"
	formatForFile          = "2006.01.02_15-04-05"
	formatForExport        = "15:04:05 02.01.2006"
	formatForCsv           = "2006.01.02_15-04-05"
	formatForInputOnlyDate = "2006.01.02"
	formatForInput         = "2006.01.02_15-04-05"
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

func (t TimeService) GetExportDateFromInput(date string) (time.Time, error) {
	exportDt, err := time.Parse(formatForInputOnlyDate, date)
	if err != nil {
		return time.Time{}, err
	}
	return exportDt, err
}

func (t TimeService) GetExportDateTimeFromInput(date, etime string) (time.Time, error) {
	exportDt, err := time.Parse(formatForInput, date+"_"+etime)
	if err != nil {
		return time.Time{}, err
	}
	return exportDt, nil
}

func (t TimeService) IsDateEqual(t1, t2 time.Time) bool {
	return t1.Day() == t2.Day() && t1.Month() == t2.Month() && t1.Year() == t1.Year()
}

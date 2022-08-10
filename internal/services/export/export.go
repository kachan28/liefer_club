package export

import (
	"fmt"

	timePkg "time"

	"github.com/kachan28/liefer_club/internal/services/file"
	timeService "github.com/kachan28/liefer_club/internal/services/time"
)

type ExportService struct{}

func (e ExportService) ExportToPdf(exportLang, date, time string) error {
	exportConfig, err := GetExportConfigService{}.GetConfig()
	if err != nil {
		return err
	}

	dtFound := false

	var protocolTitle string
	if date != "" {
		timesMap, err := file.FileService{}.GetDumpsDts()
		if err != nil {
			return err
		}
		if time != "" {
			dt, err := timeService.TimeService{}.GetExportDateTimeFromInput(date, time)
			if err != nil {
				return err
			}
			if _, ok := timesMap[dt]; ok {
				dtFound = true
			}
			protocolTitle = timesMap[dt]
		} else {
			exportDt := timePkg.Time{}
			minsub := 1000000000000000000
			date, err := timeService.TimeService{}.GetExportDateFromInput(date)
			if err != nil {
				return err
			}
			for dt := range timesMap {
				if date.After(dt) {
					dtFound = true
					if date.Sub(dt) < timePkg.Duration(minsub) {
						exportDt = dt
						minsub = int(date.Sub(dt))
					}
				}
			}
			for dt := range timesMap {
				if (timeService.TimeService{}.IsDateEqual(date, dt)) {
					dtFound = true
					if dt.After(exportDt) {
						exportDt = dt
					}
				}
			}
			protocolTitle = timesMap[exportDt]
		}
	}

	if !dtFound {
		fmt.Println("datetime was not provided or can't find rigth dt, getting last")
		protocolTitle, _, err = file.FileService{}.GetLastProtocol()
		if err != nil {
			return err
		}
	}
	result, err := file.FileService{}.ReadProtocol(protocolTitle)
	if err != nil {
		return err
	}
	exportDir, err := file.FileService{}.GetExportDirectory()
	if err != nil {
		return nil
	}
	err = CreatePdfProtocol{}.CreatePdfFile(*result, *exportConfig, file.FileService{}.SetFullExportFilename(exportDir, protocolTitle), exportLang)
	if err != nil {
		return err
	}

	return nil
}

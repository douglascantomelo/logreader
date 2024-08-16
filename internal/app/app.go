package app

import (
	"logreader/internal/repository"
	"logreader/internal/service"
)

func Run() error {
	logRepo := repository.NewLogRepository()
	logService := service.NewLogService(logRepo)

	matchData, err := logService.ParseLogs()
	if err != nil {
		return err
	}

	reportService := service.NewReportService(matchData)
	reportService.PrintReport()

	return nil
}

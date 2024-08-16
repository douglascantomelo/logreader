package service

import (
	"logreader/internal/domain/report"
)

type ReportService struct {
	matchData map[string]report.MatchData
}

func NewReportService(matchData map[string]report.MatchData) *ReportService {
	return &ReportService{matchData: matchData}
}

func (s *ReportService) PrintReport() {
	report.FormatReport(s.matchData)
}

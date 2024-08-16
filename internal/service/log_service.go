package service

import (
	"fmt"
	"logreader/internal/domain/report"
	"logreader/internal/repository"
	"strings"
)

type LogService struct {
	logRepo *repository.LogRepository
}

func NewLogService(logRepo *repository.LogRepository) *LogService {
	return &LogService{logRepo: logRepo}
}

func (s *LogService) ParseLogs() (map[string]report.MatchData, error) {
	lines, err := s.logRepo.ReadLogFile()
	if err != nil {
		return nil, err
	}

	matches := make(map[string]report.MatchData)
	var currentMatch string
	var game int = 1

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		if strings.Contains(line, "InitGame") {
			currentMatch = fmt.Sprintf("game_%02d", game)
			matches[currentMatch] = report.MatchData{
				TotalKills:   0,
				Players:      []string{},
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
			game++
		} else if strings.Contains(line, "Kill") {
			killParts := strings.Split(line, ":")
			if len(killParts) < 3 {
				continue
			}

			killInfo := strings.TrimSpace(killParts[3])
			killDetails := strings.Fields(killInfo)

			if len(killDetails) < 5 {
				continue
			}

			killer, lastKillerIndex := killerInfo(killDetails)
			victim := victimInfo(lastKillerIndex, killDetails)
			meansOfDeath := killDetails[len(killDetails)-1]
			extractKills(matches, currentMatch, victim, killer, meansOfDeath)
		}
	}

	return matches, nil
}

func extractKills(matches map[string]report.MatchData, currentMatch string, victim string, killer string, meansOfDeath string) {
	matchData := matches[currentMatch]
	matchData.TotalKills++

	if !contains(matchData.Players, victim) {
		matchData.Players = append(matchData.Players, victim)
	}

	if killer != "<world>" {
		matchData.Kills[killer]++
	} else {
		matchData.Kills[victim]--
	}
	matchData.KillsByMeans[meansOfDeath]++

	matches[currentMatch] = matchData
}

func victimInfo(lastKillerIndex int, killDetails []string) string {
	victim := ""
	for i := lastKillerIndex + 1; i < len(killDetails); i++ {
		if killDetails[i] == "by" {
			break
		}
		if victim != "" {
			victim += " "
		}
		victim += killDetails[i]
	}
	return victim
}

func killerInfo(killDetails []string) (string, int) {
	killer := ""
	var lastKillerIndex int = 0
	for i := 0; i < len(killDetails); i++ {
		if killDetails[i] == "killed" {
			lastKillerIndex = i
			break
		}
		if killer != "" {
			killer += " "
		}
		killer += killDetails[i]
	}
	return killer, lastKillerIndex
}
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

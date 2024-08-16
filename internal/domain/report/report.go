package report

import (
	"encoding/json"
	"fmt"
)

// Dados de uma partida
type MatchData struct {
	TotalKills   int
	Players      []string
	Kills        map[string]int
	KillsByMeans map[string]int
}

// Função para gerar o relatório completo
func FormatReport(matches map[string]MatchData) {
	report := make(map[string]MatchData)

	for matchID, data := range matches {
		report[matchID] = data
	}

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling report: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}

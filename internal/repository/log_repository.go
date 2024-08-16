package repository

import (
	"bufio"
	"os"
	"path/filepath"
)

type LogRepository struct {
	filePath string
}

func NewLogRepository() *LogRepository {
	logFilePath := filepath.Join("logs", "qgames.log")
	return &LogRepository{filePath: logFilePath}
}

func (repo *LogRepository) ReadLogFile() ([]string, error) {
	file, err := os.Open(repo.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

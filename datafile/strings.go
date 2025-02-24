package datafile

import (
	"bufio"
	"os"
)

func GetStrings(fileName string) ([]string, error) {
	var texts []string
	file, err := os.Open(fileName)
	if err != nil {
		return texts, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		texts = append(texts, text)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return texts, nil
}

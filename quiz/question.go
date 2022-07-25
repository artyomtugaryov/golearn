package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

type Question struct {
	Expression *string
	Answer     *string
}

func loadQuiz(filePath string) (*[]Question, error) {
	var file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var result []Question

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		var expression, answer, err = parseQuizLine(line)
		if err != nil {
			return nil, err
		}

		result = append(result, Question{&expression, &answer})
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return &result, nil
}

func parseQuizLine(line string) (string, string, error) {
	if !strings.Contains(line, ",") {
		return "", "", fmt.Errorf("string '%s' does not contain ','", line)
	}

	var splits = strings.Split(line, ",")
	if len(splits) != 2 {
		return "", "", fmt.Errorf("cannot find an expression and an answer in the string %s", line)
	}

	return splits[0], splits[1], nil
}

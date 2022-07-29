package main

import (
	"strings"
	"fmt"
)

type Question struct {
	Expression string
	Answer     string
}

func creatQuestionFromLine(line string) (Question, error) {
	if !strings.Contains(line, ",") {
		return Question{}, fmt.Errorf("string '%s' does not contain ','", line)
	}

	var splits = strings.Split(line, ",")
	if len(splits) != 2 {
		return Question{}, fmt.Errorf("cannot find an expression and an answer in the string %s", line)
	}

	return Question{splits[0], splits[1]}, nil
}

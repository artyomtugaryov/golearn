package main

import (
	"testing"
	"fmt"
	"strings"
)

func TestParseQuizLine(t *testing.T) {
	var expectedExpression = "1+1"
	var expectedAnswer = "2"
	var line = fmt.Sprintf("%s,%s", expectedExpression, expectedAnswer)

	actualExpression, actualAnswer, err := parseQuizLine(line)
	if strings.Compare(expectedExpression, actualExpression) != 0 ||
		strings.Compare(expectedAnswer, actualAnswer) != 0 ||
		err != nil {
		t.Fatalf("parseQuizLine(%s) returns unexpected result: %s, %s", line, actualExpression, actualAnswer)
	}
}

func TestParseQuizLineWrongSep(t *testing.T) {
	var expectedExpression = "1+1"
	var expectedAnswer = "2"
	var line = fmt.Sprintf("%s;%s", expectedExpression, expectedAnswer)

	actualExpression, actualAnswer, err := parseQuizLine(line)
	if strings.Compare(expectedExpression, actualExpression) == 0 ||
		strings.Compare(expectedAnswer, actualAnswer) == 0 ||
		err == nil {
		t.Fatalf("parseQuizLine(%s) returns unexpected result: %s, %s", line, actualExpression, actualAnswer)
	}
}

func TestParseQuizLineTwoSep(t *testing.T) {
	var expectedExpression = "1+1"
	var expectedAnswer = "2"
	var line = fmt.Sprintf("%s,%s,%s", expectedExpression, expectedAnswer, expectedAnswer)

	actualExpression, actualAnswer, err := parseQuizLine(line)
	if strings.Compare(expectedExpression, actualExpression) == 0 ||
		strings.Compare(expectedAnswer, actualAnswer) == 0 ||
		err == nil {
		t.Fatalf("parseQuizLine(%s) returns unexpected result: %s, %s", line, actualExpression, actualAnswer)
	}
}

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

	question, err := creatQuestionFromLine(line)
	if strings.Compare(expectedExpression, question.Expression) != 0 ||
		strings.Compare(expectedAnswer, question.Answer) != 0 ||
		err != nil {
		t.Fatalf("parseQuizLine(%s) returns unexpected result: %s, %s", line, question.Expression, question.Answer)
	}
}

func TestParseQuizLineWrongSep(t *testing.T) {
	var expectedExpression = "1+1"
	var expectedAnswer = "2"
	var line = fmt.Sprintf("%s;%s", expectedExpression, expectedAnswer)

	question, err := creatQuestionFromLine(line)
	if strings.Compare(expectedExpression, question.Expression) == 0 ||
		strings.Compare(expectedAnswer, question.Answer) == 0 ||
		err == nil {
		t.Fatalf("parseQuizLine(%s) returns unexpected result: %s, %s", line, question.Expression, question.Answer)
	}
}

func TestParseQuizLineTwoSep(t *testing.T) {
	var expectedExpression = "1+1"
	var expectedAnswer = "2"
	var line = fmt.Sprintf("%s,%s,%s", expectedExpression, expectedAnswer, expectedAnswer)

	question, err := creatQuestionFromLine(line)
	if strings.Compare(expectedExpression, question.Expression) == 0 ||
		strings.Compare(expectedAnswer, question.Answer) == 0 ||
		err == nil {
		t.Fatalf("parseQuizLine(%s) returns unexpected result: %s, %s", line, question.Expression, question.Answer)
	}
}

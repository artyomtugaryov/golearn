package main

import (
	"os"
	"fmt"
	"bufio"
	"time"
)

type answerSubmitterType func(chan<- string)

func loadQuiz(filePath string) ([]Question, error) {
	var file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var result []Question

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		var question, err = creatQuestionFromLine(line)
		if err != nil {
			return nil, err
		}

		result = append(result, question)
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func runQuiz(questions []Question, timePerQuestion int, answerSubmitter answerSubmitterType) int {
	rightAnswers := 0

	answerChanel := make(chan string)

	for index, question := range questions {
		timerChanel := time.After(time.Duration(timePerQuestion) * time.Second)

		fmt.Printf("%d: Enter the answer of the expression %s=", index+1, question.Expression)
		go answerSubmitter(answerChanel)

		select {
		case actualAnswer := <-answerChanel:
			if actualAnswer == question.Answer {
				fmt.Printf("You are absolutelly right - %s.\n", actualAnswer)
				rightAnswers++
			} else {
				fmt.Printf("You are wrong (%s). The right answer is %s.\n", actualAnswer, question.Answer)
			}
		case <-timerChanel:
			fmt.Println("\nCommon bro! You are too slow!")
		}
	}
	return rightAnswers
}

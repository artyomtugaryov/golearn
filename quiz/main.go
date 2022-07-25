package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"time"
)

func getAnswer(reader *bufio.Reader, outChannel chan<- string) {
	readLine, _ := reader.ReadString('\n')
	answer := strings.Trim(readLine, "\n")

	outChannel <- answer
}

func quiz(questions []Question, timePerQuestion int) int {
	rightAnswers := 0

	answerChanel := make(chan string)
	stdAnswerReader := bufio.NewReader(os.Stdin)

	for _, question := range questions {
		timerChanel := time.After(time.Duration(timePerQuestion) * time.Second)

		fmt.Print("Enter the answer of the expression ", *question.Expression, "=")
		go getAnswer(stdAnswerReader, answerChanel)

		select {
		case actualAnswer := <-answerChanel:
			if actualAnswer == *question.Answer {
				rightAnswers++
			} else {
				fmt.Printf("You are wrong. The right answer is %s.\n", *question.Answer)
			}
		case <-timerChanel:
			fmt.Println("\nCommon bro! You are too slow!")
		}
	}
	return rightAnswers
}

func showResult(result int) {
	fmt.Printf("Good job! Your result is %d.", result)
}

func main() {

	var parser = NewCliParser()
	var cliArgs, cliErr = parser.Parse(os.Args)
	if cliErr != nil {
		fmt.Println(cliErr)
		return
	}

	var questions, loadQuizErr = loadQuiz(*cliArgs.QuestionsFilePath)
	if loadQuizErr != nil {
		fmt.Println(loadQuizErr)
		return
	}

	fmt.Println("Welcome to the quiz. I will ask you a few questions. You need to enter the right question")

	result := quiz(*questions, *cliArgs.Time)
	showResult(result)
}

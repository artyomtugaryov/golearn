package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func getAnswer(reader *bufio.Reader, outChannel chan<- string) {
	readLine, _ := reader.ReadString('\n')
	answer := strings.Trim(readLine, "\n")

	outChannel <- answer
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

	fmt.Println("Welcome to the runQuiz. I will ask you a few questions. You need to enter the right question")

	stdAnswerReader := bufio.NewReader(os.Stdin)
	var answerGetter = func(answerChannel chan<- string) { getAnswer(stdAnswerReader, answerChannel) }
	result := runQuiz(questions, *cliArgs.Time, answerGetter)
	showResult(result)
}

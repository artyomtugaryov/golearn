package main

import (
	"testing"
	"os"
	"io"
	"bufio"
)

func lineCounter(reader io.Reader) int {

	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)

	var counter int

	for fileScanner.Scan() {
		fileScanner.Text()
		counter++
	}

	return counter
}

func getFileLinesNumber(filePath string) (int, error) {
	var file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return lineCounter(file), nil
}

func TestLengthQuiz(t *testing.T) {
	var quizFilePath = os.Getenv("TEST_GOOD_QUIZ_FILE")
	var quiz, errLoadQuiz = loadQuiz(quizFilePath)
	if errLoadQuiz != nil {
		t.Fatalf("Error on loading the file %s ", quizFilePath)
	}
	var numberOfLines, errCountFileLineNumber = getFileLinesNumber(quizFilePath)
	if errCountFileLineNumber != nil {
		t.Fatalf("Cannot open file %s ", quizFilePath)
	}
	var quizLength = len(quiz)
	if quizLength != numberOfLines {
		t.Fatalf(
			"Quize lenght and number of file lines are different. Expected %d, actual %d",
			numberOfLines, quizLength,
		)
	}
}

func TestLoadQuizFromNotExistedFile(t *testing.T) {
	var quizFilePath = "not_existed.csv"
	var quiz, errLoadQuiz = loadQuiz(quizFilePath)
	if quiz != nil {
		t.Fatalf("Quiz is not nil for not existed file: %v", quiz)
	}
	if errLoadQuiz == nil {
		t.Fatal("There is not error for loading not existed file")
	}
}

func TestLoadQuizFromBadFile(t *testing.T) {
	var quizFilePath = os.Getenv("TEST_BAD_QUIZ_FILE")
	var quiz, errLoadQuiz = loadQuiz(quizFilePath)
	if quiz != nil {
		t.Fatalf("Quiz is not nil for the bad file: %v", quiz)
	}
	if errLoadQuiz == nil {
		t.Fatal("There is not error for the bad file")
	}

}

func testAnswerSubmitter(answers *[]string, answerSubmitChannel chan<- string) {
	answerSubmitChannel <- (*answers)[0]
	*answers = (*answers)[1:]
}

func TestQuizWithRightAnswers(t *testing.T) {
	var quizFilePath = os.Getenv("TEST_GOOD_QUIZ_FILE")
	var questions, errLoadQuiz = loadQuiz(quizFilePath)
	if errLoadQuiz != nil {
		t.Fatalf("Cannot load runQuiz from the file %s", quizFilePath)
	}
	var answers = make([]string, len(questions))
	for index, comp := range questions {
		answers[index] = comp.Answer
	}
	var answerSubmitter = func(answerChannel chan<- string) { testAnswerSubmitter(&answers, answerChannel) }
	result := runQuiz(questions, 1, answerSubmitter)
	if result != len(questions) {
		t.Fatalf("The result differ (%d) with expected (%d)", result, len(questions))
	}
}

func TestQuizWithWrongAnswers(t *testing.T) {
	var quizFilePath = os.Getenv("TEST_GOOD_QUIZ_FILE")
	var questions, errLoadQuiz = loadQuiz(quizFilePath)
	if errLoadQuiz != nil {
		t.Fatalf("Cannot load run the Quiz from the file %s", quizFilePath)
	}
	var answers = make([]string, len(questions))
	for index, comp := range questions {
		answers[index] = comp.Answer
	}
	answers[0] = "wrong"
	var answerSubmitter = func(answerChannel chan<- string) { testAnswerSubmitter(&answers, answerChannel) }
	result := runQuiz(questions, 1, answerSubmitter)
	if result != len(questions)-1 {
		t.Fatalf("The result differ (%d) with expected (%d)", result, len(questions)-1)
	}
}

func TestQuizWithoutAnswers(t *testing.T) {
	var quizFilePath = os.Getenv("TEST_GOOD_QUIZ_FILE")
	var questions, errLoadQuiz = loadQuiz(quizFilePath)
	if errLoadQuiz != nil {
		t.Fatalf("Cannot load run the Quiz from the file %s", quizFilePath)
	}
	result := runQuiz(questions[0:2], 1, func(_ chan<- string) {})
	if result != 0 {
		t.Fatalf("The result of the quiz is not 0, but %d", result)
	}
}

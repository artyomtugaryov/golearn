package main

import (
	"github.com/akamensky/argparse"
)

type CliArguments struct {
	Time              *int
	QuestionsFilePath *string
}

type CliParser struct {
	argparse.Parser
	arguments CliArguments
}

func NewCliParser() *CliParser {
	parser := &CliParser{
		*argparse.NewParser("Quiz", "Test you for solving simple math problems"),
		CliArguments{},
	}

	parser.arguments.Time = parser.Int(
		"t", "time",
		&argparse.Options{
			Required: false,
			Help:     "Time for a one question.",
			Default:  3,
		})

	parser.arguments.QuestionsFilePath = parser.String(
		"q", "questions-file",
		&argparse.Options{
			Required: false,
			Help:     "File with questions in the CSV format",
			Default:  "./questions.csv",
		})
	return parser
}

func (p *CliParser) Parse(args []string) (*CliArguments, error) {
	if err := p.Parser.Parse(args); err != nil {
		return nil, err
	}
	return &p.arguments, nil
}

package main

import "github.com/aws/aws-lambda-go/lambda"

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

//show initialised and returns a new book
func show() (*book, error) {
	bk := &book{
		ISBN:   "978-1420931693",
		Title:  "The Republic",
		Author: "Plato",
	}

	return bk, nil
}

func main() {
	//pass show func to the lambda handler
	lambda.Start(show)
}

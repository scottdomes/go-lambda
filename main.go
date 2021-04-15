package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type AdditionEvent struct {
	Number1 string `json:"number1"`
	Number2 string `json:"number2"`
}

func HandleRequest(ctx context.Context, numbers AdditionEvent) (string, error) {
	return fmt.Sprintf(numbers.Number1 + numbers.Number2), nil
}

func main() {
	lambda.Start(HandleRequest)
}

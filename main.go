package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type AdditionEvent struct {
	Number1 string `json:"number1"`
	Number2 string `json:"number2"`
}

func HandleRequest(ctx context.Context, numbers AdditionEvent) (string, error) {
	number1, err1 := strconv.Atoi(numbers.Number1)
	number2, err2 := strconv.Atoi(numbers.Number2)
	result := strconv.Itoa(number1 + number2)

	if err1 != nil || err2 != nil {
		result = "Error"
	}

	return fmt.Sprint(result), nil
}

func main() {
	lambda.Start(HandleRequest)
}

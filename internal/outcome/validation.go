package outcome

import (
	"context"
)

type ModelValidationFunction[T any] func(ctx context.Context, input T) *ValidationCheckError

type ModelValidationCheck[T any] struct {
	Name string
	F    ModelValidationFunction[T]
}

type BusinessValidationFunction[T any, R any] func(ctx context.Context, repository R, input T) *ValidationCheckError

type BusinessValidationCheck[T any, R any] struct {
	Name string
	F    BusinessValidationFunction[T, R]
}

type ValidationCheckError struct {
	Field string
	Error error
}

type ValidationCheckResultAggregate []*ValidationCheckError

func ValidateModel[T any](
	ctx context.Context,
	input T,
	testFunctions []ModelValidationCheck[T],
) ValidationCheckResultAggregate {
	checkResults := ValidationCheckResultAggregate{}

	for _, tf := range testFunctions {
		testResult := tf.F(ctx, input)

		if testResult != nil {
			checkResults = append(checkResults, testResult)
		}
	}

	return checkResults
}

func ValidateBusiness[T any, R any](
	ctx context.Context,
	repository R,
	input T,
	testFunctions []BusinessValidationCheck[T, R],
) ValidationCheckResultAggregate {
	checkResults := ValidationCheckResultAggregate{}

	for _, tf := range testFunctions {
		testResult := tf.F(ctx, repository, input)

		if testResult != nil {
			checkResults = append(checkResults, testResult)
		}
	}

	return checkResults
}

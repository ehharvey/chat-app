package abc

import (
	"context"

	"github.com/ehharvey/chat-app/internal/outcome"
)

type Service struct {
	fooRepository FooRepository
}

func (s Service) CreateFoo(ctx context.Context, input InsertOneFooParams) outcome.ServiceResult[FooModel] {
	// TODO

	// modelValidationFunctions := createFooModelCheckFunctions

	// businessValidationFunctions := createFooBusinessCheckFunctions

	// s.fooRepository.InsertOneFoo(ctx, input)

	return outcome.ServiceResult[FooModel]{}
}

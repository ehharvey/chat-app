package abc

import (
	"context"

	"github.com/ehharvey/chat-app/internal/outcome"
)

type Service struct {
	fooRepository Repository
	barRepository Repository
}

func (s Service) CreateFoo(ctx context.Context, input InsertOneFooParams) outcome.ServiceResult[FooModel] {
	return outcome.PerformService(
		ctx,
		input,
		s.fooRepository,
		createFooModelCheckFunctions[:],
		createFooBusinessCheckFunctions[:],
		s.fooRepository.InsertOneFoo,
	)
}

func (s Service) CreateBar(ctx context.Context, input InsertOneBarParams) outcome.ServiceResult[BarModel] {

	return outcome.ServiceResult[BarModel]{}
}

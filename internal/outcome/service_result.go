package outcome

import (
	"context"
	"fmt"
)

type ServiceResult[F any] struct {
	Result                    F
	Model                     string
	ValidationResultAggregate ValidationCheckResultAggregate
	InternalServerError       error
}

func PerformService[T any, R any, F any](
	ctx context.Context,
	input T,
	repository R,
	modelTestFunctions []ModelValidationCheck[T],
	businessTestFunctions []BusinessValidationCheck[T, R],
	persistenceFunction PersistenceFunction[T, F],
) ServiceResult[F] {
	result := ServiceResult[F]{
		Model: fmt.Sprintf("%T", input),
	}

	result.ValidationResultAggregate = ValidateModel(ctx, input, modelTestFunctions)

	if len(result.ValidationResultAggregate) > 0 {
		return result
	}

	result.ValidationResultAggregate = append(
		result.ValidationResultAggregate,
		ValidateBusiness(ctx, repository, input, businessTestFunctions)...,
	)

	if len(result.ValidationResultAggregate) > 0 {
		return result
	}

	persistenceResult, persistenceErr := persistenceFunction(ctx, input)

	result.Result = persistenceResult
	result.InternalServerError = persistenceErr
	return result
}

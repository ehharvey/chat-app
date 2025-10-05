package abc

import (
	"context"
	"fmt"

	"github.com/ehharvey/chat-app/internal/outcome"
)

var createBarBusinessCheckFunctions = [...]outcome.BusinessValidationCheck[InsertOneBarParams, Repository]{
	{
		Name: "checkValidBarNameLength",
		F:    checkBarNameUnique,
	},
}

func checkBarNameUnique(
	ctx context.Context,
	repository Repository,
	input InsertOneBarParams,
) *outcome.ValidationCheckError {
	check, err := repository.CheckExistsBarByName(ctx, CheckExistsBarByNameParams(input))

	if err != nil {
		// DB Error
		return &outcome.ValidationCheckError{
			Field: "",
			Error: fmt.Errorf("db error: %w", err),
		}
	} else if check.Exists {
		return &outcome.ValidationCheckError{
			Field: "Name",
			Error: fmt.Errorf("bar data already exists with name %s", input.Name),
		}
	} else {
		return nil
	}

}

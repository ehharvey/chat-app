package abc

import (
	"context"
	"fmt"

	"github.com/ehharvey/chat-app/internal/outcome"
)

var createFooBusinessCheckFunctions = [...]outcome.BusinessValidationCheck[InsertOneFooParams, FooRepository]{
	{
		Name: "checkValidFooNameLengh",
		F:    checkFooNameUnique,
	},
}

func checkFooNameUnique(
	ctx context.Context,
	repository FooRepository,
	input InsertOneFooParams,
) *outcome.ValidationCheckError {
	check, err := repository.CheckExistsFooByName(ctx, CheckExistsFooByNameParams(input))

	if err != nil {
		// Uh oh, probably a DB error
		return &outcome.ValidationCheckError{
			Field: "",
			Error: fmt.Errorf("db error: %w", err),
		}
	} else if check.Exists {
		return &outcome.ValidationCheckError{
			Field: "Name",
			Error: fmt.Errorf("foo data already exists with name %s", input.Name),
		}
	} else {
		return nil
	}
}

package abc

import (
	"context"
	"fmt"

	"github.com/ehharvey/chat-app/internal/outcome"
)

var createFooBusinessCheckFunctions = [...]outcome.BusinessValidationCheck[InsertOneFooParams, Repository]{
	{
		Name: "checkValidFooNameLengh",
		F:    checkFooNameUnique,
	},
}

func checkFooNameUnique(
	ctx context.Context,
	repository Repository,
	input InsertOneFooParams,
) *outcome.ValidationCheckError {
	check, err := repository.CheckExistsFooByName(ctx, CheckExistsFooByNameParams(input))

	if err != nil {
		// Uh oh, probably a DB error
		return &outcome.ValidationCheckError{
			Field: "",
			Error: fmt.Errorf("%w", outcome.ErrInternal),
		}
	} else if check.Exists {
		return &outcome.ValidationCheckError{
			Field: "Name",
			Error: fmt.Errorf("foo data %s invalid %w", input.Name, outcome.ErrDuplicate),
		}
	} else {
		return nil
	}
}

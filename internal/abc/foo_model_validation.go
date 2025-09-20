package abc

import (
	"context"
	"fmt"
	"unicode"

	"github.com/ehharvey/chat-app/internal/outcome"
)

var createFooModelCheckFunctions = [...]outcome.ModelValidationCheck[InsertOneFooParams]{
	{
		Name: "checkValidFooNameLengh",
		F:    checkValidFooNameLength,
	},
	{
		Name: "checkValidFooNameFormat",
		F:    checkValidFooNameFormat,
	},
}

func checkValidFooNameLength(
	ctx context.Context,
	createFooParams InsertOneFooParams,
) *outcome.ValidationCheckError {
	nameLength := len(createFooParams.Name)
	if nameLength > 0 {
		return nil
	} else {
		err := fmt.Errorf("length is %d", nameLength)
		return &outcome.ValidationCheckError{
			Field: "Name",
			Error: err,
		}
	}
}

// preconditions: checkValidFooNameLength Passed
func checkValidFooNameFormat(
	ctx context.Context,
	createFooParams InsertOneFooParams,
) *outcome.ValidationCheckError {
	firstCharacter := []rune(createFooParams.Name)[0]

	// Valid when first character is capitalized
	if unicode.IsUpper(firstCharacter) {
		return nil
	} else {
		return &outcome.ValidationCheckError{
			Field: "Name",
			Error: fmt.Errorf("first character %s is not capitalized", string(firstCharacter)),
		}
	}
}

package abc

import (
	"context"
	"fmt"
	"unicode"

	"github.com/ehharvey/chat-app/internal/outcome"
)

var createBarModelCheckFunctions = [...]outcome.ModelValidationCheck[InsertOneBarParams]{
	{
		Name: "checkValidBarNameLength",
		F:    checkValidBarNameLength,
	},
	{
		Name: "checkValidBarNameFormat",
		F:    checkValidBarNameFormat,
	},
}

func checkValidBarNameLength(
	ctx context.Context,
	createBarParams InsertOneBarParams,
) *outcome.ValidationCheckError {
	nameLength := len(createBarParams.Name)

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

// preconditions: checkValidBarNameLength Passed
func checkValidBarNameFormat(
	ctx context.Context,
	createBarParams InsertOneBarParams,
) *outcome.ValidationCheckError {
	firstCharacter := []rune(createBarParams.Name)[0]

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

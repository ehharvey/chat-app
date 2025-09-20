package outcome

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func stubValidateFunctionFail(ctx context.Context, input int) *ValidationCheckError {
	return &ValidationCheckError{
		Field: "Foo",
		Error: fmt.Errorf("New Error"),
	}
}

func stubValidateFunctionPass(ctx context.Context, input int) *ValidationCheckError {
	return nil
}

func TestValidateModelReportOnlyFailures(t *testing.T) {
	validateFunctions := []ModelValidationCheck[int]{
		{
			Name: "Fail",
			F:    stubValidateFunctionFail,
		},
		{
			Name: "Pass",
			F:    stubValidateFunctionPass,
		},
		{
			Name: "Fail",
			F:    stubValidateFunctionFail,
		},
	}

	expected := ValidationCheckResultAggregate{
		&ValidationCheckError{
			Field: "Foo",
			Error: fmt.Errorf("New Error"),
		},
		&ValidationCheckError{
			Field: "Foo",
			Error: fmt.Errorf("New Error"),
		},
	}

	actual := ValidateModel(
		context.Background(),
		5,
		validateFunctions,
	)

	comparer := func(x, y error) bool {
		return x.Error() == y.Error()
	}

	if !cmp.Equal(expected, actual, cmp.Comparer(comparer)) {
		t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(expected, actual, cmp.Comparer(comparer)))
	}
}

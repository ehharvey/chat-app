package abc

import (
	"context"
	"fmt"
	"testing"

	"github.com/ehharvey/chat-app/internal/outcome"
	"github.com/google/go-cmp/cmp"
)

func TestCheckValidFooNameLength(t *testing.T) {
	tests := []struct {
		name     string
		arrange  InsertOneFooParams
		expected *outcome.ValidationCheckError
	}{
		{
			name: "Length Zero",
			arrange: InsertOneFooParams{
				Name: "",
			},
			expected: &outcome.ValidationCheckError{
				Field: "Name",
				Error: fmt.Errorf("length is 0"),
			},
		},
		{
			name: "Length One",
			arrange: InsertOneFooParams{
				Name: "A",
			},
			expected: nil,
		},
	}

	comparer := func(x, y error) bool {
		return x.Error() == y.Error()
	}

	for _, tt := range tests {
		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			actual := checkValidFooNameLength(ctx, tt.arrange)
			if !cmp.Equal(tt.expected, actual, cmp.Comparer(comparer)) {
				t.Errorf("Mismatch (-want +got):\n%s", cmp.Diff(tt.expected, actual, cmp.Comparer(comparer)))
			}
		})
	}
}

// TODO TASK: Write tests for checkValidFooNameFormat

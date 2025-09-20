// Exists to check that repositories implement interfaces

package pg

import (
	"testing"

	"github.com/ehharvey/chat-app/internal/abc"
	"github.com/ehharvey/chat-app/internal/infrastructure/db/pg/generated"
)

func TestInterfaceImplementations(t *testing.T) {
	var _ abc.FooRepository = NewFooRepository(generated.Queries{})
}

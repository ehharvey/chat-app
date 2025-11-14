package outcome

import (
	"context"
)

type PersistenceFunction[T any, F any] func(
	ctx context.Context,
	input T,
) (F, error)

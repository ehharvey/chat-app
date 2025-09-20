package pg

import (
	"context"

	"github.com/ehharvey/chat-app/internal/abc"
	"github.com/ehharvey/chat-app/internal/infrastructure/db/pg/generated"
)

type FooRepository struct {
	queries generated.Queries
}

func NewFooRepository(queries generated.Queries) FooRepository {
	return FooRepository{
		queries: queries,
	}
}

func (r FooRepository) InsertOneFoo(
	ctx context.Context,
	input abc.InsertOneFooParams,
) (abc.FooModel, error) {
	result, err := r.queries.InsertOneFoo(ctx, input.Name)

	// TODO: Consider logging
	// TODO: Consider altering err instead of just bubbling up

	return abc.FooModel{
		ID:   result.ID.String(),
		Name: result.Name,
	}, err
}

func (r FooRepository) CheckExistsFooByName(
	ctx context.Context,
	input abc.CheckExistsFooByNameParams,
) (abc.CheckExistsFooByNameModel, error) {
	result, err := r.queries.CheckExistsFooByName(ctx, input.Name)

	return abc.CheckExistsFooByNameModel{
		Exists: result,
	}, err
}

package abc

import "context"

type Repository interface {
	InsertOneFoo(
		ctx context.Context,
		input InsertOneFooParams,
	) (FooModel, error) // Generated foo and error

	CheckExistsFooByName(
		ctx context.Context,
		input CheckExistsFooByNameParams,
	) (CheckExistsFooByNameModel, error)

	InsertOneBar(
		ctx context.Context,
		input InsertOneBarParams,
	) (BarModel, error)

	CheckExistsBarByName(
		ctx context.Context,
		input CheckExistsBarByNameParams,
	) (CheckExistsBarByNameModel, error)
}

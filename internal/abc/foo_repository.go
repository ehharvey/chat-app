package abc

import "context"

type FooRepository interface {
	InsertOneFoo(
		ctx context.Context,
		input InsertOneFooParams,
	) (FooModel, error) // Generated foo and error

	CheckExistsFooByName(
		ctx context.Context,
		input CheckExistsFooByNameParams,
	) (CheckExistsFooByNameModel, error)
}

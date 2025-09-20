package abc

// FooModel and FooModel-related *Model* types
type FooModel struct {
	ID   string
	Name string
}

type CheckExistsFooByNameModel struct {
	Exists bool
}

// Foo and Foo-related *Param* types
type InsertOneFooParams struct {
	Name string
}

type CheckExistsFooByNameParams struct {
	Name string
}

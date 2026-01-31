package abc

// BarModel and Barmodel-related *Model* types
type BarModel struct {
	ID   string
	Name string
}

type CheckExistsBarByNameModel struct {
	Exists bool
}

// Foo and Boo-related *Param* types
type InsertOneBarParams struct {
	Name string
}

type CheckExistsBarByNameParams struct {
	Name string
}

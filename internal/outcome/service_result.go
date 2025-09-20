package outcome

type ServiceResult[T any] struct {
	Result                    T
	Model                     string
	ValidationResultAggregate ValidationCheckResultAggregate
	PersistenceResult         DbResult
}

package facade

// SubsystemB provides functionality related to B.
type SubsystemB struct{}

func (s *SubsystemB) OperationB() string {
	return "SubsystemB: OperationB"
}

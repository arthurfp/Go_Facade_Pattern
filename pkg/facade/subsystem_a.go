package facade

// SubsystemA provides functionality related to A.
type SubsystemA struct{}

func (s *SubsystemA) OperationA() string {
	return "SubsystemA: OperationA"
}

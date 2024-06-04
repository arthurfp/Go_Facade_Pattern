package facade

// Facade provides a simplified interface to the complex subsystems.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

// NewFacade creates a new Facade with the given subsystems.
func NewFacade(a *SubsystemA, b *SubsystemB, c *SubsystemC) *Facade {
	return &Facade{
		subsystemA: a,
		subsystemB: b,
		subsystemC: c,
	}
}

// Operation performs a series of operations using the subsystems.
func (f *Facade) Operation() string {
	result := "Facade initializes subsystems:\n"
	result += f.subsystemA.OperationA() + "\n"
	result += f.subsystemB.OperationB() + "\n"
	result += f.subsystemC.OperationC() + "\n"
	result += "Facade coordinates subsystems to complete the operation."
	return result
}

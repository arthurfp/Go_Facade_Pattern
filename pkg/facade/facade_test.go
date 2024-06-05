package facade

import "testing"

func TestFacade_Operation(t *testing.T) {
	a := &SubsystemA{}
	b := &SubsystemB{}
	c := &SubsystemC{}
	facade := NewFacade(a, b, c)

	expected := "Facade initializes subsystems:\nSubsystemA: OperationA\nSubsystemB: OperationB\nSubsystemC: OperationC\nFacade coordinates subsystems to complete the operation."
	result := facade.Operation()

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

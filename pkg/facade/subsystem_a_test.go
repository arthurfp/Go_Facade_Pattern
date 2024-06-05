package facade

import "testing"

func TestSubsystemA_OperationA(t *testing.T) {
	subsystem := &SubsystemA{}
	expected := "SubsystemA: OperationA"
	result := subsystem.OperationA()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

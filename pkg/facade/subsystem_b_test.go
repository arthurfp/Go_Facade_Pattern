package facade

import "testing"

func TestSubsystemB_OperationB(t *testing.T) {
	subsystem := &SubsystemB{}
	expected := "SubsystemB: OperationB"
	result := subsystem.OperationB()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

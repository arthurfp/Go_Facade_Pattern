package facade

import "testing"

func TestSubsystemC_OperationC(t *testing.T) {
	subsystem := &SubsystemC{}
	expected := "SubsystemC: OperationC"
	result := subsystem.OperationC()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

package main

import (
	"facade-pattern-go/pkg/facade"
	"fmt"
)

func main() {
	a := facade.SubsystemA{}
	b := facade.SubsystemB{}
	c := facade.SubsystemC{}

	fmt.Println(a.OperationA())
	fmt.Println(b.OperationB())
	fmt.Println(c.OperationC())
}

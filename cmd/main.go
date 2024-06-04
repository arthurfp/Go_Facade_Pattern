package main

import (
	"facade-pattern-go/pkg/facade"
	"fmt"
)

func main() {
	a := &facade.SubsystemA{}
	b := &facade.SubsystemB{}
	c := &facade.SubsystemC{}

	f := facade.NewFacade(a, b, c)
	fmt.Println(f.Operation())
}

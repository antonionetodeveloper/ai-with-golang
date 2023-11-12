package main

import (
	"fmt"
	"gorgonia.org/gorgonia"
	"log"
)

func main() {
	graph := gorgonia.NewGraph()

	var x, y, z *gorgonia.Node
	var err error

	// define the expression
	x = gorgonia.NewScalar(graph, gorgonia.Float64, gorgonia.WithName("x"))
	y = gorgonia.NewScalar(graph, gorgonia.Float64, gorgonia.WithName("y"))
	if z, err = gorgonia.Add(x, y); 
	err != nil {
		log.Fatal(err)
	}

	// create a VM to run the program on
	machine := gorgonia.NewTapeMachine(graph)
	defer machine.Close()

	// set initial values
	gorgonia.Let(x, 3.0)
	gorgonia.Let(y, 2.5)

	// then run
	if err = machine.RunAll() 
	err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v \n", z.Value())
	fmt.Println(machine)
	fmt.Println(graph)

}

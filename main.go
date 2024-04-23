package main

import (
	"fmt"
	dfa "toc/automata"
)

func main() {
	d := dfa.NewDFA()
	d.SetStates([]dfa.State{0, 1, 2})
	d.SetInputAlphabet([]dfa.Symbol{0, 1, 2})
	d.SetFinalState([]dfa.State{1})

	d.AddTransitionFunctions(0, '0', []dfa.State{0})
	d.AddTransitionFunctions(0, '1', []dfa.State{1})
	d.AddTransitionFunctions(1, '0', []dfa.State{0})
	d.AddTransitionFunctions(1, '1', []dfa.State{2})
	d.AddTransitionFunctions(2, '0', []dfa.State{2})
	d.AddTransitionFunctions(2, '1', []dfa.State{1})
	d.AddTransitionFunctions(0, '2', []dfa.State{2, 1})
	// read later on from a file
	// sorting is pretty important

	istrue, err := d.CheckTheLanguage("1020")
	if istrue != true {
		fmt.Println(err)
	} else {
		fmt.Println("Accepted")
	}

	fmt.Println(d.GetFinalState())
}

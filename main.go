package main

import (
	"fmt"
	dfa "toc/automata"
)

// func main() {
// 	d := dfa.NewDFA()
// 	d.SetStates([]dfa.State{0, 1, 2})
// 	d.SetInputAlphabet([]dfa.Symbol{0, 1, 2})
// 	d.SetFinalState([]dfa.State{1})

// 	d.AddTransitionFunctions(0, '0', []dfa.State{0})
// 	d.AddTransitionFunctions(0, '1', []dfa.State{1})
// 	d.AddTransitionFunctions(1, '0', []dfa.State{0})
// 	d.AddTransitionFunctions(1, '1', []dfa.State{2})
// 	d.AddTransitionFunctions(2, '0', []dfa.State{2})
// 	d.AddTransitionFunctions(2, '1', []dfa.State{1})
// 	d.AddTransitionFunctions(0, '2', []dfa.State{2,1})
// 	d.AddGraph()
// 	// read later on from a file
// 	// sorting is pretty important

// 	istrue, err := d.CheckTheLanguage("1001")
// 	if istrue != true {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Accepted")
// 	}

// 	fmt.Println(d.GetFinalState())
// }

func main() {
	d := dfa.NewDFA()
	d.SetStates([]dfa.State{1, 2, 3})
	d.SetInputAlphabet([]dfa.Symbol{'0', '1'})
	d.SetFinalState([]dfa.State{1})

	d.AddTransitionFunctions(1, '0', []dfa.State{1, 2})
	d.AddTransitionFunctions(1, '1', []dfa.State{2})
	d.AddTransitionFunctions(2, '0', []dfa.State{3})
	d.AddTransitionFunctions(2, '1', []dfa.State{3})
	d.AddTransitionFunctions(3, '1', []dfa.State{3})
	// read later on from a file
	// sorting is pretty important

	// istrue, err := d.CheckTheLanguage("0")
	// if istrue != true {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Accepted")
	// }

	n := d.NfaToDfa()
	istrue, err := n.CheckTheLanguage("0")
	if istrue != true {
		fmt.Println(err)
	} else {
		fmt.Println("Accepted")
	}

	fmt.Println(d.GetTransitionFunctions())

	fmt.Println("finalallaallyyy", n.GetTransitionFunctions())
}

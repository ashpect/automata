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

	// initiate dfa
	d := dfa.NewDFA()

	// adding states, input alphabet, final states
	d.SetStates([]dfa.State{1, 2, 3})
	d.SetInputAlphabet([]dfa.Symbol{'0', '1'})
	d.SetFinalState([]dfa.State{1})

	// adding transition functions dfa/nfa
	d.AddTransitionFunctions(1, '0', []dfa.State{1, 2})
	d.AddTransitionFunctions(1, '1', []dfa.State{2})
	d.AddTransitionFunctions(2, '0', []dfa.State{3})
	d.AddTransitionFunctions(2, '1', []dfa.State{3})
	d.AddTransitionFunctions(3, '1', []dfa.State{3})

	fmt.Println("------------------------------------checking for string given in nfa------------------------------------")
	// checking for string 0 for dfa
	istrue, err := d.CheckTheLanguage("0")
	if istrue != true {
		fmt.Println(err)
	} else {
		fmt.Println("Accepted")
	}

	// printing for dfa
	fmt.Println("State\tSymbol\tNext State")
	for key, value := range d.GetTransitionFunctions() {
		fmt.Println(key.GetSourceState(), "\t", key.GetSourceSymbol()-48, "\t", value)
	}

	// creating nfa from dfa
	n := d.NfaToDfa()

	fmt.Println("------------------------------------")
	fmt.Println("New Dfa derived from Nfa")
	fmt.Println("States: ", n.GetState())
	fmt.Println("Input Alphabet: ", n.GetInputAlphabet())
	fmt.Println("Final State: ", n.GetFinalState())
	fmt.Println("Initial State: ", n.GetInitialState())

	// print final
	fmt.Println("------------------------------------")
	fmt.Println("State\tSymbol\tNext State")
	for key, value := range n.GetTransitionFunctions() {
		fmt.Println(key.GetSourceState(), "\t", key.GetSourceSymbol()-48, "\t", value)
	}

	fmt.Println("------------------------------------checking for string given in dfa------------------------------------")

	istrue, err = n.CheckTheLanguage("0")
	if istrue != true {
		fmt.Println(err)
	} else {
		fmt.Println("Accepted")
	}

}

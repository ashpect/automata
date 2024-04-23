package automata

import (
	"fmt"
	"sync"
)

type State int

type Symbol rune

type TransitionFunction map[StateInput][]State

type StateInput struct {
	sourceState State
	symbol      Symbol
}

type DFA struct {
	states              []State
	inputAlphabet       []Symbol
	finalState          []State
	initialState        State
	transitionFunctions []TransitionFunction
}

func NewDFA() *DFA {
	return &DFA{
		states:              make([]State, 0),
		inputAlphabet:       make([]Symbol, 0),
		finalState:          make([]State, 0),
		initialState:        0,
		transitionFunctions: make([]TransitionFunction, 0, 10),
	}
}

func (d *DFA) SetStates(state []State) {
	d.states = state
}

func (d *DFA) SetInputAlphabet(inputAlphabet []Symbol) {
	d.inputAlphabet = inputAlphabet
}

func (d *DFA) SetFinalState(finalState []State) {
	d.finalState = finalState
}

func (d *DFA) SetInitialState(initialState State) {
	d.initialState = initialState
}

func (d *DFA) AddTransitionFunctions(initital State, symbol Symbol, final []State) {
	transitionFunction := make(TransitionFunction)
	transitionFunction[StateInput{initital, symbol}] = final
	d.transitionFunctions = append(d.transitionFunctions, transitionFunction)
}

func (d DFA) GetState() []State {
	return d.states
}

func (d DFA) GetInputAlphabet() []Symbol {
	return d.inputAlphabet
}

func (d DFA) GetFinalState() []State {
	return d.finalState
}

func (d DFA) GetInitialState() State {
	return d.initialState
}

func (d DFA) GetTransitionFunctions() []TransitionFunction {
	return d.transitionFunctions
}

func (d DFA) CheckTheLanguage(input string) (bool, error) {
	istrue := false
	var wg = sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error)
	go func() { errChan <- d.CheckLanguage(input, &wg, &istrue) }()
	wg.Wait()
	return istrue, <-errChan
}

// Function is compatible to check for both dfa and nfa
func (d DFA) CheckLanguage(input string, wg *sync.WaitGroup, istrue *bool, states ...State) error {
	defer wg.Done()
	err := error(nil) // err nil imp, otherwise states will get redefined inside only
	if len(states) == 0 {
		states = append(states, d.initialState)
	}

	for k, symbol := range input {
		states, err = d.Advance(states[0], symbol)
		if err != nil {
			return err
		}
		if len(states) > 1 {
			for i := range states {
				var newwg = sync.WaitGroup{}
				newwg.Add(1)
				froCh := make(chan error)
				go func() { froCh <- d.CheckLanguage(input[k+1:], &newwg, istrue, states[i]) }()
				err = <-froCh
				newwg.Wait()
			}
		}
	}

	if d.IsFinalState(states) {
		*istrue = true
	} else {
		return fmt.Errorf("err1 : The language doesn't satisfy the finite automata, string read but not reached final state")
	}

	return err
}

func (d *DFA) Advance(currentstate State, input rune) ([]State, error) {
	for _, tf := range d.transitionFunctions {
		if finalstate, ok := tf[StateInput{currentstate, Symbol(input)}]; ok {
			return finalstate, nil
		}
	}
	return []State{}, fmt.Errorf("err3 : The language doesn't satisfy the finite automata")
}

func (d *DFA) IsFinalState(states []State) bool {

	// fast lookup time O(n+m) use map, although we use more space
	element := make(map[State]bool)
	for _, ok := range states {
		element[ok] = true
	}

	for _, a := range d.finalState {
		if element[a] == true {
			return true
		}
	}

	return false
}

func inttostate(arr []int) []State {
	states := make([]State, 0)
	for _, ok := range arr {
		states = append(states, State(ok))
	}
	return states
}

// func (n *DFA) NfaToDfa() *DFA {
// 	d := NewDFA()
// }

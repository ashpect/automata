package automata

import (
	"fmt"
	"strconv"
	"sync"
	g "toc/automata/helpers"
)

type State int

type Symbol rune

const phi State = -1

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
	transitionFunctions TransitionFunction
	graphForm           g.Graph
}

func NewDFA() *DFA {
	d := &DFA{
		states:              make([]State, 0),
		inputAlphabet:       make([]Symbol, 0),
		finalState:          make([]State, 0),
		initialState:        1,
		transitionFunctions: make(TransitionFunction),
		graphForm:           *g.NewGraph(),
	}
	d.AddGraph()
	return d
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
	d.transitionFunctions[StateInput{initital, symbol}] = final
	d.graphForm.AddEdge(int(initital), int(compressState(final)), rune(symbol))
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

func (d DFA) GetTransitionFunctions() TransitionFunction {
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

	if finalstate, ok := d.transitionFunctions[StateInput{currentstate, Symbol(input)}]; ok {
		return finalstate, nil
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

func compressState(states []State) State {
	state := ""
	for _, ok := range states {
		state += strconv.Itoa(int(ok))
	}
	result, _ := strconv.Atoi(state)
	return State(result)
}

func decompressState(state State) ([]State, error) {
	result := make([]State, 0)
	a := strconv.Itoa(int(state))
	for _, char := range a {
		charStr := string(char)

		num, err := strconv.Atoi(charStr)
		if err != nil {
			return []State{}, nil
		}
		result = append(result, State(num))
	}

	return result, nil
}

func (n *DFA) NfaToDfa() *DFA {
	d := NewDFA()
	d.SetInitialState(n.initialState)
	n.AddState(d, d.initialState)
	return d
}

func (n *DFA) AddState(d *DFA, state State) error {

	states, err := decompressState(state)
	if err != nil {
		return err
	}

	fmt.Println(n.GetTransitionFunctions())

	for _, input := range n.inputAlphabet {
		var final []State
		for _, state := range states {
			b, okk := n.transitionFunctions[StateInput{0, input}]
			if okk {
				fmt.Println(b)
			}
			a, ok := n.transitionFunctions[StateInput{state, input}]
			if ok {
				exists := false
				for _, check := range a {
					if len(final) > 0 {
						for _, state := range final {
							if state == check {
								exists = true
								break
							}
						}
					}
					if !exists {
						final = append(final, check)
					}
				}
			}
		}
		if len(final) == 0 || (compressState(final)) == state {
			continue
		}

		fmt.Println(state)
		d.AddTransitionFunctions(state, input, []State{compressState(final)})
		err := n.AddState(d, compressState(final))
		if err != nil {
			return err
		}

	}

	return nil
}

// Normal dfa to graph ; make it general and put in contructir
func (d *DFA) AddGraph() {
	gr := d.graphForm
	for _, ok := range d.states {
		gr.AddVertex(int(ok))
	}

	for si, st := range d.transitionFunctions {
		gr.AddEdge(int(si.sourceState), int(st[0]), rune(si.symbol))
	}

}

// for state {qi}, if there exists a walk from qi to qk, qk will be in the return slice of states.
func (d *DFA) getWalkFunction(state State, input Symbol, v *[]State) {

	*v = append(*v, state)
	a, ok := d.transitionFunctions[StateInput{state, input}]
	if !ok && (len(a) == 1 && a[0] == state) {
		return
	}
	for _, ok := range a {
		if !ifExists(*v, ok) {
			d.getWalkFunction(ok, input, v)
		}
	}

}

func ifExists(state []State, input State) bool {
	for _, ok := range state {
		if ok == input {
			return true
		}
	}
	return false
}

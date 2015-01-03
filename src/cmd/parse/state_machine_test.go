package parse

import (
	"testing"
)

type state struct {
	id string
}

func (s *state) isAcceptState() bool {
	return false
}

type acceptState struct {
	id string
}
func (s *acceptState) isAcceptState() bool {
	return true
}

func TestStateMachine1(t *testing.T) {
	first := &state{id: "1"}
	second := &acceptState{id: "2"}
	
	var firstToSecond Handler = func(context []rune) (State, []rune) {
		if context[0] == 'a' {
			return second, context
		} else {
			return nil, context
		}
	}

	machine := NewMachine(first)
	machine.AddState(first, firstToSecond)
	machine.AddState(second, nil)
	if machine.run([]rune("a")) != true {
		t.Errorf("got %v\nwant %v", "FAILURE", true)
	}
}

func TestStateMachine2(t *testing.T) {
	first := &state{id: "1"}
	second := &acceptState{id: "2"}

	var firstToSecond Handler = func(context []rune) (State, []rune) {
		if context[0] == 'a' {
			return second, context
		} else {
			return nil, context
		}
	}

	machine := NewMachine(first)
	machine.AddState(first, firstToSecond)
	machine.AddState(second, nil)
	if machine.run([]rune("b")) != false {
		t.Errorf("got %v\nwant %v", "FAILURE", false)
	}
}

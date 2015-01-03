package parse

type Machine struct {
	Handlers   map[State]func([]rune) (State, []rune)
	StartState State
}

func NewMachine(state State) *Machine {
	return &Machine{Handlers: make(map[State]func([]rune) (State, []rune)), StartState: state}
}

type State interface {
	isAcceptState() bool
}

type Handler func([]rune) (State, []rune)

func (m *Machine) AddState(state State, handler Handler) {
	m.Handlers[state] = handler
}

func (m *Machine) run(input []rune) bool {
	if handler, present := m.Handlers[m.StartState]; present {
		for {
			nextState, nextContext := handler(input)
			if nextState != nil {
				if nextState.isAcceptState() {
					return true
				} else {
					handler, present = m.Handlers[nextState]
					input = nextContext
				}
			} else {
				return false
			}
		}
	}
	return false
}

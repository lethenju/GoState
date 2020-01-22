package state_machine

type Moves interface {
	Get_transition_to(State) func()
	Get_reason_for(State) func()
	State_function() func()
}

type Connection struct {
	Connection_state  *State
	Reason_to_move     func() bool
	Transition         func()
}
type State struct {
	Name string
	Core_function func()
	Connected    []Connection
}

func (s State) Get_transition_to ( to *State) func() {
	for _, connected_state := range s.Connected {
		if (connected_state.Connection_state == to) {
			return connected_state.Transition
		}
	}
	// If we didnt find the transition, return empty function
	return func() {}
}


func (s State) Get_reason_for( to *State) func() bool {
	for _, connected_state := range s.Connected {
		if (connected_state.Connection_state == to) {
			return connected_state.Reason_to_move
		}
	}
	// If we didnt find the reason, return empty function
	return func() bool {
		return false
	}
}

func (s State) State_function () *State {
	var alive = true
	var next_state = &s
	s.Core_function()
	for alive {
		for _, connected_state := range s.Connected {
			if (connected_state.Reason_to_move() == true) { 
				connected_state.Transition()
				next_state = connected_state.Connection_state
				alive = false
				break
			}
		}
	}
	return next_state;
}

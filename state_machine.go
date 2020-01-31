/*  -- State Machine Framework 
 /  LICENCE : MIT
 /  Author : Julien LE THENO
*/
package state_machine

/** The interface Moves declares functions available for a State object
 */
type Moves interface {
	Get_transition_to(State) func()
	Get_reason_for(State) func()
	State_function() func()
}

/** Represents a Connection ( a transition) between states. It is strictly directional,
 *  if you want bidirectionnality you need 2 Connections.
 */ 
type Connection struct {
	Connection_state  *State       // The state it is connected to
	Reason_to_move     func() bool // The "reason to move" test function: code that tests if the transition is possible
	Transition         func()      // The transition function : code that is executed during a transition
}

/** Represents a State.
  */
type State struct {
	Name string               // Name of the State
	Core_function func()      // user-defined Function that gets executed when you enter that state. 
	Connected    []Connection // List of connections from that state
}

/** Implementation of the interface Moves
 *  return the transition function from a state to another one
 */
func (s State) Get_transition_to ( to *State) func() {
	for _, connected_state := range s.Connected {
		if (connected_state.Connection_state == to) {
			return connected_state.Transition
		}
	}
	// If we didnt find the transition, return empty function
	return func() {}
}


/** Implementation of the interface Moves
 *  return the "reason_to_move" function from a state to another one
 */
func (s State) Get_reason_for( to *State) func() bool {
	for _, connected_state := range s.Connected {
		if (connected_state.Connection_state == to) {
			return connected_state.Reason_to_move
		}
	}
	// If we didnt find the reason, return empty function
	// TODO Error handling instead of this !!!
	return func() bool {
		return false
	}
}


/** State function : it is the underlying function that makes it all work.
 *  It manages the transitions, given the results of reason to move. 
 *  Returns the next state. 
 */
func (s State) State_function () *State {
	var next_state = &s
	for _, connected_state := range s.Connected {
		if (connected_state.Reason_to_move() == true) { 
			connected_state.Transition()
			next_state = connected_state.Connection_state
			break
		}
	}

	return next_state;
}

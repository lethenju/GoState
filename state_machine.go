/*  -- State Machine Framework
/  LICENCE : MIT
/  Author : Julien LE THENO
*/

package statemachine

/*Moves The interface Moves declares functions available for a State object
 */
type Moves interface {
	Get_transition_to(State) func()
	Get_reason_for(State) func()
	State_function() func()
}

/*Connection Represents a Connection ( a transition) between states. It is strictly directional,
 *  if you want bidirectionnality you need 2 Connections.
 */
type Connection struct {
	ConnectionState *State      // The state it is connected to
	ReasonToMove    func() bool // The "reason to move" test function: code that tests if the transition is possible
	Transition      func()      // The transition function : code that is executed during a transition
}

/*State Represents a State.
 */
type State struct {
	Name         string       // Name of the State
	CoreFunction func()       // user-defined Function that gets executed when you enter that state.
	Connected    []Connection // List of connections from that state
}

/*GetTransitionTo Implementation of the interface Moves
 *  return the transition function from a state to another one
 */
func (s State) GetTransitionTo(to *State) func() {
	for _, connectedState := range s.Connected {
		if connectedState.ConnectionState == to {
			return connectedState.Transition
		}
	}
	// If we didnt find the transition, return empty function
	return func() {}
}

/*GetReasonFor Implementation of the interface Moves
 *  return the "reason_to_move" function from a state to another one
 */
func (s State) GetReasonFor(to *State) func() bool {
	for _, connectedState := range s.Connected {
		if connectedState.ConnectionState == to {
			return connectedState.ReasonToMove
		}
	}
	// If we didnt find the reason, return empty function
	// TODO Error handling instead of this !!!
	return func() bool {
		return false
	}
}

/*StateFunction State function : it is the underlying function that makes it all work.
 *  It manages the transitions, given the results of reason to move.
 *  Returns the next state.
 */
func (s State) StateFunction() *State {
	var nextState = &s
	for _, connectedState := range s.Connected {
		if connectedState.ReasonToMove() == true {
			connectedState.Transition()
			nextState = connectedState.ConnectionState
			break
		}
	}

	return nextState
}

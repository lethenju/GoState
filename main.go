package main

import (
    "fmt"
    "time"
)

var counter = 0;

type moves interface {
	get_transition_to(state) func()
	get_reason_for(state) func()
	state_function() func()
}

type connection struct {
	connection_state  *state
	reason_to_move     func() bool
	transition         func()
}
type state struct {
	name string
	core_function func()
	connected    []connection
}

func (s state) get_transition_to ( to *state) func() {
	for _, connected_state := range s.connected {
		if (connected_state.connection_state == to) {
			return connected_state.transition
		}
	}
	// If we didnt find the transition, return empty function
	return func() {}
}


func (s state) get_reason_for( to *state) func() bool {
	for _, connected_state := range s.connected {
		if (connected_state.connection_state == to) {
			return connected_state.reason_to_move
		}
	}
	// If we didnt find the reason, return empty function
	return func() bool {
		return false
	}
}

func (s state) state_function () *state {
	var alive = true
	var next_state = &s
	s.core_function()
	for alive {
		for _, connected_state := range s.connected {
			if (connected_state.reason_to_move() == true) { 
				connected_state.transition()
				next_state = connected_state.connection_state
				alive = false
				break
			}
		}
	}
	return next_state;
}

func runtime(s *state) {
	var alive = true 
	for alive {
		fmt.Println("Counter = ", counter)
		time.Sleep(500 * time.Millisecond)
		s = (*s).state_function()
	}
}


func main() {
	fmt.Println("Welcome to state machine")

	var normal_state state
	normal_state.name = "Normal"
	normal_state.core_function = func() {
		fmt.Println(".")
		counter++
	}
	
	var fizz_state  state
	fizz_state.name = "Fizz"
	fizz_state.core_function = func() {
		fmt.Println(" == FIZZ == ")
	}

	var buzz_state  state
	buzz_state.name = "Buzz"
	buzz_state.core_function = func() {
		fmt.Println(" == BUZZ == ")
	}

	// Describing connection from normal state to fizz state
	normal_state.connected   = append(normal_state.connected, 
		connection{ connection_state : &fizz_state,
			reason_to_move : func () bool { 
					if (counter % 3 == 0) {
						return true
					}
					return false 
				},
				transition : func () {
					fmt.Println("[Normal] -> [Fizz]")
				}},
	// Describing connection from normal state to buzz state
		connection{ connection_state : &buzz_state,
			reason_to_move : func () bool { 
					if (counter % 5 == 0) {
						return true
					}
					return false 
				},
				transition : func () {
					fmt.Println("[Normal] -> [Buzz]")
				}},
	// Describing connection from normal state to normal state
		connection{ connection_state : &normal_state,
			reason_to_move : func () bool { return true },
			    transition : func () {
					fmt.Println("[Normal] -> [Normal]")
				}});

	// Describing connection from fizz to buzz
	fizz_state.connected = append(fizz_state.connected,
		connection{ connection_state : &buzz_state,
			reason_to_move : func () bool { 
				if (counter % 5 == 0) {
					return true
				}
				return false 
			},
			transition : func () {
				fmt.Println("[Fizz] -> [Buzz]")
			}},
	// Describing connection from fizz to normal
		connection{ connection_state : &normal_state,
			reason_to_move : func () bool { 
				return true
			},
			transition : func () {
				fmt.Println("[Fizz] -> [Normal]")
			}});


	// Buzz to normal
	buzz_state.connected = append(buzz_state.connected,
		connection{ connection_state : &normal_state,
			reason_to_move : func () bool { 
				return true
			},
			transition : func () {
				fmt.Println("[Buzz] -> [Normal]")
			}});


	runtime(&normal_state)
}

package main

import (
    "fmt"
	"time"
	sm "go/state_machine"
)

var counter = 0;

func runtime(s *sm.State) {
	var alive = true 
	for alive {
		s.Core_function()
		fmt.Println("Counter = ", counter)
		time.Sleep(500 * time.Millisecond)

		s = (*s).State_function()
	}
}


func main() {
	fmt.Println("Welcome to state machine")

	var normal_state sm.State
	normal_state.Name = "Normal"
	normal_state.Core_function = func() {
		fmt.Println(".")
		counter++
	}
	
	var fizz_state  sm.State
	fizz_state.Name = "Fizz"
	fizz_state.Core_function = func() {
		fmt.Println(" == FIZZ == ")
	}

	var buzz_state  sm.State
	buzz_state.Name = "Buzz"
	buzz_state.Core_function = func() {
		fmt.Println(" == BUZZ == ")
	}

	// Describing connection from normal state to fizz state
	normal_state.Connected   = append(normal_state.Connected, 
		sm.Connection{ Connection_state : &fizz_state,
			Reason_to_move : func () bool { 
					if (counter % 3 == 0) {
						return true
					}
					return false 
				},
				Transition : func () {
					fmt.Println("[Normal] -> [Fizz]")
				}},
	// Describing connection from normal state to buzz state
		sm.Connection{ Connection_state : &buzz_state,
			Reason_to_move : func () bool { 
					if (counter % 5 == 0) {
						return true
					}
					return false 
				},
				Transition : func () {
					fmt.Println("[Normal] -> [Buzz]")
				}},
	// Describing connection from normal state to normal state
		sm.Connection{ Connection_state : &normal_state,
			Reason_to_move : func () bool { return true },
			    Transition : func () {
					fmt.Println("[Normal] -> [Normal]")
				}});

	// Describing connection from fizz to buzz
	fizz_state.Connected = append(fizz_state.Connected,
		sm.Connection{ Connection_state : &buzz_state,
			Reason_to_move : func () bool { 
				if (counter % 5 == 0) {
					return true
				}
				return false 
			},
			Transition : func () {
				fmt.Println("[Fizz] -> [Buzz]")
			}},
	// Describing connection from fizz to normal
		sm.Connection{ Connection_state : &normal_state,
			Reason_to_move : func () bool { 
				return true
			},
			Transition : func () {
				fmt.Println("[Fizz] -> [Normal]")
			}});


	// Buzz to normal
	buzz_state.Connected = append(buzz_state.Connected,
		sm.Connection{ Connection_state : &normal_state,
			Reason_to_move : func () bool { 
				return true
			},
			Transition : func () {
				fmt.Println("[Buzz] -> [Normal]")
			}});


	runtime(&normal_state)
}
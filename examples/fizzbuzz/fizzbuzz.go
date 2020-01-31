/** 
 *  This is a simple example on how to use the state machine framework in your project.
 *  This example declares the state machine statically in the code.
 *  Do you know Fizzbuzz ? It is a simple game where you count numbers and you need to say "fizz"
 *  for each multiple of 3, and buzz for each multiple of 5.
 *  If it is both a multiple of 3 and 5, you should say fizzbuzz.
 *
 *  We will modelize the game as a state machine described as so :
 * 
 * (Fizz state) --> (Buzz state)
 *         ^          ^
 *         v          V
 *         (Normal state)
 *
 * Core functions :
 * The normal state increments the counter
 * The fizz state display FIZZ
 * The buzz state display BUZZ
 *
 * Reasons to move, and transition :
 * We go from normal state to Fizz state if the number is a multiple of 3, and we display the transition
 * We go from normal state to Buzz state if the number is a multiple of 5, and we display the transition
 * We stay in normal state otherwise, and we display the non-transition
 *
 * We go from Fizz state to Buzz state if the number is also a multiple of 5, and we display the transition
 * We go from Fizz state to normal directly, to continue the game, and we display the transition
 * We go from Buzz state to normal directly, to continue the game, and we display the transition
 */
package main

import (
    "fmt"
    "time"
    sm "go/state_machine"
)

// Our counter 
var counter = 0;

/** The runtime function is the heart of the State Machine. You need this minimal code to make the framework run
 *
 *  
 func runtime(s *sm.State) { // The parameter is the entry state of the SM.
	for {
	        // Launch the user-defined core function of the state 
		s.Core_function()
		// Transition to another state if possible (the state function will take care of everything)
		s = (*s).State_function()
	}
}
 *
 */
func runtime(s *sm.State) {
	for {
		s.Core_function()
		fmt.Println("Counter = ", counter)
		time.Sleep(500 * time.Millisecond)

		s = (*s).State_function()
	}
}

/** Our main function
 */
func main() {
	fmt.Println("Welcome to Fizz-Buzz Demo!")
        
	// Declaring normal state
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

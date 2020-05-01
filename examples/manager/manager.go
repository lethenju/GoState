package main

import (
	"fmt"
	"time"

	sm "github.com/lethenju/state_machine_framework"
)

func runtime(s *sm.State) {
	for {
		s.CoreFunction()

		// TODO What to do in the transitions ?
		time.Sleep(500 * time.Millisecond)

		s = (*s).StateFunction()
	}
}


func main() {
	fmt.Println("Welcome to Manager!")
	var running_flag = true;
	while (running_flag) {
		/** TODO Get user command
		 -> Create variable (name, value)

		 -> Create state (Name)
		 -> Add variable change in state
				(State Name, Entering/Leaving, VAR, OPERATION, VALUE or VAR)			 
						Variable change include :
							=  	(VAR) (INCR) (VALUE or VAR) // incrementation
							=  	(VAR) (=) (VALUE or VAR) // affectation
							
		 -> Create Transition (Transition Name, From state, To state)
		 -> Add Reason to move
		 		(Transition Name, VAR, COMPARISON, VALUE or VAR)
					    A reason to move is based on variables
						   = when (VAR) > (VALUE or VAR) 
						   = when (VAR) == (VALUE or VAR)
		 -> Display states
					- State Name 
							In : (VAR) (OPERATION) (VAR OR VALUE)
							      ...
							Out : (VAR) (OPERATION) (VAR OR VALUE)
							      ...
						Transitions In :
							Name transitions
							..
						Transition Out : 
							Name transitions
							..
		 -> Display Transitions
					- Transition Name
							In : State
							Reason : (VAR) (CONDITION) (VAR)
									 ...
							Out : State

		-> Run (manually)
			-> Step
			-> Display Current State
			-> Display Current Transition
			-> Display vars
					Display all vars with values
		-> Run (auto) with timer (sleep between transitions)
		*/
	}
}
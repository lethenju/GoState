package main

import (
	"fmt"
	"time"
	"strings"
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


type stateMachine struct {
	states []State;
}


func main() {
	fmt.Println("Welcome to Manager!")
	var runningFlag = true
	for (runningFlag) {

		var userInput string
		fmt.Print("> ")
		fmt.Scan(&userInput)
		userInput = strings.ToUpper(userInput)
		command := strings.Split(userInput, " ")
		switch (command[0]) {
		case "CREATE":
			switch (command[1]) {
				
		 //-> Create variable (name, value)
			case "VARIABLE":
				//TODO Add variable command[2] with optional value command[3]
				break
				
		 //-> Create state (Name)
			case "STATE":
				//TODO Add state command[2]
				break
		 //-> Create Transition (Transition Name, From state, To state)
			case "TRANSITION":
				// Verify command[3] and command[4] are in the list of states
				// Todo add transition command[2] to the state command[3]
				// command[4]
				break
			}
		case "ADD":
			switch (command[1]) {
			
			/*-> Add variable change in state
				(State Name, Entering/Leaving, VAR, OPERATION, VALUE or VAR)			 
					Variable change include :
						=  	(VAR) (INCR) (VALUE or VAR) // incrementation
						=  	(VAR) (=)    (VALUE or VAR) // affectation */
			case: "STATE_BEHAVIOUR":
				// TODO Verify command in command[3] command[4] command[5]
				// command[3] must be in variables
				// command[4] must be in commandlist
				// command[5] must be in variable OR is possible to cast as int
				
				switch (command[2]) {
				case "ENTERING":
					// TODO Add command in core function of state
					break;
				case "LEAVING":
					// TODO Add command in ?
					break;
				}
				break;
			
		    /*-> Add Reason to move
		 		(State from, State to, VAR, COMPARISON, VALUE or VAR)
					    A reason to move is based on variables
						   = when (VAR) > (VALUE or VAR) 
						   = when (VAR) == (VALUE or VAR)*/
			case: "REASON_TO_MOVE":
				// TODO verify in all states if the transition exists
				// if not create it
			}
		}

		/** TODO Get user command


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
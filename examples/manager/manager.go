package main

import (
	"errors"
	"fmt"
	"strings"
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

// A used defined variable
// Can be a string or an int, but the value is always stored as a string
type Variable struct {
	Name  string
	Value string
}

// The state machine representation
type stateMachine struct {
	States    []sm.State
	Variables []Variable
}

func commandHandler(command []string, machine *stateMachine) (err error) {

	switch command[0] {
	case "CREATE":
		switch command[1] {
		//-> Create variable (name, value)
		case "VARIABLE":
			if len(command) < 3 || len(command) > 4 {
				return errors.New("Number of parameters incorrect")
			}
			if len(command) < 4 {
				machine.Variables = append(machine.Variables, Variable{
					Name:  command[2],
					Value: ""})
			}
			machine.Variables = append(machine.Variables, Variable{
				Name:  command[2],
				Value: command[3]})
			break

		//-> Create state (Name)
		case "STATE":
			if len(command) != 3 {
				return errors.New("Number of parameters incorrect")
			}

			machine.States = append(machine.States, sm.State{
				Name:         command[3],
				CoreFunction: func() {}})
			break
		//-> Create Transition ( From state, To state)
		case "TRANSITION":
			// Verify command[2] and command[3] are in the list of states
			// Then add the transition

			if len(command) != 4 {
				return errors.New("Number of parameters incorrect")
			}
			// Search for the state
			var from *sm.State
			var to *sm.State

			for _, state := range machine.States {
				if command[2] == state.Name {
					from = &state
				}
				if command[3] == state.Name {
					to = &state
				}
			}
			if from == nil {
				return errors.New("Didnt find state " + command[2])
			} else if to == nil {
				return errors.New("Didnt find state " + command[3])

			}
			// Todo add transition  to the state command[2]
			from.Connected = append(from.Connected, sm.Connection{ConnectionState: to,
				ReasonToMove: func() bool {
					return true
				},
				Transition: func() {
					fmt.Println("[" + from.Name + "] -> [" + to.Name + "]")
				}})
			break
		}
	case "ADD":
		switch command[1] {

		/*-> Add variable change in state
		(State Name, Entering/Leaving, VAR, OPERATION, VALUE or VAR)
			Variable change include :
				=  	(VAR) (INCR) (VALUE or VAR) // incrementation
				=  	(VAR) (=)    (VALUE or VAR) // affectation */
		case "STATE_BEHAVIOUR":
			// TODO Verify command in command[3] command[4] command[5]
			// command[3] must be in variables
			// command[4] must be in commandlist
			// command[5] must be in variable OR is possible to cast as int

			switch command[2] {
			case "ENTERING":
				// TODO Add command in core function of state
				break
			case "LEAVING":
				// TODO Add command in ?
				break
			}
			break

			/*-> Add Reason to move
			 (State from, State to, VAR, COMPARISON, VALUE or VAR)
					A reason to move is based on variables
					   = when (VAR) > (VALUE or VAR)
					   = when (VAR) == (VALUE or VAR)*/
		case "REASON_TO_MOVE":
			// TODO verify in all states if the transition exists
			// if not create it
			break
		}
	/*
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
				..*/
	case "DISPLAY":
		// Todo display state
		break
	case "RUN":
		switch command[1] {
		case "MANUAL":
			break
		case "AUTO":
			break
		}
	}

	/** TODO Get user command

	-> Run (manually)
		-> Step
		-> Display Current State
		-> Display Current Transition
		-> Display vars
				Display all vars with values
	-> Run (auto) with timer (sleep between transitions)
	*/
	return nil
}

func main() {
	fmt.Println("Welcome to Manager!")

	var machine stateMachine

	var runningFlag = true
	for runningFlag {

		var userInput string
		fmt.Print("> ")
		fmt.Scan(&userInput)
		userInput = strings.ToUpper(userInput)
		command := strings.Split(userInput, " ")
		commandHandler(command, &machine)
	}
}

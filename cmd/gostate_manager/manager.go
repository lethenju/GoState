package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	sm "github.com/lethenju/gostate/pkg"
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

func getState(stateStr string, states *[]sm.State) *sm.State {
	// We cannot do a for range as we're looking for a reference.
	// The range operation is copying objects and not references
	for i := 0; i < len(*states); i++ {
		if stateStr == (*states)[i].Name {
			return &(*states)[i]
		}
	}
	return nil
}

func getVariable(variableStr string, variables []Variable) *Variable {
	for _, variable := range variables {
		if variableStr == variable.Name {
			return &variable
		}
	}
	return nil
}

func commandHandler(command []string, machine *stateMachine) (err error) {

	switch strings.ToUpper(command[0]) {
	case "CREATE":
		if len(command) < 2 {
			return errors.New("Not enough arguments")
		}
		switch strings.ToUpper(command[1]) {
		//-> Create variable (name, value)
		case "VARIABLE":
			if len(command) < 3 || len(command) > 4 {
				return errors.New("Number of parameters incorrect")
			}
			if len(command) < 4 {
				machine.Variables = append(machine.Variables, Variable{
					Name:  command[2],
					Value: ""})
			} else {
				machine.Variables = append(machine.Variables, Variable{
					Name:  command[2],
					Value: command[3]})
			}
			break

		//-> Create state (Name)
		case "STATE":
			if len(command) != 3 {
				return errors.New("Number of parameters incorrect")
			}

			machine.States = append(machine.States, sm.State{
				Name:         command[2],
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

			from := getState(command[2], &machine.States)
			to := getState(command[3], &machine.States)

			if from == nil {
				return errors.New("Didnt find state " + command[2])
			} else if to == nil {
				return errors.New("Didnt find state " + command[3])

			}
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
		switch strings.ToUpper(command[1]) {

		/*-> Add variable change in state
		(State Name, VAR, OPERATION, VALUE or VAR)
			Variable change include :
				=  	(VAR) (INCR) (VALUE or VAR) // incrementation
				=  	(VAR) (=)    (VALUE or VAR) // affectation */
		case "STATE_BEHAVIOUR":
			// TODO Verify command in command [2] command[3] command[4] command[5]
			// command[2] must be in states
			// command[3] must be in variables
			// command[4] must be in commandlist
			// command[5] must be in variable OR is possible to cast as int
			if len(command) != 6 {
				return errors.New("Number of parameters incorrect")
			}
			state := getState(command[2], &machine.States)
			if state == nil {
				return errors.New("State " + command[2] + " is unknown")
			}
			firstVar := getVariable(command[3], machine.Variables)
			if firstVar == nil {
				return errors.New("Variable " + command[3] + " is unknown")
			}
			if strings.ToUpper(command[4]) != "INCR" && command[4] != "=" {
				return errors.New("Operation '" + command[4] + "' unknown. Use '=' or 'INCR'")
			}
			operation := strings.ToUpper(command[4])
			secondVar := getVariable(command[5], machine.Variables)
			var value string
			if secondVar == nil {
				// Verify it is castable
				_, err := strconv.Atoi(command[5])
				if err != nil {
					return errors.New("Second variable '" + command[5] + "' is not a known variable nor a int value")
				}
				_, err = strconv.Atoi(firstVar.Value)
				if err != nil {
					return errors.New("First variable '" + firstVar.Name + "' is not a int variable (Value='" + firstVar.Value + "')")
				}

				value = command[5]
			} else {
				value = secondVar.Value
			}
			switch operation {
			case "=":
				state.CoreFunction = func() {
					firstVar.Value = value
				}
				break
			case "INCR":

				state.CoreFunction = func() {
					// Append if string, increment if int
					secondVarInt, err := strconv.Atoi(value)
					if err != nil {
						// Append
						firstVar.Value = firstVar.Value + secondVar.Value
					} else {
						// Increment
						firstVarInt, err := strconv.Atoi(firstVar.Value)
						if err != nil {
							fmt.Print("First variable '" + firstVar.Name + "' is not a int variable (Value='" + firstVar.Value + "')")
							panic(0)
						}
						firstVarInt += secondVarInt
						firstVar.Value = strconv.Itoa(firstVarInt)
					}
				}

			}

			break

			/*-> Add Reason to move
			 (State from, State to, VAR, COMPARISON, VALUE or VAR)
					A reason to move is based on variables
					   = when (VAR) > (VALUE or VAR)
					   = when (VAR) == (VALUE or VAR)*/
		case "REASON_TO_MOVE":
			// TODO verify state from exists,
			// TODO verify State to exists,
			// TODO verify VAR1 exists
			// TODO verify comparison is in ">", "==", "<", ">=", "<="
			//    And only '==' if VAR1 and VAR2 are strings
			// TODO verify VAR2 exists or is castable
			// TODO verify State From if the transition exists
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
		states := machine.States
		for _, state := range states {
			fmt.Println("State : " + state.Name)
			for _, transition := range state.Connected {
				fmt.Println(" -> " + transition.ConnectionState.Name)
			}
		}

		for _, variable := range machine.Variables {
			fmt.Println("Variable : " + variable.Name + " (" + variable.Value + ")")
		}
		break
	case "RUN":
		switch strings.ToUpper(command[1]) {
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

		scanner := bufio.NewScanner(os.Stdin)
		var command []string
		for len(command) == 0 {

			fmt.Print("> ")
			if !scanner.Scan() {
				return
			}
			userInput := scanner.Text()
			command = strings.Fields(userInput)

		}
		if strings.ToUpper(command[0]) == "EXIT" {
			fmt.Println("Goodbye !")
			runningFlag = false
			break
		}
		res := commandHandler(command, &machine)
		if res != nil {
			fmt.Println("ERROR :" + res.Error())
		}
	}
}

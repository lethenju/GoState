/*  -- State Machine Framework 
 /  LICENCE : MIT
 /  Author : Julien LE THENO
*/
package state_machine

import (
	"io"
        "fmt"
	"os"
	"encoding/csv" 
)

/** List of states
 *  TODO DONT LET IT STATIC ! 
 */
var list_of_states []*State;

/** Parses the "state" csv file and the "transition" csv file to construct the model in state machine go objects and structures.
 *  Takes in parameter maps with the callback functions to introduce in the model, and their names as a String.
 *  Maps are separated between "map_function" for void functions
 *  and "map_reasons" that have to return a bool
 *  Returns the first state of the SM.
 */
func Parse_and_install( map_functions map[string] func (),  map_reasons map[string] func () bool) *State {
	
	// open the State file 
	// TODO files in argument 
	input_file_states, err := os.Open("states")
	if (err != nil) {
		fmt.Println(err)
		return nil
	}
	// instantiate a Reader from the CSV package.
	lines := csv.NewReader(input_file_states)
	
	// Parsing states
	for {
		line, err := lines.Read()
		if (err == io.EOF) {
			break
		}
		list_of_states = append(list_of_states, &State {
			Name : line[0],
			Core_function : map_functions[line[1]],
		})
	}
	
	// open the transition file
	input_file_transitions, err := os.Open("transitions")
		if (err != nil) {
		fmt.Println(err)
		return nil
	}
        // instantiate a Reader from the CSV package.
	lines_transitions := csv.NewReader(input_file_transitions)
	// Parse the transitions
	for {
		line, err := lines_transitions.Read()
		if (err == io.EOF) {
			break
		}
		// TODO Handle cases when the referenced state doesnt exist
		for _, state_from := range list_of_states{
			if state_from.Name == line[0] {
				for _, state_to := range list_of_states{
					if state_to.Name == line[1] { 
						state_from.Connected = 
						append(state_from.Connected,
						       Connection{ Connection_state : state_to,
				                                   Reason_to_move: map_reasons[line[2]],
							           Transition: map_functions[line[3]],})
						break;
					}
				}
				break;		
			}
		}
	}

	return list_of_states[0]	
}

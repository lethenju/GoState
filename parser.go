package state_machine

import (
	"io"
    "fmt"
	"os"
	"encoding/csv" 
)


var list_of_states []*State;


func Parse_and_install( map_functions map[string] func (),  map_reasons map[string] func () bool) *State {
	
	input_file_states, err := os.Open("states")
	if (err != nil) {
		fmt.Println(err)
		return nil
	}
	
	lines := csv.NewReader(input_file_states);
	
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

	input_file_transitions, err := os.Open("transitions")
	
	lines_transitions := csv.NewReader(input_file_transitions);
	
	for {
		line, err := lines_transitions.Read()
		if (err == io.EOF) {
			break
		}
		for _, state_from := range list_of_states{
			if state_from.Name == line[0] {
				for _, state_to := range list_of_states{
					if state_to.Name == line[1] { 
						state_from.Connected = append(state_from.Connected,
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
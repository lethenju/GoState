/*  -- State Machine Framework
/  LICENCE : MIT
/  Author : Julien LE THENO
*/

package statemachine

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

/** List of states
 *  TODO DONT LET IT STATIC !
 */
var listOfStates []*State

/*ParseAndInstall Parses the "state" csv file and the "transition" csv file to construct the model in state machine go objects and structures.
*  Takes in parameter maps with the callback functions to introduce in the model, and their names as a String.
*  Maps are separated between "map_function" for void functions
*  and "map_reasons" that have to return a bool
*  Returns the first state of the SM, otherwise an error.
 */
func ParseAndInstall(stateFile string, transitionFile string, mapFunctions map[string]func(), mapReasons map[string]func() bool) (*State, error) {

	// open the State file
	// TODO files in argument
	inputFileStates, err := os.Open(stateFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// instantiate a Reader from the CSV package.
	lines := csv.NewReader(inputFileStates)

	// Parsing states
	for {
		line, err := lines.Read()
		if err == io.EOF {
			break
		}
		listOfStates = append(listOfStates, &State{
			Name:         line[0],
			CoreFunction: mapFunctions[line[1]],
		})
	}

	// open the transition file
	inputFileTransitions, err := os.Open(transitionFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// instantiate a Reader from the CSV package.
	linesTransitions := csv.NewReader(inputFileTransitions)
	// Parse the transitions
	for {
		line, err := linesTransitions.Read()
		if err == io.EOF {
			break
		}
		// TODO Handle cases when the referenced state doesnt exist
		for _, stateFrom := range listOfStates {
			if stateFrom.Name == line[0] {
				for _, stateTo := range listOfStates {
					if stateTo.Name == line[1] {
						stateFrom.Connected =
							append(stateFrom.Connected,
								Connection{ConnectionState: stateTo,
									ReasonToMove: mapReasons[line[2]],
									Transition:   mapFunctions[line[3]]})
						break
					}
				}
				break
			}
		}
	}

	return listOfStates[0], nil
}

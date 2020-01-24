package main
import (
	"bufio"
	"io"
    "fmt"
	"os"
	"encoding/csv" 
	"strings"
	sm "go/state_machine"
)

var position = "Chambre"
var reader = bufio.NewReader(os.Stdin)
		
func runtime(s *sm.State) {
	var alive = true 
	var old_position string
	for alive {
		s.Core_function()
		fmt.Println("Déplacements possibles : ")
		for _, possible_move := range s.Connected {
			fmt.Println(" - ", possible_move.Connection_state.Name)
		}
		fmt.Print("Vous voulez aller a ")
		old_position = position
		position, _ = reader.ReadString('\n')
		position = strings.TrimRight(position, "\b")
		position = strings.TrimSpace(position)
		
		s2 := (*s).State_function()
		if (s.Name == s2.Name) {
			fmt.Println("Nope vous pouvez pas")
			position = old_position 
		}
		s = s2
	}
}


func r_chambre_salon() bool { return (position == "Salon") } 
func r_salon_chambre() bool { return (position == "Chambre") } 


func t_chambre_salon() {
	fmt.Println(" Vous ouvrez la porte vers le salon ")
} 
func t_salon_chambre() {
	fmt.Println(" Vous ouvrez la porte vers la chambre ")
}

func core_chambre() {
	fmt.Println("Vous voila dans votre chambre ! Ah que c'est beau !")
}
func core_salon() {
	fmt.Println("Vous voila dans le salon, il est grand et personne n'est la..")
}
var list_of_states []*sm.State;


func parse_and_install() {
	map_functions := map[string] func ()  {
		"core_chambre": core_chambre,
		"core_salon"  : core_salon,
		"t_chambre_salon": t_chambre_salon,
		"t_salon_chambre": t_salon_chambre,
	}
	map_reasons := map[string] func () bool {
		"r_chambre_salon": r_chambre_salon,
		"r_salon_chambre": r_salon_chambre,
	}

	input_file_states, err := os.Open("states")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	
	lines := csv.NewReader(input_file_states);
	
	for {
		line, err := lines.Read()
		if (err == io.EOF) {
			break
		}
		list_of_states = append(list_of_states, &sm.State {
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
							sm.Connection{ Connection_state : state_to,
										   Reason_to_move: map_reasons[line[2]],
										   Transition: map_functions[line[3]],})

						break;
					}
				}
				break;		
			}
		}

	}
	
}
func main() {
	parse_and_install()
	runtime(list_of_states[0])
}

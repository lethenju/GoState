package main
import (
	"bufio"
	"os"
    "fmt"
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


func r_to_salon() bool { return (position == "Salon") } 
func r_to_chambre() bool { return (position == "Chambre") } 
func r_to_cuisine() bool { return (position == "Cuisine") } 

func t_to_salon() {
	fmt.Println(" Vous ouvrez la porte vers le salon ")
} 
func t_to_chambre() {
	fmt.Println(" Vous ouvrez la porte vers la chambre ")
}
func t_to_cuisine() {
	fmt.Println(" Vous ouvrez la porte vers la cuisine ")
}

func core_chambre() {
	fmt.Println("Vous voila dans votre chambre ! Ah que c'est beau !")
}
func core_salon() {
	fmt.Println("Vous voila dans le salon, il est grand et personne n'est la..")
}
func core_cuisine() {
	fmt.Println("La cuisine est sale..")
}

func main() {
	map_functions := map[string] func ()  {
		"core_chambre": core_chambre,
		"core_salon"  : core_salon,
		"core_cuisine" : core_cuisine,
		"t_to_salon": t_to_salon,
		"t_to_chambre": t_to_chambre,
		"t_to_cuisine": t_to_cuisine,
		
	}
	map_reasons := map[string] func () bool {
		"r_to_salon": r_to_salon,
		"r_to_chambre": r_to_chambre,
		"r_to_cuisine": r_to_cuisine,
	}
	
	first_state := sm.Parse_and_install(map_functions, map_reasons)
	
	runtime(first_state)
}
/** 
 *  This is a simple example on how to use the state machine framework in your project.
 *  This example parses a csv-described state machine and launches it
 *  You have a little house with a kitchen (Cuisine in French),
 *  a living room (Salon in French), and a bedroom (Chambre in French). You can move from a room to another.
 *
 *  We will modelize the "game" as a state machine described as so :
 * 
 * (Cuisine) <--> (Salon) <--> (Chambre)
 *
 * Core functions :
 *   in each state we just display info on the room we are in 
 *
 * Reasons to move, and transition :
 * We go from Cuisine to Salon if the user types "Salon", and we display the transition
 * We go from Salon to Chambre if the user types "Chambre", and we display the transition
 * We go from Salon to Cuisine if the user types "Cuisine", and we display the transition
 * We go from Chambre to Salon if the user types "Salon", and we display the transition
 *
 */
package main
import (
	"bufio"
	"os"
        "fmt"
	"strings"
	sm "go/state_machine"
)

// We start from the "Chambre" position.
var position = "Chambre"

var reader = bufio.NewReader(os.Stdin)


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
	var old_position string
	for {
		s.Core_function()
		// Listing possible moves
		fmt.Println("Déplacements possibles : ")
		for _, possible_move := range s.Connected {
			fmt.Println(" - ", possible_move.Connection_state.Name)
		}
		fmt.Print("Vous voulez aller a ")
		old_position = position // saving actual position
		// asking where to go to the user.
		position, _ = reader.ReadString('\n')
		position = strings.TrimRight(position, "\b")
		position = strings.TrimSpace(position)
		
		s2 := (*s).State_function() // Getting new state
		
		if (s.Name == s2.Name) {
			// Means we didnt move
			fmt.Println("Nope vous pouvez pas")
			position = old_position // We overwrite what the user said..
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
	// Declaring our void callbacks
	map_functions := map[string] func ()  {
		"core_chambre": core_chambre,
		"core_salon"  : core_salon,
		"core_cuisine" : core_cuisine,
		"t_to_salon": t_to_salon,
		"t_to_chambre": t_to_chambre,
		"t_to_cuisine": t_to_cuisine,
		
	}
	// Declaring our boolean-returning callbacks
	map_reasons := map[string] func () bool {
		"r_to_salon": r_to_salon,
		"r_to_chambre": r_to_chambre,
		"r_to_cuisine": r_to_cuisine,
	}
	
	// Launch the parsing, install the model and get a reference to the first state
	first_state := sm.Parse_and_install(map_functions, map_reasons)
	
	// Launch the SM with the first state.
	runtime(first_state)
}

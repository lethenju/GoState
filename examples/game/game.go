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
	"fmt"
	"os"
	"strings"

	sm "github.com/lethenju/state_machine_framework"
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
	var oldPosition string
	for {
		s.CoreFunction()
		// Listing possible moves
		fmt.Println("Déplacements possibles : ")
		for _, possibleMove := range s.Connected {
			fmt.Println(" - ", possibleMove.ConnectionState.Name)
		}
		fmt.Print("Vous voulez aller a ")
		oldPosition = position // saving actual position
		// asking where to go to the user.
		position, _ = reader.ReadString('\n')
		position = strings.TrimRight(position, "\b")
		position = strings.TrimSpace(position)

		s2 := (*s).StateFunction() // Getting new state

		if s.Name == s2.Name {
			// Means we didnt move
			fmt.Println("Nope vous pouvez pas")
			position = oldPosition // We overwrite what the user said..
		}
		s = s2
	}
}

func rToSalon() bool   { return (position == "Salon") }
func rToChambre() bool { return (position == "Chambre") }
func rToCuisine() bool { return (position == "Cuisine") }

func tToSalon() {
	fmt.Println(" Vous ouvrez la porte vers le salon ")
}
func tToChambre() {
	fmt.Println(" Vous ouvrez la porte vers la chambre ")
}
func tToCuisine() {
	fmt.Println(" Vous ouvrez la porte vers la cuisine ")
}

func coreChambre() {
	fmt.Println("Vous voila dans votre chambre ! Ah que c'est beau !")
}
func coreSalon() {
	fmt.Println("Vous voila dans le salon, il est grand et personne n'est la..")
}
func coreCuisine() {
	fmt.Println("La cuisine est sale..")
}

func main() {
	// Declaring our void callbacks
	mapFunctions := map[string]func(){
		"core_chambre": coreChambre,
		"core_salon":   coreSalon,
		"core_cuisine": coreCuisine,
		"t_to_salon":   tToSalon,
		"t_to_chambre": tToChambre,
		"t_to_cuisine": tToCuisine,
	}
	// Declaring our boolean-returning callbacks
	mapReasons := map[string]func() bool{
		"r_to_salon":   rToSalon,
		"r_to_chambre": rToChambre,
		"r_to_cuisine": rToCuisine,
	}

	// Launch the parsing, install the model and get a reference to the first state
	firstState := sm.ParseAndInstall("states", "transitions", mapFunctions, mapReasons)

	// Launch the SM with the first state.
	runtime(firstState)
}

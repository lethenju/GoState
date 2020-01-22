package main

// Game 
// State diagram
// Jardin  <--> Station de bus <--> Ecole
//    ^                        <--> Magasin
//    |
//    v
//  Cuisine <--> Salon <--> Chambre

import (
    "bufio"
    "fmt"
	"os"
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

func main() {

	var chambre sm.State
	chambre.Name = "Chambre"
	chambre.Core_function = func() {
		fmt.Println("Vous voila dans votre chambre ! Ah que c'est beau !")
	}
	var salon sm.State
	salon.Name = "Salon"
	salon.Core_function = func() {
		fmt.Println("Vous voila dans le salon, il est grand et personne n'est la..")
	}
	var cuisine sm.State
	cuisine.Name = "Cuisine"
	cuisine.Core_function = func() {
		fmt.Println("Vous êtes dans la cuisine.. Ca sent mauvais")
	}
	var jardin sm.State
	jardin.Name = "Jardin"
	jardin.Core_function = func() {
		fmt.Println("Vous voila dans le jardin ! Ah que c'est beau !")
	}
	var station sm.State
	station.Name = "Station de bus"
	station.Core_function = func() {
		fmt.Println("Ya personne a la Station de bus")
	}
	var ecole sm.State
	ecole.Name = "Ecole"
	ecole.Core_function = func() {
		fmt.Println("Vous êtes a l'école.. Il faudrait peut etre travailler")
	}
	var magasin sm.State
	magasin.Name = "Magasin"
	magasin.Core_function = func() {
		fmt.Println("Vous êtes au magasin, envie de qqchose ?")
	}


	// Describing connection from chambre to salon
	chambre.Connected   = append(chambre.Connected, 
		sm.Connection{ Connection_state : &salon,
			Reason_to_move : func () bool { 
					return (position == salon.Name)
				},
				Transition : func () {
					fmt.Println(" Vous ouvrez la porte vers le salon ")
				}});

	// Describing connection from salon to chambre
	salon.Connected   = append(salon.Connected, 
		sm.Connection{ Connection_state : &chambre,
			Reason_to_move : func () bool { 
					return (position == chambre.Name)
				},
				Transition : func () {
					fmt.Println(" Vous ouvrez la porte vers la chambre ")
				}}, 
		sm.Connection{ Connection_state : &cuisine,
			Reason_to_move : func () bool { 
					return (position == cuisine.Name)
				},
				Transition : func () {
					fmt.Println(" Vous ouvrez la porte vers la cuisine ")
				}});
	
	cuisine.Connected = append(cuisine.Connected,
		sm.Connection{ Connection_state : &salon,
			Reason_to_move : func () bool {
				return (position == salon.Name)
			},
			Transition : func () {
				fmt.Println(" Vous ouvrez la porte vers le salon ")
			}},
		sm.Connection{ Connection_state : &jardin,
				Reason_to_move : func () bool {
					return (position == jardin.Name)
				},
				Transition : func () {
					fmt.Println(" Vous ouvrez la porte vers le jardin ")
				}});
	
	jardin.Connected = append(jardin.Connected,
		sm.Connection{ Connection_state : &cuisine,
			Reason_to_move : func () bool {
				return (position == cuisine.Name)
			},
			Transition : func () {
				fmt.Println(" Vous ouvrez la porte vers la cuisine ")
			}},
		sm.Connection{ Connection_state : &station,
			Reason_to_move : func () bool {
				return (position == station.Name)
			},
			Transition : func () {
				fmt.Println(" Vous marchez vers la station ")
			}});


	station.Connected = append(station.Connected,
		sm.Connection{ Connection_state : &ecole,
			Reason_to_move : func () bool {
				return (position == ecole.Name)
			},
			Transition : func () {
				fmt.Println("Le bus pour l'école arrive ! ")
			}},
		sm.Connection{ Connection_state : &magasin,
			Reason_to_move : func () bool {
				return (position == magasin.Name)
			},
			Transition : func () {
				fmt.Println("Le bus pour le magasin arrive")
			}},
		sm.Connection{ Connection_state : &jardin,
			Reason_to_move : func () bool {
				return (position == jardin.Name)
			},
			Transition : func () {
				fmt.Println("Le bus pour rentrer chez soi arrive")
			}});

		ecole.Connected   = append(ecole.Connected, 
		sm.Connection{ Connection_state : &station,
			Reason_to_move : func () bool { 
					return (position == station.Name)
				},
				Transition : func () {
					fmt.Println("Vous retournez a la station")
				}});

		magasin.Connected   = append(magasin.Connected, 
		sm.Connection{ Connection_state : &station,
			Reason_to_move : func () bool { 
					return (position == station.Name)
				},
				Transition : func () {
					fmt.Println("Vous retournez a la station")
				}});

	runtime(&chambre)
}
package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

type players struct {
	list []*player
}


type player struct {
	id int
	life float64
	name string
	weapons map[string]float64
	coins float64
	playerType string
}

var playerId int = 0

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getOption(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)

	opt, err := r.ReadString('\n')

	opt = strings.TrimSpace(opt)

	switch opt {
	case "A":
		break
	case "B":
		break
	case "C":
		break
	case "D":
		break
	default:
		getOption("Incorrect Option please chose a valid one\n" + prompt, r)	
	}
	return opt, err

}

func printSinglePlayer(p *player) {
	if (p == nil) {
		return 
	}
	fmt.Println("Player Name: ", p.name)
	fmt.Println("Player id: ", p.id)
	fmt.Println("Player Life: ", p.life)
	fmt.Println("Player coins: ", p.coins)
	fmt.Println("Player Weapons: ")
	for weapon, damage := range p.weapons {
		fmt.Printf("    Weapon: %s ----- cost: %.2f\n", weapon, damage)
	}
}


func newPlayer(playerType string) player {
	reader := bufio.NewReader(os.Stdin)
	playerName, _ := getInput("Enter player name: ", reader)
	newPlayer := player{
		id: playerId,
		life: 100.0,
		name: playerName,
		weapons: map[string]float64{},
		coins: 100.0,
		playerType: playerType,
	}
	playerId++
	return newPlayer
}

func createNewPlayer(ps *players) player {
	reader := bufio.NewReader(os.Stdin)
	playerType, _ := getOption(`
Chose player type:

A) Atacking player (their attacks are more strong due to the speed however they have less life because their skin is to thin)
B) Defensive player (their attack is not so strong but can handle attacks due to their thick skin)\n
C) Strong player (their attack is strong and can handle attacks very well due that they have a strong sking)
D) Archer (they can attack from distance and their attacks are strong but they cannot handle attacks due to the thin skin)

 Option:`, reader)
	newPlayer := newPlayer(playerType)
	ps.addPlayer(&newPlayer)
	return newPlayer
}

func (ps *players) addPlayer(p *player) {
	if ps == nil {
		return }
	ps.list = append(ps.list, p)
}
							
func (ps *players) printPlayers() {
	if (ps == nil){
		fmt.Println("You got not players get one!")}
	for _, p := range ps.list {
		fmt.Println("-> Player id ", p.id)
		fmt.Println(" -- player Name: ", p.name)
		fmt.Println(" -- player Weapons: ", p.weapons)
		fmt.Println(" -- player Life: ", p.life)
		fmt.Println(" -- player coins: ", p.coins)
	}
}

func (ps *players) getSinglePlayer(id int) *player {
	for _, p := range ps.list {
		if p.id == id {
			return p
		}
	}
	return nil
}

func showWeaponsToBuy() map[string]float64{

	fmt.Println("Chose the weapons to buy by typing their name:")
	weapons := map[string]float64{"gun": 2.0, "shotgun": 4.5, "sub mchine": 3.0, "rifle": 6.5}
	for k , v := range weapons {
		fmt.Printf("name: %v    cost: %v\n", k, v)
	}

	return weapons
}

func selectWeapon(p *player, input string, weapons map[string]float64) (string, float64) {
	for k, v := range weapons {
		if k == input {
			return k, v
		}
	}
	return "", 0.0
}

func checkWeapons(p *player, r *bufio.Reader) {
	if p == nil {
		return
	}
	weapons := showWeaponsToBuy()
	input, _ := getInput("Enter the weapon name to buy (to not buy anything write CLOSE)\n ->", r)
	if (input == "CLOSE") {
		return
	}
	key, value := selectWeapon(p, input, weapons)
	if (key == "") {
		fmt.Printf("The input: '%v' is incorrect\n", input)
		checkWeapons(p, r)
		return
	}
	p.weapons[key] = value
	p.coins -= value
	printSinglePlayer(p)
	checkWeapons(p, r)
}

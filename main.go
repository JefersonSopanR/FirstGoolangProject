package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func usePlayer(ps *players, r *bufio.Reader) {
	ps.printPlayers()
	inputStr, _ := getInput("chose player by typing their Id: ", r)
	input , err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("Invalid Id\nYou must type a valid Id!")
		return
	}
	player := ps.getSinglePlayer(input)
	if player == nil {
		fmt.Printf("Player wiht the id %v, does not exist!\n", input)
		return 
	}
	printSinglePlayer(player)
	fmt.Printf("You can update the '%v' player\n", player.name)
	option, _ := getOption(" A) add Weapons\n B) Increment life \n C) change name\n D) delete Player\n -> ", r)
	switch option {
	case "A":
		checkWeapons(player, r)
	case "B":
		//incrementLife(player, r)
	case "C":
		//changePlayerName(player, r)
	case "D":
		//deletePlayer(ps, player)
	default:
		return
	}
}

func main() {
	players := players{
		list: []*player{},
	}

	input := ""
	reader := bufio.NewReader(os.Stdin)

	for input != "E" {
		fmt.Println("You can chose this options:")
		input, _ = getInput(`C) create a new Player P) print players U) usePlayer E) EXIT
`, reader)
		switch	input {
		case "C":
				player := createNewPlayer(&players)
				fmt.Println(player)
		case "P":
				players.printPlayers()
		case "U":
			usePlayer(&players,  reader)
		case "E":
				break
			default:
				fmt.Println("Option Not valid chose a valid one!")
		}
	}
}

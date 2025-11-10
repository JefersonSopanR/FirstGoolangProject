package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


type player struct {
	level int
	exp float64
	name string
	weapons map[string]float64
	life float64
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err = r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createPlayer() player {
	reader = bufio.NewReader(os.stdin)

	name, _= getInput("Enetr the player's name", reader)

	pl := player{
		level: 0,
		exp: 0.0,
		name: name,
		weapons: map[string]float64{},
		life: 100.0
	}
	return pl
}







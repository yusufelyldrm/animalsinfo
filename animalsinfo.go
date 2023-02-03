package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal is a struct that contains the food, locomotion and noise of an animal
type Animal struct{ food, locomotion, noise string }

// Eat prints the food of an animal
func (a Animal) Eat() { fmt.Println(a.food) }

// Move prints the locomotion of an animal
func (a Animal) Move() { fmt.Println(a.locomotion) }

// Speak prints the noise of an animal
func (a Animal) Speak() { fmt.Println(a.noise) }

// animals contains informations about animals
var animals = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter animal name and action (e.g. cow eat)")
		fmt.Print("> ")
		if scanner.Scan() {
			line := scanner.Text()
			tokens := strings.Split(line, " ")
			if len(tokens) != 2 {
				fmt.Println("Invalid request")
				continue
			}
			animal, info := tokens[0], tokens[1]
			a, ok := animals[animal]
			if !ok {
				fmt.Println("Invalid animal name")
				continue
			}
			switch info {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Println("Invalid request")
			}
		} else {
			break
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal is an interface that contains the Eat, Move and Speak methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Animal is a struct that contains the food, locomotion and noise of an animal
type Animals struct{ food, locomotion, noise string }

func (a Animals) Eat()   { fmt.Println(a.food) }
func (a Animals) Move()  { fmt.Println(a.locomotion) }
func (a Animals) Speak() { fmt.Println(a.noise) }

// animals contains informations about animals
var animals = map[string]Animals{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

// addAnimal adds an animal to the map
func addAnimal(name, food, locomotion, noise string) {
	_, exists := animals[name]

	if exists {
		fmt.Println("Animal already exists")
		return
	}
	animals[name] = Animals{food, locomotion, noise}
	fmt.Println("Animal added successfully!")
	askingWhatToDo()
}

// asking asks the user what to do
func askingWhatToDo() {
	fmt.Println("What do you want to do? Add animal or query animal? (e.g. add or query)")
}

func main() {
	askingWhatToDo()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("> ")
		line := scanner.Text()
		if line == "add" {
			fmt.Println("Enter animal name, food, locomotion and noise (e.g. cow grass walk moo)")
			fmt.Print("> ")
			if scanner.Scan() {
				line := scanner.Text()
				tokens := strings.Split(line, " ")
				if len(tokens) != 4 {
					fmt.Println("Invalid request")
					askingWhatToDo()
					continue
				}
				addAnimal(tokens[0], tokens[1], tokens[2], tokens[3])
			}
		} else if line == "query" {

			fmt.Println("Enter animal name and action (e.g. cow eat)")
			fmt.Print("> ")
			if scanner.Scan() {
				line := scanner.Text()
				tokens := strings.Split(line, " ")
				if len(tokens) != 2 {
					fmt.Println("Invalid request")
					askingWhatToDo()
					continue
				}
				animal, info := tokens[0], tokens[1]
				anim, ok := animals[animal]
				if !ok {
					fmt.Println("Invalid animal name")
					continue
				}
				switch info {
				case "eat":
					anim.Eat()
					askingWhatToDo()
				case "move":
					anim.Move()
					askingWhatToDo()
				case "speak":
					anim.Speak()
					askingWhatToDo()
				default:
					fmt.Println("Invalid request")
					askingWhatToDo()
				}
			}

		} else {
			fmt.Println("Invalid request")
			askingWhatToDo()
		}
	}
}

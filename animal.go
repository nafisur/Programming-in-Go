package main

import "fmt"
import "os"

type Animal struct {
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (x *Animal) Eat() {
	fmt.Println(x.foodEaten)
}

func (x *Animal) Speak() {
	fmt.Println(x.spokenSound)
}

func (x *Animal) Move() {
	fmt.Println(x.locomotionMethod)
}

func main() {
	var animal, command string
	for {
		fmt.Printf("> ")
		if _, err := fmt.Scanf("%s %s", &animal, &command); err != nil {
			os.Exit(0)
		}

		var x Animal
		switch animal {
		case "cow":
			x = Animal{foodEaten: "grass", locomotionMethod: "walk", spokenSound: "moo"}
		case "bird":
			x = Animal{foodEaten: "worms", locomotionMethod: "fly", spokenSound: "peep"}
		case "snake":
			x = Animal{foodEaten: "mice", locomotionMethod: "slither", spokenSound: "hsss"}
		}

		switch command {
		case "eat":
			x.Eat()
		case "move":
			x.Move()
		case "speak":
			x.Speak()
		}
	}
}

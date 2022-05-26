package main

import (
	"fmt"

	"github.com/mrrizal/go-design-pattern/strategy_pattern/duck"
)

func main() {
	wildDuck := duck.DuckStruct{
		Flying:     duck.JetFly{},
		Quacking:   duck.SuperQuack{},
		Displaying: duck.Display{},
	}
	duckName := "Mio"

	fmt.Println(wildDuck.Display(duckName))
	fmt.Println(wildDuck.Quack(duckName))
	fmt.Println(wildDuck.Fly())

	cityDuck := duck.DuckStruct{
		Flying: duck.SimpleFly{},
		Quacking: duck.SimpleQuack{},
		Displaying: duck.Display{},
	}
	duckName = "Alex"

	fmt.Println(cityDuck.Display(duckName))
	fmt.Println(cityDuck.Quack(duckName))
	fmt.Println(cityDuck.Fly())
}

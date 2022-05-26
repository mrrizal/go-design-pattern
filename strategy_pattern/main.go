package main

import (
	"fmt"

	"github.com/mrrizal/go-design-pattern/strategy_pattern/duck"
)

func main() {
	wildDuck := duck.DuckStruct{
		duck.JetFly{},
		duck.SuperQuack{},
		duck.Display{},
	}

	duckName := "Mio"
	fmt.Println(wildDuck.Fly())
	fmt.Println(wildDuck.Quack(duckName))
	fmt.Println(wildDuck.Display(duckName))
}

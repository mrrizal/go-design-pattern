package duck

import "fmt"

// define interfaces
type FlyInterface interface {
	Fly() string
}

type QuackInterface interface {
	Quack(name string) string
}

type DisplayInterface interface {
	Display(name string) string
}

// define struct/object
type DuckStruct struct {
	Flying     FlyInterface
	Quacking   QuackInterface
	Displaying DisplayInterface
}

// define method/func of struct/object
func (d *DuckStruct) Fly() string {
	return d.Flying.Fly()
}

func (d *DuckStruct) Quack(name string) string {
	return d.Quacking.Quack(name)
}

func (d *DuckStruct) Display(name string) string {
	return d.Displaying.Display(name)
}

// implementation

type SimpleFly struct{}

func (SimpleFly) Fly() string {
	return "Simple fly!"
}

type JetFly struct{}

func (JetFly) Fly() string {
	return "Jet fly!"
}

type SimpleQuack struct{}

func (SimpleQuack) Quack(name string) string {
	return fmt.Sprintf("%s: Quack!", name)
}

type SuperQuack struct{}

func (SuperQuack) Quack(name string) string {
	return fmt.Sprintf("%s: Quack!Quack!Quack!Quack!Quack!", name)
}

type Display struct{}

func (Display) Display(name string) string {
	return fmt.Sprintf("i'm %s", name)
}

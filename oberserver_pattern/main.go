package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		eventTime time.Time
		eventName string
	}

	ObserverInterface interface {
		Display(Event)
	}

	ObserverableInterface interface {
		Add(ObserverStruct)
		Remove(ObserverStruct)
		Notify(Event)
	}

	DisplayInterface interface {
		Display(int, Event)
	}
)

type (
	ObserverStruct struct {
		id int
	}

	ObserverableStruct struct {
		observers map[ObserverStruct]struct{}
	}
)

func (o *ObserverStruct) Display(event Event) {
	fmt.Printf("%s: observer %d: receive %s\n", event.eventTime.Format(time.RFC3339), o.id, event.eventName)
}

func (o *ObserverableStruct) Add(ob ObserverStruct) {
	o.observers[ob] = struct{}{}
}

func (o *ObserverableStruct) Remove(ob ObserverStruct) {
	delete(o.observers, ob)
}

func (o *ObserverableStruct) Notify(event Event) {
	for observer := range o.observers {
		observer.Display(event)
	}
}

func main() {
	n := ObserverableStruct{
		observers: map[ObserverStruct]struct{}{},
	}

	n.Add(ObserverStruct{id: 1})
	n.Add(ObserverStruct{id: 2})
	n.Add(ObserverStruct{id: 3})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	counter := 1
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{eventName: fmt.Sprintf("event %d", counter), eventTime: t})
			counter++
		}
	}
}

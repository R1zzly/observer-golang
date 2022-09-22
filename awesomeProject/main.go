package main

import (
	"awesomeProject/observer"
)

func main() {
	hhkz := observer.NewJobSite()
	bob := observer.NewPerson("Bob")
	hhkz.AddVacancies("SuperMan")
	hhkz.Subscribe(bob)
	hhkz.AddVacancies("Batman")
	hhkz.RemoveVacancies("SuperMan")

	hhkz.Unsubscribe(bob)
	aibek := observer.NewPerson("Aibek")
	hhkz.Subscribe(aibek)
	hhkz.AddVacancies("Wonder-woman")

}

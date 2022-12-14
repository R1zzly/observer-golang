package observer

import "fmt"

type Person struct {
	name string
}

func (P *Person) HandleEvent(vacancies []string) {
	fmt.Println("Hi dear " + P.name)
	fmt.Println("vacancies updated: ")
	for i := range vacancies {
		fmt.Println(vacancies[i])
	}
}

func NewPerson(name string) *Person {
	person := new(Person)
	person.name = name
	return person
}

package main

import "fmt"


func mainLearn() {

	//* Learn pointer & and * */
	person := Person{Name: "Aisyah", Age: 24}
	updateName(person, "Yasmin")
	updateAge(&person, 26)
	fmt.Println(person.Name, person.Age)
}

type Person struct {
	Name string
	Age int
}

func updateName(p Person, newName string){
	p.Name = newName
}

func updateAge(p *Person, newAge int){
	p.Age = newAge
}
package main

import (
	"fmt"
)

type Address struct {
	Country string
	City    string
}
type Person struct {
	Name string
	Age  int
	Address
}

func (p Person) introduce() {
	fmt.Printf("Hello, my name is %s and I'm %d years old. I'm from %s, %s.\n", p.Name, p.Age, p.Country, p.City)
}
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

type Counter struct {
	count int
}

func (c *Counter) increment() {
	c.count++
}

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func newSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}

func changeArray(a *[3]int) {
    a[0] = 20
}

func main() {
	Person1 := Person{
		Name: "Abemelek",
		Age:  22,
		Address: Address{
			Country: "Ethiopia",
			City:    "Addis Ababa",
		},
	}
	Person1.introduce()
	updateAge(&Person1, 21)
	Person1.introduce()

	counter := Counter{}
	fmt.Println("intial counter", counter.count)

	counter.increment()
	fmt.Println("and then ", counter.count)

	goku := newSaiyan("Abemelek", 2000)
	fmt.Printf("my name is %s and i have power of %d", goku.Name, goku.Power)
	fmt.Println()

	gohan := &Saiyan{
		Name:  "Abemelek",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Daniel",
			Power:  2001,
			Father: nil,
		},
	}
	fmt.Printf("my name is %s and i have power of %d , my father is %s and his power is %d", gohan.Name, gohan.Power, gohan.Father.Name, gohan.Father.Power)
	fmt.Println()
	scores := [4]int{43,25,89,22}
	for i:=0; i < len(scores); i++{
		fmt.Println(scores[i])
	}
	for index ,value := range scores{
		fmt.Printf("index = %d its value %d ",index,value)
	}
	fmt.Println()

	arr3 := [3]int{1,2,3}
	changeArray(&arr3)
	fmt.Println(arr3)
}

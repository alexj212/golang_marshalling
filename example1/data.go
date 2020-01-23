package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Email     string
	City      string
	Country   string
	Age       int32
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s - email: %s city: %s country: %s age: %d", p.FirstName, p.LastName, p.Email, p.City, p.Country, p.Age)
}

type People struct {
	Members []*Person
}

func (f *People) String() string {
	return fmt.Sprintf("members: %d", len(f.Members))
}

var Alex = &Person{
	FirstName: "Alex",
	LastName:  "Jeannopoulos",
	Email:     "alexj@backpocket.com",
	City:      "Plantation",
	Country:   "USA",
	Age:       48,
}

var Mia = &Person{
	FirstName: "Mia",
	LastName:  "Jeannopoulos",
	Email:     "mia@backpocket.com",
	City:      "Plantation",
	Country:   "USA",
	Age:       11,
}

var Karyn = &Person{
	FirstName: "Karyn",
	LastName:  "Jeannopoulos",
	Email:     "karyn@backpocket.com",
	City:      "Plantation",
	Country:   "USA",
	Age:       -1,
}

var Family = &People{Members: make([]*Person, 0)}

func init() {
	Family.Members = append(Family.Members, Alex)
	Family.Members = append(Family.Members, Mia)
	Family.Members = append(Family.Members, Karyn)
}

package main

import "fmt"

func (p *Person) CustomString() string {
	return fmt.Sprintf("%s %s - email: %s city: %s country: %s age: %d", p.FirstName, p.LastName, p.Email, p.City, p.Country, p.Age)
}

func (f *People) CustomString() string {
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

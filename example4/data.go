package main

import (
	"fmt"
	"math/rand"
	"time"
)
import "github.com/Pallinder/go-randomdata"

func (p *Person) CustomString() string {
	return fmt.Sprintf("%s %s - email: %s city: %s country: %s age: %d", p.FirstName, p.LastName, p.Email, p.City, p.Country, p.Age)
}

func (f *People) CustomString() string {
	return fmt.Sprintf("members: %d", len(f.Members))
}

var Homer = &Person{
	FirstName: "Homer",
	LastName:  "Simpson",
	Email:     "homer@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       36,
}

var Marge = &Person{
	FirstName: "Marge",
	LastName:  "Simpson",
	Email:     "marge@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       34,
}

var Bart = &Person{
	FirstName: "Bart",
	LastName:  "Simpson",
	Email:     "bart@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       10,
}
var Lisa = &Person{
	FirstName: "Lisa",
	LastName:  "Simpson",
	Email:     "lisa@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       8,
}
var Maggie = &Person{
	FirstName: "Maggie",
	LastName:  "Simpson",
	Email:     "maggie@simpsons.com",
	City:      "Springfield",
	Country:   "USA",
	Age:       1,
}

var Family = &People{Members: make([]*Person, 0)}
var Family5000 = &People{Members: make([]*Person, 0)}
var r1 *rand.Rand

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)

	Family.Members = append(Family.Members, Homer)
	Family.Members = append(Family.Members, Marge)
	Family.Members = append(Family.Members, Bart)
	Family.Members = append(Family.Members, Lisa)
	Family.Members = append(Family.Members, Maggie)

	for i := 0; i < 1000; i++ {
		Family5000.Members = append(Family5000.Members, Homer)
		Family5000.Members = append(Family5000.Members, Marge)
		Family5000.Members = append(Family5000.Members, Bart)
		Family5000.Members = append(Family5000.Members, Lisa)
		Family5000.Members = append(Family5000.Members, Maggie)
	}

}

func FakePeople(sz int) (peps *People) {
	peps = &People{Members: make([]*Person, 0)}

	for i := 0; i < sz; i++ {
		p := &Person{
			FirstName: randomdata.FirstName(randomdata.Male),
			LastName:  randomdata.LastName(),
			Email:     randomdata.Email(),
			City:      randomdata.City(),
			Country:   randomdata.Country(randomdata.ThreeCharCountry),
			Age:       r1.Int31n(99),
		}

		peps.Members = append(peps.Members, p)
	}
	return peps
}

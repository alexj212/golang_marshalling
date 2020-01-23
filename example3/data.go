package main

import "fmt"

type Person struct {
	FirstName     string `json:"first"`
	LastName      string `json:"last"`
	Email         string `json:"email"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Age           int32  `json:"age"`
	OmittedInt    int32  `json:"OmittedInt,omitempty"`
	OmittedString string `json:"OmittedString,omitempty"`
	Excluded      string `json:"-"`
	private       string
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s - email: %s city: %s country: %s age: %d", p.FirstName, p.LastName, p.Email, p.City, p.Country, p.Age)
}

type People struct {
	Members []*Person `json:"members"`
}

func (f *People) String() string {
	return fmt.Sprintf("members: %d", len(f.Members))
}

var Alex = &Person{
	FirstName:     "Alex",
	LastName:      "Jeannopoulos",
	Email:         "alexj@backpocket.com",
	City:          "Plantation",
	Country:       "USA",
	Age:           48,
	OmittedInt:    156,
	OmittedString: "not omitted",
	Excluded:      "hide me",
}

var Mia = &Person{
	FirstName:     "Mia",
	LastName:      "Jeannopoulos",
	Email:         "mia@backpocket.com",
	City:          "Plantation",
	Country:       "USA",
	Age:           11,
	OmittedInt:    156,
	OmittedString: "",
	Excluded:      "",
}

var Karyn = &Person{
	FirstName:     "Karyn",
	LastName:      "Jeannopoulos",
	Email:         "karyn@backpocket.com",
	City:          "Plantation",
	Country:       "USA",
	Age:           -1,
	OmittedInt:    156,
	OmittedString: "",
	Excluded:      "",
}

var Family = &People{Members: make([]*Person, 0)}

func init() {
	Family.Members = append(Family.Members, Alex)
	Family.Members = append(Family.Members, Mia)
	Family.Members = append(Family.Members, Karyn)
}

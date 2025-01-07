package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name      string
	Age       int
	Salary    int
	Education string
}

type People struct {
	People []Person
}

type Stats struct {
	AverageAge       float64
	AverageSalary    float64
	Youngest         []string
	Oldest           []string
	HighestPaid      []string
	LowestPaid       []string
	CountByEducation map[string]int
}

func main() {

	file, err := os.Open("people.json")

	lastThing := func() {
		fmt.Println("Closing file")
		file.Close()
	}

	defer lastThing()

	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	data, _ := io.ReadAll(file)

	var people People

	err = json.Unmarshal(data, &people)

	if err != nil {
		fmt.Println("Error unmarshalling data" + err.Error())
		return
	}
	//fmt.Println(people)

	fmt.Println(people.averageAge())
	fmt.Println(people.averageSalary())
	fmt.Println(people.NameofYouguest())
	fmt.Println(people.Nameofoldest())
	fmt.Println(people.Nameofhighestpaid())
	fmt.Println(people.Nameoflowestpaid())
	fmt.Println(people.countByEducation())

	stats := Stats{
		AverageAge:       people.averageAge(),
		AverageSalary:    people.averageSalary(),
		Youngest:         people.NameofYouguest(),
		Oldest:           people.Nameofoldest(),
		HighestPaid:      people.Nameofhighestpaid(),
		LowestPaid:       people.Nameoflowestpaid(),
		CountByEducation: people.countByEducation(),
	}

	statsJson, err := json.MarshalIndent(stats, "", " ")
	_ = os.WriteFile("stats.json", statsJson, 0644)

}

func (p *People) averageAge() float64 {
	var sum int
	for _, person := range p.People {
		sum += person.Age
	}
	return float64(sum) / float64(len(p.People))
}

func (p *People) averageSalary() float64 {
	var sum int
	for _, person := range p.People {
		sum += person.Salary
	}
	return float64(sum) / float64(len(p.People))
}

func (p *People) NameofYouguest() []string {
	var youguest Person
	res := []string{}
	for i, person := range p.People {
		if i == 0 {
			youguest = person
		}
		if person.Age <= youguest.Age {
			youguest = person
			res = append(res, person.Name)
		}
	}
	return res
}

func (p *People) Nameofoldest() []string {
	var youguest Person
	res := []string{}
	for i, person := range p.People {
		if i == 0 {
			youguest = person
		}
		if person.Age > youguest.Age {
			youguest = person
			res = append(res, person.Name)
		}
	}
	return res
}

func (p *People) Nameofhighestpaid() []string {
	var youguest Person
	res := []string{}
	for i, person := range p.People {
		if i == 0 {
			youguest = person
		}
		if person.Salary > youguest.Salary {
			youguest = person
			res = append(res, person.Name)
		}
	}
	return res
}

func (p *People) Nameoflowestpaid() []string {
	var youguest Person
	res := []string{}
	for i, person := range p.People {
		if i == 0 {
			youguest = person
		}
		if person.Salary < youguest.Salary {
			youguest = person
			res = append(res, person.Name)
		}
	}
	return res
}

func (p *People) countByEducation() map[string]int {
	res := map[string]int{}
	for _, person := range p.People {
		_, ok := res[person.Education]
		if ok {
			res[person.Education]++
		} else {
			res[person.Education] = 1
		}
	}
	return res

}

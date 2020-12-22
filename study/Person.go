package study

import (
	"fmt"
)

/*

 */
type person struct {
	name   string
	age    int
	iDCard string
}

type AgeError struct {
	age int
}

func CreatePerson(name string) *person {
	return &person{
		name: name,
	}
}

func (e *AgeError) Error() string {
	return fmt.Sprintf("人的年龄：%d不在范围内", e.age)
}

func (p *person) SetAge(age int) error {
	if age < 0 {
		return &AgeError{
			age: age,
		}
	}
	p.age = age
	return nil
}

func (p *person) SetIDCard(IDCard string) {
	p.iDCard = IDCard
}

type PersonList []*person

func (p PersonList) Len() int {
	return len(p)
}

func (p PersonList) Less(i, j int) bool {
	if p[i].age != p[j].age {
		return p[i].age-p[j].age > 0
	}
	return p[i].name < p[j].name
}

func (p PersonList) Swap(i, j int) {
	p1 := p[i]
	p[i] = p[j]
	p[j] = p1
}

func CreatePersonList(length int) PersonList {
	return PersonList(make([]*person, length))
}

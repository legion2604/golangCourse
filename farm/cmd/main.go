package main

import (
	"fmt"
	"golangCourse/farm/models"
)

type Animal interface {
	GetVoice()
	GetGender()
	GetAge()
}

type Cow models.Cow
type Chicken models.Chicken
type Dog models.Dog
type Cat models.Cat

func main() {
	cow := Cow{
		Gender: "female",
		Age:    5,
		Voice:  "moo",
		Weight: 700,
	}
	chicken := Chicken{
		Gender: "male",
		Age:    2,
		Voice:  "cluck",
	}
	dog := Dog{
		Gender:   "male",
		Age:      3,
		Voice:    "bark",
		Domestic: true,
	}
	cat := Cat{
		Gender:   "female",
		Age:      4,
		Voice:    "meow",
		Domestic: true,
	}
	MakeNoise(&cow)
	MakeNoise(&chicken)
	MakeNoise(&dog)
	MakeNoise(&cat)

}

func MakeNoise(a Animal) {
	a.GetVoice()
	a.GetGender()
	a.GetAge()
}

func (c *Cow) GetVoice()  { fmt.Println(c.Voice) }
func (c *Cow) GetGender() { fmt.Println(c.Gender) }
func (c *Cow) GetAge()    { fmt.Println(c.Age) }

func (c *Chicken) GetVoice()  { fmt.Println(c.Voice) }
func (c *Chicken) GetGender() { fmt.Println(c.Gender) }
func (c *Chicken) GetAge()    { fmt.Println(c.Age) }

func (c *Dog) GetVoice()  { fmt.Println(c.Voice) }
func (c *Dog) GetGender() { fmt.Println(c.Gender) }
func (c *Dog) GetAge()    { fmt.Println(c.Age) }

func (c *Cat) GetVoice()  { fmt.Println(c.Voice) }
func (c *Cat) GetGender() { fmt.Println(c.Gender) }
func (c *Cat) GetAge()    { fmt.Println(c.Age) }

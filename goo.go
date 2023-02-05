package main

import "fmt"

type teacher struct {
	Name       string
	Surname    string
	Experience int
}

func (t teacher) choose() {
	if t.Experience >= 5 {
		fmt.Println("Dear " + t.Name + " " + t.Surname + ", welcome to our school!")
	} else {
		fmt.Println("Dear " + t.Name + " " + t.Surname + ", unfortunately, we are looking for a candidate with a lot of experience!")
	}
}
func main() {
	res := teacher{
		Name:       "Madina",
		Surname:    "Tolegenkyzy",
		Experience: 12,
	}
	res1 := teacher{
		Name:       "Beibarys",
		Surname:    "Talgatov",
		Experience: 3,
	}
	res2 := teacher{
		Name:       "Ivan",
		Surname:    "Ivanov",
		Experience: 30,
	}
	res.choose()
	res1.choose()
	res2.choose()
}

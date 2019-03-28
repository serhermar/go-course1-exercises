package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []user

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	u1 := user{
		Name: "Fermina",
		Age:  35,
	}

	u2 := user{
		Name: "Gergio",
		Age:  34,
	}

	users := []user{u1, u2}

	fmt.Println(users)

	sort.Sort(ByAge(users))

	fmt.Println(users)

	sort.Sort(ByName(users))

	fmt.Println(users)

}

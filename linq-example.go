package main


import (
	"fmt"
	"github.com/ahmetb/go-linq"
)

type Car struct {
	year int
	owner, model string
}

func main() {
	var owners []string

	cars:=[]Car{ Car{year:2000, owner:"seb", model:"BMW"}, Car{year:2017, owner:"seb", model:"Mini"} }
	fmt.Println(cars)

	linq.From(cars).WhereT(func(c Car) bool {
		return c.year >= 2015
	}).SelectT(func(c Car) string {
		return c.owner
	}).ToSlice(&owners)

	fmt.Println(owners)


	var res string =  linq.From(cars).WhereT(func(c Car) bool {
		println("filter")
		return c.year >= 2015
	}).SelectT(func(c Car) string {
		println("select")
		return c.owner
	}).First().(string)
	fmt.Println("res: "+res)
}

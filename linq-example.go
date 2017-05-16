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

	cars:=[]Car{
		Car{year:2001, owner:"seb1", model:"car1"},
		Car{year:2017, owner:"seb2", model:"car2"},
		Car{year:2012, owner:"seb3", model:"car3"}}
	fmt.Println(cars)

	linq.From(cars).WhereT(func(c Car) bool {
		println("filter: "+c.owner)
		return c.year >= 2001
	}).
	SelectT(func(c Car) Car {
		println("do nothing: "+c.owner)
		return c
	}).
		SelectT(func(c Car) string {
		println("map to owner: "+c.owner)
		return c.owner
	}).SelectT(func(it interface{}) interface{}{
		println("peek: "+it.(string))
		return it
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

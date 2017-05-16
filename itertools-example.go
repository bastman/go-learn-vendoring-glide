package main

import (
	"fmt"
	"github.com/yanatan16/itertools"
)

func main() {

	srcData := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("srcData: ", srcData)
	fmt.Println("srcData.len: ", len(srcData))

	var iter itertools.Iter = itertools.New(srcData)
	fmt.Println("iter.len: ", len(iter))

	iter = itertools.Filter(func(interface{}) bool {
		return true
	}, iter)



	fmt.Println("iter: ", iter)
	fmt.Println("iter.len: ", len(iter))

	iter2 := itertools.Map(func (i interface{}) interface{} {
		return len(i.(string))
	}, iter)

	fmt.Println("iter: ", iter2)
	fmt.Println("iter.len: ", len(iter2))
	//for it := range iter2 {
		//println(it)
		//fmt.Println("item: ",(it.(int)).(string))
	//}

}

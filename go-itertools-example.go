package main
import (
	"fmt"
	"github.com/doloopwhile/go-itertools"
)

func values(vals ...int) chan int {
	ch := make(chan int)
	go func() {
		for _, x := range vals {
			ch <- x
		}
		close(ch)
	}()
	return ch
}
func array(ch <-chan int) []int {
	var arr []int
	for x := range ch {
		arr = append(arr, x)
	}
	return arr
}

func main() {
	println("==== TAKE ===")
	r := itertools.Take(values(1, 2, 3, 4, 5), 3)
	for it := range r {
		println(it)
	}
	println("==== Accumulate ===")
	ch := itertools.Accumulate(
		values(1, 2, 3, 4, 5),
		func(x int, y int) int { return x + y },
	)
	vals := array(ch)

	for it := range vals {
		fmt.Println("item: ",it)
	}
}

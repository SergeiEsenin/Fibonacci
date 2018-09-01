package main

import (
	"fmt"
	"time"
)

var rightGuess int
var wrong int
var timeoutChan (<-chan time.Time)
var intChan chan int

func getInp() {
	var v int
	fmt.Println("Insert your value:")
	fmt.Scan(&v)
	intChan <- v
}

func fibonacci() func() int {
	n := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case n == 0:
			n++
			ret = 0
		case n == 1:
			n++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}

func main() {
	var in int
	intChan = make(chan int)
	f := fibonacci()
	for i := 0; rightGuess < 10 && wrong < 3; i++ {
		var num int

		go getInp()
		select {
		case <-time.After(10 * time.Second):
			num = f()
			wrong++
			fmt.Println("Timeout!")
			fmt.Printf("You made %d mistake(s)\n", wrong)
		case in = <-intChan:
			num = f()
			if in == num {
				rightGuess++
				fmt.Printf("Right guess. Keep going %d remain \n", 10-rightGuess)
			} else if in != num {
				wrong++
				rightGuess = 0
				fmt.Printf("Wrong. You made %d mistake(s)\n", wrong)
			}
		}
		fmt.Printf("Index %d : Answer %d \n", i+1, num)

	}
}

package main

import "fmt"

func fizzBuzz(val int) (res string) {
	if val % 3 == 0 {
		res = res + "Fizz"
	}
	if val % 5 == 0 {
		res = res + "Buzz"
	}

	return
}

func main()  {
	for i := 0; i <= 100; i++ {
		if fizzBuzz(i) != "" {
			fmt.Println(fizzBuzz(i))
		} else {
			fmt.Println(i)
		}
	}
}

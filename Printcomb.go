package main

import "fmt"

func PrintComb2() {

	a := 0
	b := 0
	for n := 0; n < 10; n++ {
		for a <= 9 {
			b = a + 1
			for b <= 9 {
				fmt.Print(a / 10)
				fmt.Print(a % 10)
				fmt.Print("")
				fmt.Print(b / 10)
				fmt.Print(b % 10)
				if a < b {
					fmt.Print(",")
				}
				b++
			}
			a++
		}
	}

}

func main() {

	PrintComb2()
}


package main

import (
	"fmt"
)

func PrintComb2() {
	a := 0
	b := 0
	for a <= 9 {
		b = a + 1
		for b <= 9 {
			fmt.Printf("%d%d", a, b)
			if a < b {
				fmt.Print(",")
			}
			b++
		}
		a++
	}
	fmt.Println("\n")
}

func main() {
	PrintComb2()
}


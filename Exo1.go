package main

import "fmt"

func main() {
	T := 0
	fmt.Print("Entrez la température:")
	fmt.Scan(&T)

	if T <= 0 {
		fmt.Println("c'est de la glace")
	}
	if 0 < T && T < 25 {
		fmt.Println("C'est liquide")
	}

	if 25 < T && T < 100 {
		fmt.Println("c'est de la vapeur")
	}

}

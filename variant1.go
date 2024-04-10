package main

import "fmt"

func kakezan() {
	for x := 1; x <= 12; x++ {
		for y := 1; y <= 12; y++ {
			fmt.Printf("%4d", x*y)
		}
		fmt.Println()
	}
}

func main() {
	kakezan()
}

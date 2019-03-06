package main

import (
	"fmt"
	"time"
)

func getLetters(s string) {
	for _, c := range s {
		fmt.Println(string(c))
		time.Sleep(10 * time.Millisecond)
	}
}

func getNumbers(n []int) {
	for _, d := range n {
		fmt.Println(d)
		time.Sleep(15 * time.Millisecond)
	}
}

func main() {

	fmt.Println("Starting..")

	go getLetters("goLang")

	go getNumbers([]int{1, 2, 3, 4, 5})

	time.Sleep(300 * time.Millisecond)

	fmt.Println("Ending...")
}

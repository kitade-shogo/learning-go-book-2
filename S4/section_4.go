package main

import (
	"fmt"
	// "math/rand/v2"
)

func main() {
	//	var s []int
	//	for i := 0; i < 100; i++ {
	//		s = append(s, rand.IntN(100))
	//	}

	//	for _, v := range s {
	//		switch {
	//		case v%2 == 0 && v%3 == 0:
	//			fmt.Println("Six!")
	//		case v%2 == 0:
	//			fmt.Println("Two!")
	//		case v%3 == 0:
	//			fmt.Println("Three")
	//		default:
	//			fmt.Println("Never mind")
	//		}
	//	}

	var total int
	for i := 0; i < 10; i++ {
		total += 1
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Println(total)
}

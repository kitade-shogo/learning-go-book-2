package main

import "fmt"

func main() {
	//	===1===
	//	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	//	var ss1 = greetings[:2]
	//	var ss2 = greetings[1:4]
	//	var ss3 = greetings[3:]

	//	fmt.Println(greetings)
	//	fmt.Println(ss1)
	//	fmt.Println(ss2)
	//	fmt.Println(ss3)

	//	===2===
	//	var message string = "Hi 👩 and 👨"
	//	var rs []rune = []rune(message)
	//	var r rune = rs[3]
	//	var s string = string(r)
	//	fmt.Println(s)

	// 	===3===
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	me1 := struct {
		firstName string
		lastName  string
		id        int
	}{
		firstName: "Shogo",
		lastName:  "Kitade",
		id:        1,
	}

	me2 := Employee{
		firstName: "Shogo",
		lastName:  "kitade",
		id:        2,
	}

	var me3 struct {
		firstName string
		lastName  string
		id        int
	}

	me3.firstName = "Shogo"
	me3.lastName = "Kitade"
	me3.id = 3

	fmt.Println(me1)
	fmt.Println(me2)
	fmt.Println(me3)
}

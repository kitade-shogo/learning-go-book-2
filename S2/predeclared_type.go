package main

import "fmt"

func main() {
	//	===1===
	//	var i int = 20
	//	var f float64
	//	f = float64(i)
	//	fmt.Println(i, f)

	//	===2===
	//	const value = 10
	//	var i int = value
	//	var f float64 = value
	//	fmt.Println(i, f)

	var (
		b      byte   = 255
		smallI int32  = 2147483647
		bigI   uint64 = 18446744073709551615
	)

	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
}

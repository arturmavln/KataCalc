package main

import "fmt"

func main() {
	const (
		A int = 45

		B // скопирует А

		C float32 = 3.3

		D // скопирует С

		Sunday = iota

		Monday

		Tuesday

		Wednesday

		Thursday

		Friday

		Saturday
	)
	fmt.Println(Friday)
}

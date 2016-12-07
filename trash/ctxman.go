package main

import "fmt"

func main() {
	message := "Hello, world."
	var counts []int

	func() {
		defer func(x string) {
			message = x
		}(message)

		message = "Go to Tokyo!"
		fmt.Println(message)
		counts = append(counts, len(message))
	}()

	fmt.Println(message)
	counts = append(counts, len(message))

	fmt.Println(counts)
}

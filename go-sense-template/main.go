package main

import "fmt"

func main() {
	const name = "${{ values.name }}"

	fmt.Printf("Hello %s", name)
}
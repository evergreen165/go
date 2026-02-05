package main

import (
	"fmt"

	"github.com/speps/go-hashids/v2"
)

func main() {
	// 1. Setup Hashids data with salt and minimum length
	hd := hashids.NewData()
	hd.Salt = "secret"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)

	// 2. Encode integer(s) into a string
	id, _ := h.Encode([]int{1, 2, 3})
	fmt.Println("Encoded:", id) // Example: "nJ8rXzQj1B"

	// 3. Decode string back to integer(s)
	numbers, _ := h.DecodeWithError(id)
	fmt.Println("Decoded:", numbers) // Output: [1 2 3]
}

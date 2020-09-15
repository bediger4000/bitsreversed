package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%064b\n", n)

	var reversed uint64

	for i := 0; i < 64; i++ {
		reversed <<= 1
		bit := n & 0x1
		n >>= 1
		reversed |= bit
	}

	fmt.Printf("%064b\n", reversed)
}

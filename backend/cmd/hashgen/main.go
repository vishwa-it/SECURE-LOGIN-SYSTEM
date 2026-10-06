package main

import (
	"fmt"
	"log"
	"os"

	"secure-login-system/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: go run ./cmd/hashgen <password>")
	}

	hash, err := utils.HashPassword(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)
}

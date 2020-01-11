package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := flag.String("password", "", "password to be hashed")
	cost := flag.Int("cost", bcrypt.DefaultCost, "password hash cost")
	flag.Parse()

	passwordBytes := []byte(*password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, *cost)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	fmt.Println(string(hashedPassword))
}

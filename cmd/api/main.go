package main

import (
	"MyDrive/internal/env"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	fmt.Printf("ENV var = %s\n", env.GetString("STRONGEST_AVENGER", "NO"))
}

package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kailash360/BookStore-Backend/utils"
)

func main() {
	fmt.Println("Book Store")
	godotenv.Load()

	utils.Connect()

}

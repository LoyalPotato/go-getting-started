package main

import (
	"fmt"

	"github.com/LoyalPotato/go-getting-started/models"
)

func main() {
	u := models.User {
		ID: 2,
		FirstName: "Robbert",
		LastName: "Miller",
	}

	fmt.Println(u)
}
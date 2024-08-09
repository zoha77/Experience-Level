package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Enter your years of experience: ")
	fmt.Scanln(&input)

	years, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	determineExperienceLevel(years)
}
func determineExperienceLevel(years int) {
	if years < 0 {
		fmt.Println("Experience years cannot be negative.")
	} else if years <= 2 {
		fmt.Println("You are a Junior developer.")
	} else if years <= 5 {
		fmt.Println("You are a Mid-Level developer.")
	} else {
		fmt.Println("You are a Senior developer.")
	}
}

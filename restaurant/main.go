package main

import "fmt"

var dishes = []string{
	"Spaghetti Carbonara",
	"Grilled Salmon",
	"Caesar Salad",
	"Margherita Pizza",
	"Beef Steak",
}

func main() {
	fmt.Println("Dishes:")
	for _, dish := range dishes {
		fmt.Println(dish);
	}
}

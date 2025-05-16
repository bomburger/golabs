package main

import (
	"fmt"
	"time"
	"sync"
)

// dish -> cooking time
var dishes = map[string]float32{
	"Spaghetti Carbonara":  1,
	"Grilled Salmon": 2,
	"Caesar Salad": 1.5,
	"Margherita Pizza": 3,
	"Beef Steak": 3.5,
}

type Cook struct {
	Name string;
	CookingTime float32;
	DishesCooked int;
}

func NewCook(name string) Cook {
	c := Cook{}
	c.Name = name
	return c
}

func cook(dish string, cooks chan Cook, wg *sync.WaitGroup) {
	defer wg.Done()
	cook := <- cooks
	cookDuration := dishes[dish]
	time.Sleep(time.Duration(cookDuration) * time.Second)

	cook.CookingTime += cookDuration;
	cook.DishesCooked += 1

	fmt.Println("dish ", dish, " is cooked by ", cook.Name)
	cooks <- cook
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan Cook, 3) // free cooks
	ch <- NewCook("cook#1")
	ch <- NewCook("cook#2")
	ch <- NewCook("cook#3")
	fmt.Println("Restaurant is open!")
	start := time.Now()
	for dish := range dishes {
		wg.Add(1)
		go cook(dish, ch, wg)
	}
	wg.Wait()
	close(ch)

	fmt.Println("Restaurant is closed!")
	fmt.Println("Cook times:")
	for c := range ch {
		fmt.Printf("%s: cooking time: %.1f | dishes cooked: %d\n", c.Name, c.CookingTime, c.DishesCooked)
	}
	fmt.Printf("Total working time: %s\n", time.Since(start))
}

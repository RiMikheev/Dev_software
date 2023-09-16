package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, my dear friends")
	timer := 10
	for timer >= 0 {
		fmt.Println("Countdown... ", timer)
		time.Sleep(1 * time.Second)
		timer -= 1
	}
	fmt.Println("Bye")
}

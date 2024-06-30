package main

import (
	"KennyZeu69/particles"
	"fmt"
	"time"
)

func main() {
	// The NewArt is the whole system ...
	art := particles.NewArt(5, 3)
	art.Start()

	timer := time.NewTicker(100 * time.Millisecond)

	for {
		<-timer.C
		fmt.Print("\033[H\033[2J")
		art.Update()
		fmt.Println(art.Display())
	}
}

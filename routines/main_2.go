package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {

	// simple go routine
	go ProcessUsers()
}

func ProcessUsers() {

	i := 0

	for true {

		ProcessUser(i)

		i++
	}
}

func ProcessUser(userID int) {

    fmt.Println("I'm processing user", userID)

    time.Sleep(time.Duration(200 + rand.Intn(300)) * time.Millisecond)
}
package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println("The current time is:", current.Format(time.ANSIC))
	// 6 hours, 6 minutes, 6 seconds = 21966
	sss := time.Duration(21966 * time.Second)
	future := current.Add(sss)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be: ", future.Format(time.ANSIC))
}

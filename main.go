package main

import (
	// "fmt"
	//"runtime"
	"./driver"
	"time"
)

func main() {
	driver.Elev_init()
	driver.Elev_set_speed(0)

	for {
		//driver.Elev_set_speed(150)
		time.Sleep(time.Second)
	}
}

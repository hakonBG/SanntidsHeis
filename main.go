package main

import (
	// "fmt"
	//"runtime"
	"./driver"
	"./logic"
	//"time"
)

func main() {
	orderChan := make(chan logic.Order_call_s)
	floorSensorchan := make(chan int)
	passOrders := make(chan chan [3][4]int)
	exit := make(chan int)
	driver.Elev_init()
	go logic.Pull_panel_signals(orderChan, floorSensorchan)
	go logic.Handle_orders(orderChan, passOrders)
	go logic.Motor_control(passOrders, floorSensorchan)
	<-exit
}

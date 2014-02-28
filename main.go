package main

import (
	// "fmt"
	//"runtime"
	"./driver"
	"./logic"
	//"time"
	. "runtime"
)

func main() {
	GOMAXPROCS(NumCPU())
	orderChan := make(chan logic.Order_call_s)
	floorSensorchan := make(chan int)
	passOrders := make(chan chan logic.Orders_s)
	exit := make(chan int)
	driver.Elev_init()
	logic.Init_elevator()
	go logic.Poll_panel_signals(orderChan, floorSensorchan)
	go logic.Handle_orders(orderChan, passOrders)
	go logic.Motor_control(passOrders, floorSensorchan, orderChan)
	go logic.Adjust_lights(passOrders)
	<-exit
}

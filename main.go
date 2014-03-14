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
	addOrderChan, removeOrderChan := make(chan logic.Order_call_s), make(chan logic.Order_call_s)
	floorSensorchan := make(chan int)
	selfElevatorChan := make(chan logic.Elevator_s)
	passOrders := make(chan chan logic.Orders_s)
	exit := make(chan int)
	driver.Elev_init()
	logic.Init_elevator()
	go logic.Poll_panel_orders(addOrderChan, passOrders)
	go logic.Poll_floor_sensor_signal(floorSensorchan)
	go logic.Handle_orders(addOrderChan, removeOrderChan, passOrders)
	go logic.Motor_control(passOrders, floorSensorchan, removeOrderChan, selfElevatorChan)
	go logic.Adjust_lights(passOrders)
	<-exit
}

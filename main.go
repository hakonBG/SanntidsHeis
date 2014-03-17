package main

import (
	// "fmt"
	//"runtime"
	"./communication"
	"./driver"
	"./logic"
	"./ownVar"
	//"time"
	. "runtime"
)

func main() {
	GOMAXPROCS(NumCPU())
	addOrderChan, removeOrderChan := make(chan ownVar.Order_call_s), make(chan ownVar.Order_call_s)
	floorSensorchan := make(chan int)
	selfElevatorChan, addElevatorChan, removeElevatorChan := make(chan ownVar.Elevator_s), make(chan ownVar.Elevator_s), make(chan ownVar.Elevator_s)
	passElevators := make(chan chan map[string]ownVar.Elevator_s)
	passOrders := make(chan chan ownVar.Orders_s)
	exit := make(chan int)
	pushAddGlobalOrderChan := make(chan ownVar.Order_call_s)
	pushRemoveGlobalOrderChan := make(chan ownVar.Order_call_s)

	driver.Elev_init()
	logic.Init_elevator()

	assignedOrdersChan := make(chan [ownVar.N_GLOBALBUTTONTYPES][ownVar.N_FLOORS]string)

	go logic.Poll_panel_orders(addOrderChan, passOrders)
	go logic.Poll_floor_sensor_signal(floorSensorchan)
	go logic.Handle_orders(addOrderChan, removeOrderChan, passOrders, pushAddGlobalOrderChan, pushRemoveGlobalOrderChan)
	go logic.Motor_control(passOrders, floorSensorchan, removeOrderChan, selfElevatorChan)
	go logic.Adjust_lights(passOrders)

	go communication.Push_elevator(selfElevatorChan)
	go communication.Receive_elevator(addElevatorChan)
	go communication.Receive_add_global_order(addOrderChan, passOrders)
	go communication.Receive_remove_global_order(removeOrderChan, passOrders)
	go communication.Push_add_global_order(pushAddGlobalOrderChan)
	go communication.Push_remove_global_order(pushRemoveGlobalOrderChan)

	go communication.Handle_elevators(addElevatorChan, removeElevatorChan, passElevators)
	go communication.Find_best_elevator(passOrders, passElevators, addOrderChan, removeOrderChan, assignedOrdersChan)
	//go communication.Merge_global_orders(passOrders, passElevators, assignedOrdersChan, removeOrderChan, addOrderChan)

	<-exit
}

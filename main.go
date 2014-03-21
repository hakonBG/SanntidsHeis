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

	//Network Connection Chans

	//Init chans
	exitImNewChan := make(chan int)
	startElevatorProgram := make(chan int)
	ReceiveNewMsgChan := make(chan ownVar.Elevator_s)

	//Elev OrderChans
	ordersWhenNewChan := make(chan [ownVar.N_BUTTONTYPES][ownVar.N_FLOORS]int, 1)
	addOrderElevChan := make(chan ownVar.Order_call_s)
	addOrderUDPChan := make(chan ownVar.Order_call_s)
	addOrderCostChan := make(chan ownVar.Order_call_s)
	removeOrderElevChan := make(chan ownVar.Order_call_s)
	removeOrderUDPChan := make(chan ownVar.Order_call_s)
	removeOrderCostChan := make(chan ownVar.Order_call_s)

	newElevatorFoundChan := make(chan string)

	floorSensorchan := make(chan int)
	selfElevatorChan, addElevatorChan, removeElevatorChan := make(chan ownVar.Elevator_s), make(chan ownVar.Elevator_s), make(chan ownVar.Elevator_s)
	passElevators := make(chan chan map[string]ownVar.Elevator_s)
	passLostElevators := make(chan chan map[string]ownVar.Elevator_s)
	passOrders := make(chan chan ownVar.Orders_s)

	exit := make(chan bool)

	pushAddGlobalOrderChan := make(chan ownVar.Order_call_s)
	pushRemoveGlobalOrderChan := make(chan ownVar.Order_call_s)

	driver.Elev_init()
	logic.Init_elevator()

	//New elevator Routine

	go communication.Im_new_spam(
		exitImNewChan)

	go communication.Handle_msg_when_new(
		exitImNewChan,
		startElevatorProgram,
		ordersWhenNewChan,
		ReceiveNewMsgChan)
	go communication.Receive_msg_when_new(
		ReceiveNewMsgChan)

	<-startElevatorProgram

	//Communication Elevators
	go communication.Handle_elevators(
		addElevatorChan,
		removeElevatorChan,
		passElevators,
		passLostElevators)

	go communication.Assign_global_orders(
		passOrders,
		passElevators,
		addOrderCostChan,
		removeOrderCostChan)

	go communication.Update_new_elevator(
		passOrders,
		newElevatorFoundChan,
		passLostElevators)
	go communication.Find_new_elevator(
		newElevatorFoundChan)
	go communication.Detect_lost_elevators(
		removeElevatorChan,
		passElevators)

	//Elevator Control Routines
	go logic.Poll_panel_orders(
		addOrderElevChan,
		passOrders)
	go logic.Poll_floor_sensor_signal(
		floorSensorchan)
	go logic.Handle_orders(
		addOrderElevChan,
		addOrderUDPChan,
		addOrderCostChan,
		removeOrderElevChan,
		removeOrderUDPChan,
		removeOrderCostChan,
		passOrders,
		pushAddGlobalOrderChan,
		pushRemoveGlobalOrderChan,
		ordersWhenNewChan)
	go logic.Motor_control(
		passOrders,
		floorSensorchan,
		removeOrderElevChan,
		selfElevatorChan)
	go logic.Adjust_lights(
		passOrders)
	go logic.NotifyCtrlC()

	//Communication Orders Routines
	go communication.Push_elevator(
		selfElevatorChan)
	go communication.Receive_elevator(
		addElevatorChan)
	go communication.Receive_add_global_order(
		addOrderUDPChan,
		passOrders)
	go communication.Receive_remove_global_order(
		removeOrderUDPChan,
		passOrders)
	go communication.Push_add_global_order(
		pushAddGlobalOrderChan)
	go communication.Push_remove_global_order(
		pushRemoveGlobalOrderChan)

	//go communication.Merge_global_orders(passOrders, passElevators, assignedOrdersChan, removeOrderChan, addOrderChan)

	<-exit

}

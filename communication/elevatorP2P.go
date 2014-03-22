package communication

import (
	"./../ownVar"
	"fmt"
	"time"
)

const (
	N_SPAM = 10
)

func Receive_elevator(addElevatorChan chan<- ownVar.Elevator_s) {
	elevatorReadSocket := Set_up_udp_readSocket(ELEVATOR_STRUCT_PORT)

	var elevator ownVar.Elevator_s
	var msgElev []byte
	var connAddress string
	for {

		msgElev, connAddress = Udp_receive_msg(elevatorReadSocket)
		elevator = Json_decode_elevator(msgElev)
		elevator.Ip = connAddress
		elevator.LastTime = time.Now()
		addElevatorChan <- elevator
	}
}

func Push_elevator(selfElevatorChan <-chan ownVar.Elevator_s) {
	elevatorSendSocket := Set_up_udp_sendSocket(ELEVATOR_STRUCT_PORT)
	var elevator ownVar.Elevator_s
	var sendMsg []byte

	for {

		elevator = <-selfElevatorChan
		sendMsg = Json_encode_elevator(elevator)
		Udp_send_msg(elevatorSendSocket, sendMsg)
		<-time.After(25 * time.Millisecond)
	}
}

func Receive_add_global_order(
	addOrderUDPChan chan<- ownVar.Order_call_s,
	passOrders chan<- chan ownVar.Orders_s) {
	//Function Start

	//Channels
	passOrdersChan := make(chan ownVar.Orders_s)

	//Variables

	ownIp := Get_own_ip()
	globalOrderReadAddSocket := Set_up_udp_readSocket(ADD_GLOBAL_ORDER_PORT)
	var orderCall ownVar.Order_call_s
	var msgOrder []byte
	var address string
	var orders ownVar.Orders_s

	//Do
	for {
		select {
		case passOrders <- passOrdersChan:
			orders = <-passOrdersChan
		default:
			msgOrder, address = Udp_receive_msg(globalOrderReadAddSocket)
			if address != ownIp {
				orderCall = Json_decode_order(msgOrder)
				if orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] == 0 {
					orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] = 1
					addOrderUDPChan <- orderCall
				}
			}
		}
	}
}
func Receive_remove_global_order(
	removeOrderUDPChan chan<- ownVar.Order_call_s,
	passOrders chan<- chan ownVar.Orders_s) {
	//Start of function

	//Channels
	passOrdersChan := make(chan ownVar.Orders_s)

	//Variables
	ownIp := Get_own_ip()
	globalOrderReadRemoveSocket := Set_up_udp_readSocket(REMOVE_GLOBAL_ORDER_PORT)
	var orderCall ownVar.Order_call_s
	var msgOrder []byte
	var address string
	var orders ownVar.Orders_s

	//Do
	for {
		select {
		case passOrders <- passOrdersChan:
			orders = <-passOrdersChan
		default:
			msgOrder, address = Udp_receive_msg(globalOrderReadRemoveSocket)
			fmt.Println("recieved")
			if address != ownIp {
				fmt.Println("another elev whit global Order:", orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor])
				orderCall = Json_decode_order(msgOrder)
				if orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] == 1 {
					orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] = 0
					removeOrderUDPChan <- orderCall
				}
			}
		}
	}
}

func Push_add_global_order(pushAddGlobalOrderChan <-chan ownVar.Order_call_s) {

	globalOrderSendAddSocket := Set_up_udp_sendSocket(ADD_GLOBAL_ORDER_PORT)
	var orderCall ownVar.Order_call_s
	var sendMsg []byte
	for {
		orderCall = <-pushAddGlobalOrderChan
		sendMsg = Json_encode_order(orderCall)
		fmt.Println("push add")
		for i := 0; i < N_SPAM; i++ {
			Udp_send_msg(globalOrderSendAddSocket, sendMsg)
			<-time.After(5 * time.Millisecond)
		}
	}
}

func Push_remove_global_order(pushRemoveGlobalOrderChan <-chan ownVar.Order_call_s) {
	globalOrderSendRemoveSocket := Set_up_udp_sendSocket(REMOVE_GLOBAL_ORDER_PORT)
	var orderCall ownVar.Order_call_s
	var sendMsg []byte
	for {

		orderCall = <-pushRemoveGlobalOrderChan
		orderCall.OrderType = ownVar.GLOBAL
		fmt.Println("push remove")
		sendMsg = Json_encode_order(orderCall)
		for i := 0; i < N_SPAM; i++ {
			Udp_send_msg(globalOrderSendRemoveSocket, sendMsg)
			<-time.After(5 * time.Millisecond)
		}
	}
}

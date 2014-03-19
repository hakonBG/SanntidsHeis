package logic

import (
	//"./../driver"
	"./../ownVar"
	"fmt"
)

type call_type_t int

const (
	N_BUTTONTYPES       = 3
	N_GLOBALBUTTONTYPES = 2
	LAMP_ON             = 1
	LAMP_OFF            = 0
)

func Handle_orders(
	addOrderElevChan chan ownVar.Order_call_s,
	addOrderUDPChan chan ownVar.Order_call_s,
	addOrderCostChan chan ownVar.Order_call_s,
	removeOrderElevChan chan ownVar.Order_call_s,
	removeOrderUDPChan chan ownVar.Order_call_s,
	removeOrderCostChan chan ownVar.Order_call_s,
	passOrders chan chan ownVar.Orders_s,
	pushAddGlobalOrderChan chan ownVar.Order_call_s,
	pushRemoveGlobalOrderChan chan ownVar.Order_call_s) {
	//Start of function

	//Channels
	passOrdersChan := make(chan ownVar.Orders_s)

	//variables
	var orderCall ownVar.Order_call_s
	var orders ownVar.Orders_s

	orders.LocalOrders = Init_localOrders()
	orders.GlobalOrders = Init_globalOrders()

	//Do
	for {
		select {
		case orderCall = <-removeOrderElevChan:

			orders.LocalOrders[orderCall.ButtonType][orderCall.Floor] = 0
			orders.LocalOrders[BUTTON_COMMAND][orderCall.Floor] = 0
			if orderCall.OrderType == ownVar.GLOBAL {
				orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] = 0
				pushRemoveGlobalOrderChan <- orderCall
				fmt.Println("Remove global elev Order ")
			} else {
				fmt.Println("Remove COmmand ORder")
			}
			fmt.Println("ButtonType:", orderCall.ButtonType, "OrderFloor:", orderCall.Floor)

		case orderCall = <-removeOrderUDPChan:
			orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] = 0
			fmt.Println("Remove UDPglobal Order")
			fmt.Println("ButtonType:", orderCall.ButtonType, "OrderFloor:", orderCall.Floor)
		case orderCall = <-removeOrderCostChan:
			orders.LocalOrders[orderCall.ButtonType][orderCall.Floor] = 0
			fmt.Println("Remove localordercost")
			fmt.Println("ButtonType:", orderCall.ButtonType, "OrderFloor:", orderCall.Floor)
		case orderCall = <-addOrderElevChan:
			if orderCall.OrderType == ownVar.LOCAL {
				orders.LocalOrders[orderCall.ButtonType][orderCall.Floor] = 1
				fmt.Println("legger til local")
			} else if orderCall.OrderType == ownVar.GLOBAL {
				orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] = 1
				pushAddGlobalOrderChan <- orderCall

			}
			fmt.Println("ButtonType:", orderCall.ButtonType, "OrderFloor:", orderCall.Floor)

		case orderCall = <-addOrderUDPChan:
			orders.GlobalOrders[orderCall.ButtonType][orderCall.Floor] = 1
			fmt.Println("legger til UDPglobal")
			fmt.Println("ButtonType:", orderCall.ButtonType, "OrderFloor:", orderCall.Floor)
		case orderCall = <-addOrderCostChan:
			orders.LocalOrders[orderCall.ButtonType][orderCall.Floor] = 1
			fmt.Println("legger til localCost")
			fmt.Println("ButtonType:", orderCall.ButtonType, "OrderFloor:", orderCall.Floor)
		case passOrdersChan = <-passOrders:

			passOrdersChan <- orders
		}

	}

}

func Init_localOrders() [N_BUTTONTYPES][N_FLOORS]int {
	var localOrders [N_BUTTONTYPES][N_FLOORS]int
	for i := 0; i < N_BUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			localOrders[i][j] = 0
		}

	}
	return localOrders
}
func Init_globalOrders() [N_GLOBALBUTTONTYPES][N_FLOORS]int {
	var globalOrders [N_GLOBALBUTTONTYPES][N_FLOORS]int
	for i := 0; i < N_GLOBALBUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			globalOrders[i][j] = 0
		}

	}
	return globalOrders

}

func print_orders(orders ownVar.Orders_s) {
	fmt.Println("LOCAL ORDERS:")
	for i := 0; i < N_BUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			fmt.Printf("%b ", orders.LocalOrders[i][j])
		}
		fmt.Printf("\n")

	}
	fmt.Println("GLOBAL ORDERS:")
	for i := 0; i < N_BUTTONTYPES-1; i++ {
		for j := 0; j < N_FLOORS; j++ {
			fmt.Printf("%b ", orders.GlobalOrders[i][j])
		}
		fmt.Printf("\n")

	}
}

package logic

import (
	"./../driver"
	"fmt"
)

type call_type_t int
type order_type_t int

const (
	N_BUTTONTYPES       = 3
	N_GLOBALBUTTONTYPES = 2
	LAMP_ON             = 1
	LAMP_OFF            = 0
)

const (
	ADD_ORDER order_type_t = iota
	REMOVE_ORDER
)

type Order_call_s struct {
	orderType  order_type_t
	buttonType driver.Elev_button_type_t
	floor      int
}

type Orders_s struct {
	globalOrders [N_GLOBALBUTTONTYPES][N_FLOORS]int
	localOrders  [N_BUTTONTYPES][N_FLOORS]int
}

func Handle_orders(orderChan chan Order_call_s, passOrders chan chan Orders_s) {
	var orderCall Order_call_s
	var orders Orders_s
	passOrdersChan := make(chan Orders_s)
	orders.localOrders = init_localOrders()
	orders.globalOrders = init_globalOrders()
	for {
		select {
		case orderCall = <-orderChan:
			if orderCall.orderType == REMOVE_ORDER {
				orders.localOrders[orderCall.buttonType][orderCall.floor] = 0
				orders.localOrders[BUTTON_COMMAND][orderCall.floor] = 0
				fmt.Println("Remove Order")
			} else if orderCall.orderType == ADD_ORDER {

				if orderCall.buttonType == driver.BUTTON_CALL_UP || orderCall.buttonType == driver.BUTTON_CALL_DOWN {
					orders.localOrders[orderCall.buttonType][orderCall.floor] = 1
					fmt.Println("Legger til GLOBAL")
				} else {
					orders.localOrders[orderCall.buttonType][orderCall.floor] = 1
					fmt.Println("legger til command!!")
				}

			}
		case passOrdersChan = <-passOrders:
			fmt.Println("Passing orders")
			passOrdersChan <- orders
		}

	}

}

func init_localOrders() [N_BUTTONTYPES][N_FLOORS]int {
	var localOrders [N_BUTTONTYPES][N_FLOORS]int
	for i := 0; i < N_BUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			localOrders[i][j] = 0
		}

	}
	return localOrders
}
func init_globalOrders() [N_GLOBALBUTTONTYPES][N_FLOORS]int {
	var globalOrders [N_GLOBALBUTTONTYPES][N_FLOORS]int
	for i := 0; i < N_GLOBALBUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			globalOrders[i][j] = 0
		}

	}
	return globalOrders

}

func print_orders(orders Orders_s) {
	fmt.Println("LOCAL ORDERS:")
	for i := 0; i < N_BUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			fmt.Printf("%b ", orders.localOrders[i][j])
		}
		fmt.Printf("\n")

	}
	fmt.Println("GLOBAL ORDERS:")
	for i := 0; i < N_BUTTONTYPES-1; i++ {
		for j := 0; j < N_FLOORS; j++ {
			fmt.Printf("%b ", orders.globalOrders[i][j])
		}
		fmt.Printf("\n")

	}
}

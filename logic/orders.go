package logic

import (
	"./../driver"
	"fmt"
)

const (
	N_BUTTONTYPES = 3
	LAMP_ON       = 1
	LAMP_OFF      = 0
)

func Handle_orders(orderChan chan Order_call_s, passOrders chan chan [3][4]int) {
	var orderCall Order_call_s
	passOrdersChan := make(chan [3][4]int)
	orders := init_orders()
	for {
		select {
		case orderCall = <-orderChan:
			if orderCall.orderType == REMOVE_ORDER {
				orders[orderCall.buttonType][orderCall.floor] = 0
				driver.Elev_set_button_lamp(orderCall.buttonType, orderCall.floor, LAMP_OFF)

			} else if orderCall.orderType == ADD_ORDER {
				orders[orderCall.buttonType][orderCall.floor] = 1
				driver.Elev_set_button_lamp(orderCall.buttonType, orderCall.floor, LAMP_ON)
			}
		case passOrdersChan = <-passOrders:
			fmt.Println("orders")
			passOrdersChan <- orders
		}
	}

}

func init_orders() [3][4]int {
	var orders [3][4]int
	for i := 0; i < N_BUTTONTYPES; i++ {
		for j := 0; j < N_FLOORS; j++ {
			orders[i][j] = 0
		}

	}
	return orders
}

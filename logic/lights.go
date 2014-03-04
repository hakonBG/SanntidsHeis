package logic

import (
	"./../driver"
	"fmt"
	"time"
)

const (
	BUTTON_CALL_UP   = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND   = 2
)

func Adjust_lights(passOrders chan chan Orders_s) {
	var orders Orders_s
	lightOrders := make(chan Orders_s)
	for {
		select {
		case passOrders <- lightOrders:
			fmt.Println("Light Orders")
			orders = <-lightOrders

			set_lights(orders)
		case <-time.After(time.Second):
			print_orders(orders)
		}

		time.Sleep(25 * time.Millisecond)
	}

}
func set_lights(orders Orders_s) {

	for i := 0; i < N_FLOORS; i++ {
		if i < N_FLOORS-1 {
			driver.Elev_set_button_lamp(driver.BUTTON_CALL_UP, i, orders.localOrders[BUTTON_CALL_UP][i])
		}
		if i > 0 {
			driver.Elev_set_button_lamp(driver.BUTTON_CALL_DOWN, i, orders.localOrders[BUTTON_CALL_DOWN][i])
		}
		driver.Elev_set_button_lamp(driver.BUTTON_COMMAND, i, orders.localOrders[BUTTON_COMMAND][i])

	}

}

package logic

import (
	"./../driver"
	"./../ownVar"
	//"fmt"
	"time"
)

func Adjust_lights(passOrders chan chan ownVar.Orders_s) {
	var orders ownVar.Orders_s
	lightOrders := make(chan ownVar.Orders_s)
	for {
		select {
		case passOrders <- lightOrders:
			//fmt.Println("Light Orders")
			orders = <-lightOrders

			set_lights(orders)
		}
		<-time.After(25 * time.Millisecond)

	}

}
func set_lights(orders ownVar.Orders_s) {

	for i := 0; i < N_FLOORS; i++ {
		if i < N_FLOORS-1 {
			driver.Elev_set_button_lamp(driver.BUTTON_CALL_UP, i, orders.GlobalOrders[BUTTON_CALL_UP][i])
		}
		if i > 0 {
			driver.Elev_set_button_lamp(driver.BUTTON_CALL_DOWN, i, orders.GlobalOrders[BUTTON_CALL_DOWN][i])
		}
		driver.Elev_set_button_lamp(driver.BUTTON_COMMAND, i, orders.LocalOrders[BUTTON_COMMAND][i])

	}

}

package logic

import (
	"./../driver"
	"./../ownVar"
	"time"
)

const (
	N_FLOORS = 4
)

const (
	BUTTON_CALL_UP   = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND   = 2
)

func Poll_panel_orders(addOrderElevChan chan ownVar.Order_call_s, passOrders chan chan ownVar.Orders_s) {
	init_panel_lights()
	panelOrderChan := make(chan ownVar.Orders_s)
	var orders ownVar.Orders_s
	orders.LocalOrders = Init_localOrders()
	orders.GlobalOrders = Init_globalOrders()
	for {
		select {

		case passOrders <- panelOrderChan:

			orders = <-panelOrderChan

			for i := 0; i < N_FLOORS; i++ {
				if (driver.Elev_get_button_signal(driver.BUTTON_COMMAND, i) == 1) && (orders.LocalOrders[BUTTON_COMMAND][i] == 0) {

					addOrderElevChan <- order_call(i, driver.BUTTON_COMMAND)
					orders.LocalOrders[BUTTON_COMMAND][i] = 1
				}

				if (i > 0) && (driver.Elev_get_button_signal(driver.BUTTON_CALL_DOWN, i) == 1) && (orders.GlobalOrders[BUTTON_CALL_DOWN][i] == 0) {
					addOrderElevChan <- order_call(i, driver.BUTTON_CALL_DOWN)
					orders.GlobalOrders[BUTTON_CALL_DOWN][i] = 1
				}
				if (i < N_FLOORS-1) && (driver.Elev_get_button_signal(driver.BUTTON_CALL_UP, i) == 1) && (orders.GlobalOrders[BUTTON_CALL_UP][i] == 0) {
					addOrderElevChan <- order_call(i, driver.BUTTON_CALL_UP)
					orders.GlobalOrders[BUTTON_CALL_UP][i] = 1
				}

			}

		}

		time.Sleep(25 * time.Millisecond)
	}

}

func Poll_floor_sensor_signal(floorSensorChan chan int) {
	for {
		currentFloor := driver.Elev_get_floor_sensor_signal()

		if currentFloor != -1 {
			floorSensorChan <- currentFloor
			driver.Elev_set_floor_indicator(currentFloor)
		}
		time.Sleep(25 * time.Millisecond)

	}

}

func order_call(floor int, buttonType driver.Elev_button_type_t) ownVar.Order_call_s {

	var orderCall ownVar.Order_call_s
	orderCall.ButtonType = buttonType
	orderCall.Floor = floor
	if (buttonType == driver.BUTTON_CALL_UP) || (buttonType == BUTTON_CALL_DOWN) {
		orderCall.OrderType = ownVar.GLOBAL
	} else if buttonType == driver.BUTTON_COMMAND {
		orderCall.OrderType = ownVar.LOCAL
	}

	return orderCall
}

func init_panel_lights() {
	for i := 0; i < N_FLOORS; i++ {
		driver.Elev_set_button_lamp(driver.BUTTON_COMMAND, i, LAMP_OFF)
		if i > 0 {
			driver.Elev_set_button_lamp(driver.BUTTON_CALL_DOWN, i, LAMP_OFF)
		}
		if i < N_FLOORS-1 {
			driver.Elev_set_button_lamp(driver.BUTTON_CALL_UP, i, LAMP_OFF)
		}
	}

}

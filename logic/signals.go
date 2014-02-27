package logic

import (
	"./../driver"
)

type call_type_t int
type order_type_t int

const (
	N_FLOORS = 4
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

func Pull_panel_signals(OrderChan chan Order_call_s, floorSensorChan chan int) {
	init_panel_lights()
	for {

		if driver.Elev_get_floor_sensor_signal() > -1 {
			floorSensorChan <- driver.Elev_get_floor_sensor_signal()
		}

		for i := 0; i < N_FLOORS; i++ {
			if driver.Elev_get_button_signal(driver.BUTTON_COMMAND, i) == 1 {

				OrderChan <- order_call(i, driver.BUTTON_COMMAND)
			}
			if (i > 0) && (driver.Elev_get_button_signal(driver.BUTTON_CALL_DOWN, i) == 1) {
				OrderChan <- order_call(i, driver.BUTTON_CALL_DOWN)
			}
			if (i < N_FLOORS-1) && (driver.Elev_get_button_signal(driver.BUTTON_CALL_UP, i) == 1) {
				OrderChan <- order_call(i, driver.BUTTON_CALL_UP)
			}
		}
	}
	<-floorSensorChan

}

func order_call(floor int, buttonType driver.Elev_button_type_t) Order_call_s {
	var orderCall Order_call_s
	orderCall.buttonType = buttonType
	orderCall.floor = floor
	orderCall.orderType = ADD_ORDER
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

package logic

import (
	"./../driver"
	"time"
)

const (
	N_FLOORS = 4
)

func Poll_panel_orders(OrderChan chan Order_call_s) {
	init_panel_lights()

	for {

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

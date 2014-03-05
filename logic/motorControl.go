package logic

import (
	"./../driver"
	"fmt"
	"time"
)

type direction_t int

const (
	UP direction_t = iota
	DOWN
	STAND_STILL
)
const (
	BOTTOM_FLOOR       = 0
	TOP_FLOOR          = N_FLOORS - 1
	MOTOR_SPEED        = 150
	DURATION_DOOR_OPEN = 2
)

func Motor_control(passOrders chan chan Orders_s, floorSensorChan chan int, orderChan chan Order_call_s) {
	direction := DOWN
	var orders Orders_s
	var orderCall Order_call_s
	orderCall.orderType = REMOVE_ORDER
	var currentFloor, nextFloor int
	var breakDirection direction_t
	stopped := true
	readyToGo := true
	passOrdersChan := make(chan Orders_s)
	timeCheckpoint := time.Now().Add(-3 * time.Second)

	for {
		select {
		case currentFloor = <-floorSensorChan:
			//fmt.Println("Received floorsensor")
			nextFloor, direction = findNextStop(currentFloor, direction, orders.localOrders)

			//fmt.Printf("Next Floor: %b \n", nextFloor)
			//fmt.Printf("CurrentFloor: %b \n", currentFloor)
			if nextFloor == currentFloor {
				fmt.Println("NEXTFLOOR == CURRENT")
				fmt.Printf("Direction: %d\n", direction)
				stopped = stop_motor(breakDirection, stopped)
				timeCheckpoint = time.Now()
				if direction == UP {
					orderCall.floor = currentFloor
					orderCall.buttonType = BUTTON_CALL_UP
					orderChan <- orderCall
				} else {
					orderCall.floor = currentFloor
					orderCall.buttonType = BUTTON_CALL_DOWN
					orderChan <- orderCall
				}

			} else if nextFloor == -1 {
				stopped = stop_motor(breakDirection, stopped)
			} else if readyToGo {
				stopped = run_motor(direction, stopped)
				breakDirection = direction
			}
			if time.Now().After(timeCheckpoint.Add(DURATION_DOOR_OPEN * time.Second)) {
				driver.Elev_set_door_open_lamp(0)
				fmt.Printf("Direction: %d\n", direction)
				readyToGo = true
			} else {
				readyToGo = false
				driver.Elev_set_door_open_lamp(1)

			}

		case passOrders <- passOrdersChan:
			//fmt.Println("MotorOrders")
			orders = <-passOrdersChan

		}
		time.Sleep(25 * time.Millisecond)
	}
}

func run_motor(direction direction_t, stopped bool) bool {
	if direction == UP {
		driver.Elev_set_speed(MOTOR_SPEED)
	}
	if direction == DOWN {
		driver.Elev_set_speed(-MOTOR_SPEED)
	}
	return false
}

func stop_motor(breakDirection direction_t, stopped bool) bool {
	if !stopped {
		if breakDirection == UP {
			driver.Elev_set_speed(-100)
		} else if breakDirection == DOWN {
			driver.Elev_set_speed(100)
		}
	}

	<-time.After(50 * time.Millisecond)
	driver.Elev_set_speed(0)
	return true
}

func findNextStop(currentFloor int, direction direction_t, localOrders [N_BUTTONTYPES][N_FLOORS]int) (int, direction_t) {
	if direction == UP {
		for i := currentFloor; i < N_FLOORS; i++ {
			if (localOrders[BUTTON_CALL_UP][i] == 1) || (localOrders[BUTTON_COMMAND][i] == 1) {
				return i, UP
			}
		}
		for i := N_FLOORS - 1; i >= 0; i-- {
			if (localOrders[BUTTON_CALL_DOWN][i] == 1) || (localOrders[BUTTON_COMMAND][i] == 1) {
				if i > currentFloor {
					return i, UP
				} else {
					return i, DOWN
				}

			}
		}
		for i := 0; i < currentFloor; i++ {
			if localOrders[BUTTON_CALL_UP][i] == 1 {
				return i, DOWN
			}
		}

	} else if direction == DOWN {
		for i := currentFloor; i >= 0; i-- {
			if (localOrders[BUTTON_CALL_DOWN][i] == 1) || (localOrders[BUTTON_COMMAND][i] == 1) {
				return i, DOWN
			}

		}
		for i := 0; i < N_FLOORS; i++ {
			if (localOrders[BUTTON_CALL_UP][i] == 1) || (localOrders[BUTTON_COMMAND][i] == 1) {
				if i < currentFloor {
					return i, DOWN
				} else {
					return i, UP
				}

			}
		}
		for i := N_FLOORS - 1; i >= 0; i-- {
			if localOrders[BUTTON_CALL_DOWN][i] == 1 {
				return i, UP
			}
		}

	}
	if currentFloor == BOTTOM_FLOOR {
		return -1, UP
	} else if currentFloor == TOP_FLOOR {
		return -1, DOWN
	} else {
		return -1, direction
	}

}

func Init_elevator() {
	for {
		driver.Elev_set_speed(-150)
		if driver.Elev_get_floor_sensor_signal() != -1 {
			break
		}
	}
	driver.Elev_set_speed(0)

}

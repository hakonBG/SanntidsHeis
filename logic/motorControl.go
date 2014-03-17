package logic

import (
	"./../driver"
	"./../ownVar"
	"fmt"
	"time"
)

const (
	BOTTOM_FLOOR       = 0
	TOP_FLOOR          = N_FLOORS - 1
	MOTOR_SPEED        = 150
	DURATION_DOOR_OPEN = 2
)

//Equal as the type in communication

func Motor_control(
	passOrders chan chan ownVar.Orders_s,
	floorSensorChan chan int,
	removeOrderChan chan ownVar.Order_call_s,
	selfElevatorChan chan ownVar.Elevator_s) {
	//Start of function

	//Channels
	passOrdersChan := make(chan ownVar.Orders_s)

	//Variables
	direction := ownVar.DOWN
	var elevator ownVar.Elevator_s
	var orders ownVar.Orders_s
	var orderCall ownVar.Order_call_s
	orderCall.OrderType = ownVar.LOCAL
	var currentFloor, nextFloor int
	var breakDirection ownVar.Direction_t
	stopped := true
	readyToGo := true
	timeCheckpoint := time.Now().Add(-3 * time.Second)

	for {
		select {
		case currentFloor = <-floorSensorChan:

			nextFloor, direction = findNextStop(currentFloor, direction, orders.LocalOrders)

			//Update Elevator variables
			elevator.CurrentFloor = currentFloor
			elevator.NextFloor = nextFloor
			elevator.Direction = direction

			if nextFloor == currentFloor {
				fmt.Println("NEXTFLOOR == CURRENT")
				//fmt.Printf("Direction: %d\n", direction)
				stopped = stop_motor(breakDirection, stopped)
				elevator.Moving = false
				timeCheckpoint = time.Now()

				if direction == ownVar.UP {
					orders.LocalOrders[BUTTON_CALL_UP][currentFloor] = 0
					orders.LocalOrders[BUTTON_COMMAND][currentFloor] = 0
					orderCall.Floor = currentFloor
					orderCall.ButtonType = driver.BUTTON_CALL_UP
					removeOrderChan <- orderCall
				} else if direction == ownVar.DOWN {
					orders.LocalOrders[BUTTON_CALL_DOWN][currentFloor] = 0
					orders.LocalOrders[BUTTON_COMMAND][currentFloor] = 0
					orderCall.Floor = currentFloor
					orderCall.ButtonType = driver.BUTTON_CALL_DOWN
					removeOrderChan <- orderCall
				}

			} else if nextFloor == -1 {
				stopped = stop_motor(breakDirection, stopped)
				elevator.Moving = false
				elevator.Direction = ownVar.IDLE
			} else if readyToGo {
				stopped = run_motor(direction, stopped)
				elevator.Moving = true
				breakDirection = direction

			}
			if time.Now().After(timeCheckpoint.Add(DURATION_DOOR_OPEN * time.Second)) {
				driver.Elev_set_door_open_lamp(0)
				//fmt.Printf("Direction: %d\n", direction)
				readyToGo = true
			} else {
				readyToGo = false
				driver.Elev_set_door_open_lamp(1)

			}

		case passOrders <- passOrdersChan:
			//fmt.Println("MotorOrders")
			orders = <-passOrdersChan
			elevator.Orders[BUTTON_COMMAND] = orders.LocalOrders[BUTTON_COMMAND]
			elevator.Orders[BUTTON_CALL_UP] = orders.GlobalOrders[BUTTON_CALL_UP]
			elevator.Orders[BUTTON_CALL_DOWN] = orders.GlobalOrders[BUTTON_CALL_DOWN]

		case selfElevatorChan <- elevator:

		}

	}
}

func run_motor(direction ownVar.Direction_t, stopped bool) bool {
	if direction == ownVar.UP {
		driver.Elev_set_speed(MOTOR_SPEED)
	}
	if direction == ownVar.DOWN {
		driver.Elev_set_speed(-MOTOR_SPEED)
	}
	return false
}

func stop_motor(breakDirection ownVar.Direction_t, stopped bool) bool {
	if !stopped {
		if breakDirection == ownVar.UP {
			driver.Elev_set_speed(-100)
		} else if breakDirection == ownVar.DOWN {
			driver.Elev_set_speed(100)
		}
	}

	<-time.After(50 * time.Millisecond)
	driver.Elev_set_speed(0)
	return true
}

func findNextStop(
	currentFloor int,
	direction ownVar.Direction_t,
	localOrders [N_BUTTONTYPES][N_FLOORS]int) (int, ownVar.Direction_t) {

	//Start of function
	if direction == ownVar.UP {
		for i := currentFloor; i < N_FLOORS; i++ {
			if (localOrders[driver.BUTTON_CALL_UP][i] == 1) || (localOrders[driver.BUTTON_COMMAND][i] == 1) {
				return i, ownVar.UP
			}
		}
		for i := N_FLOORS - 1; i >= 0; i-- {
			if (localOrders[driver.BUTTON_CALL_DOWN][i] == 1) || (localOrders[driver.BUTTON_COMMAND][i] == 1) {
				if i > currentFloor {
					return i, ownVar.UP
				} else {
					return i, ownVar.DOWN
				}

			}
		}
		for i := 0; i < currentFloor; i++ {
			if localOrders[driver.BUTTON_CALL_UP][i] == 1 {
				return i, ownVar.DOWN
			}
		}

	} else if direction == ownVar.DOWN {
		for i := currentFloor; i >= 0; i-- {
			if (localOrders[driver.BUTTON_CALL_DOWN][i] == 1) || (localOrders[driver.BUTTON_COMMAND][i] == 1) {
				return i, ownVar.DOWN
			}

		}
		for i := 0; i < N_FLOORS; i++ {
			if (localOrders[driver.BUTTON_CALL_UP][i] == 1) || (localOrders[driver.BUTTON_COMMAND][i] == 1) {
				if i < currentFloor {
					return i, ownVar.DOWN
				} else {
					return i, ownVar.UP
				}

			}
		}
		for i := N_FLOORS - 1; i >= 0; i-- {
			if localOrders[driver.BUTTON_CALL_DOWN][i] == 1 {
				return i, ownVar.UP
			}
		}

	}
	if currentFloor == BOTTOM_FLOOR {
		return -1, ownVar.UP
	} else if currentFloor == TOP_FLOOR {
		return -1, ownVar.DOWN
	} else {
		return -1, direction
	}

}

func Init_elevator() {

	position := driver.Elev_get_floor_sensor_signal()

	if position == -1 {
		for {
			driver.Elev_set_speed(-150)
			if driver.Elev_get_floor_sensor_signal() != -1 {
				break
			}
		}
		driver.Elev_set_speed(150)
		<-time.After(25 * time.Millisecond)
		driver.Elev_set_speed(0)
	}

}

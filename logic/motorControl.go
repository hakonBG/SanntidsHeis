package logic

import (
	"./../driver"
	"fmt"
)

type direction_t int

const (
	UP direction_t = iota
	DOWN
)

func Motor_control(passOrders chan chan [3][4]int, floorSensorChan chan int) {
	direction := DOWN
	var orders [3][4]int
	var floorSensor int
	passOrdersChan := make(chan [3][4]int)
	set_motor_speed(direction)
	for {
		select {
		case floorSensor = <-floorSensorChan:
			if floorSensor == 0 {
				direction = UP
				set_motor_speed(direction)

			} else if floorSensor == 3 {
				direction = DOWN
				set_motor_speed(direction)
			}
		case passOrdersChan = <-passOrders:
			fmt.Println("hei")
			passOrdersChan <- orders
		}
	}
}

func set_motor_speed(direction direction_t) {
	if direction == UP {
		driver.Elev_set_speed(150)
	}
	if direction == DOWN {
		driver.Elev_set_speed(-150)
	}
}

package communication

import (
	"./../driver"
	"./../logic"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type Status_t int

type Elevator_s struct {
	Direction    logic.Direction_t
	Moving       bool
	Orders       [N_BUTTONTYPES][N_FLOORS]int
	LastTime     time.Time
	Ip           string
	NextFloor    int
	CurrentFloor int
}

const (
	UP Direction_t = iota
	DOWN
	IDLE
)

const (
	BUTTON_CALL_UP   = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND   = 2
)

func Handle_elevators(addElevatorChan chan Elevator_s, removeElevatorChan chan Elevator_S, passElevators chan chan map[string]Elevator_s) {
	var elevators map[string]Elevator_s
	var elevator, elevatorSelf, elevatorSelfTemp Elevator_s
	elevatorSelf
	findBestElevatorChan := make(chan Elevator_s)
	for {
		select {
		case findBestElevatorChan = <-passElevators:
			findBestElevatorChan <- elevators

		case elevator = <-addElevatorChan:
			elevators[elevator.Ip] = elevator
		case elevator = <-removeElevatorChan:
			delete(elevators, elevator.Ip)
		}

	}

}

func initializeElevator(ip string) Elevator_s {
	var elev Elevator_s
	elev.Ip = ip
	elev.Status = IDLE
	elev.NextFloor = 1
	elev.LastTime = 0
	elev.CurrentFloor = -1
}

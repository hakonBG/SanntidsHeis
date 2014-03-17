package communication

import (
	"./../ownVar"
	//"fmt"
	//"math"
	"time"
)

const (
	BUTTON_CALL_UP   = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND   = 2
)

func Handle_elevators(
	addElevatorChan chan ownVar.Elevator_s,
	removeElevatorChan chan ownVar.Elevator_s,
	passElevators chan chan map[string]ownVar.Elevator_s) {
	//Start of function

	//Channels
	findBestElevatorChan := make(chan map[string]ownVar.Elevator_s)
	elevators := make(map[string]ownVar.Elevator_s)

	//Variables
	var elevator ownVar.Elevator_s

	//Do
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

func initializeElevator(ip string) ownVar.Elevator_s {
	var elev ownVar.Elevator_s
	elev.Ip = ip
	elev.Direction = ownVar.IDLE
	elev.NextFloor = 1
	elev.LastTime = time.Now()
	elev.CurrentFloor = -1
	return elev
}

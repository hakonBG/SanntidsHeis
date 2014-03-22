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
	addElevatorChan <-chan ownVar.Elevator_s,
	removeElevatorChan <-chan ownVar.Elevator_s,
	passElevators <-chan chan map[string]ownVar.Elevator_s,
	passLostElevators <-chan chan map[string]ownVar.Elevator_s) {
	//Start of function

	//Channels
	passElevatorsChan := make(chan map[string]ownVar.Elevator_s)
	elevators := make(map[string]ownVar.Elevator_s)
	lostElevators := make(map[string]ownVar.Elevator_s)
	//Variables
	var elev ownVar.Elevator_s

	//Do
	for {
		select {
		case passElevatorsChan = <-passElevators:
			passElevatorsChan <- elevators
		case passElevatorsChan = <-passLostElevators:
			passElevatorsChan <- lostElevators
		case elev = <-addElevatorChan:
			if check_for_elevator_in_map(elev, lostElevators) {
				delete(lostElevators, elev.Ip)
			}
			elevators[elev.Ip] = elev
		case elev = <-removeElevatorChan:
			delete(elevators, elev.Ip)
			lostElevators[elev.Ip] = elev
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

func check_for_elevator_in_map(elev ownVar.Elevator_s, elevators map[string]ownVar.Elevator_s) bool {
	for ip, _ := range elevators {
		if ip == elev.Ip {
			return true
		}
	}
	return false
}

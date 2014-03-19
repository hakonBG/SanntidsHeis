package communication

import (
	"./../ownVar"
	"fmt"
	"time"
)

func Detect_lost_elevators(
	removeElevatorChan chan ownVar.Elevator_s,
	passElevators chan chan map[string]ownVar.Elevator_s) {

	//Start of function

	//Channels
	passElevatorsChan := make(chan map[string]ownVar.Elevator_s)

	//Variables
	var elevators map[string]ownVar.Elevator_s
	for {
		select {
		case passElevators <- passElevatorsChan:
			elevators = <-passElevatorsChan
			for _, elev := range elevators {
				if time.Now().Sub(elev.LastTime) > time.Second {
					removeElevatorChan <- elev
					fmt.Println("Elev Lost:", elev.Ip)
				}
			}
		}
	}
}

/*func Recover_lost_elevator(
	passElevators chan chan map[string]ownVar.Elevator_s,
	passLostElevators chan chan map[string]ownVar.Elevator_s) {
	//Start of function

	//Channels
	passElevatorsChan := make(chan map[string]ownVar.Elevator_s)
	passLostElevatorsChan := make(chan map[string]ownVar.Elevator_s)

	//Variables
	elevators := make(map[string]ownVar.Elevator_s)
	lostElevators := make(map[string]ownVar.Elevator_s)

	//Do
	for {
		select {
		case passElevators <- passElevatorsChan:
			elevators = <-passElevatorsChan
		case passLostElevators <- passLostElevatorsChan:
			lostElevators = <-passLostElevatorsChan

		default:
			for _, elev := range elevators {
				for _, lostelev := range lostElevators {
					if elev.Ip == lostelev.Ip {

					}
				}
			}
		}
	}
}

func Update_lost_elevator(lostElev ownVar.Elevator_s, updateLostElevatorChan chan ownVar.Elevator_s) {

}
*/

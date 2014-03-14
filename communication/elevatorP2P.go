package communication

import (
	"./../ownVar"
	"fmt"
	"time"
)

func Receive_elevator(addElevatorChan chan ownVar.Elevator_s) {
	elevatorReadSocket := Set_up_udp_readSocket(ELEVATOR_STRUCT_PORT)

	var elevator ownVar.Elevator_s
	var msgElev []byte
	var connAddress string
	for {

		msgElev, connAddress = Udp_receive_msg(elevatorReadSocket)

		elevator = Json_decode_elevator(msgElev)
		elevator.Ip = connAddress
		elevator.LastTime = time.Now()
		addElevatorChan <- elevator

	}
}

func Push_elevator(selfElevatorChan chan ownVar.Elevator_s) {
	elevatorSendSocket := Set_up_udp_sendSocket(ELEVATOR_STRUCT_PORT)
	var elevator ownVar.Elevator_s
	var sendMsg []byte

	for {
		fmt.Println("Pusher heis")
		elevator = <-selfElevatorChan

		sendMsg = Json_encode_elevator(elevator)
		Udp_send_msg(elevatorSendSocket, sendMsg)

	}

}

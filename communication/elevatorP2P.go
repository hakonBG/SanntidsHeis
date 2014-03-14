package communication

import (
	"fmt"
	"time"
)

func Receive_elevator(addElevatorChan chan Elevator_s) {
	elevatorReadSocket := Set_up_udp_readSocket(ELEVATOR_READ_PORT)
	var stringElev string
	var elevator Elevator_s

	for {
		stringElev, connAddress = Udp_receive_msg(elevatorReadSocket)
		elevator = Json_decode_elevator(stringElev)
		elevator.Ip = connAddress
		elevator.LastTime = time.Now()
		addElevatorChan <- elevator

	}
}

func Push_elevator(selfElevatorChan chan Elevator_s) {
	elevatorSendSocket := Set_up_udp_sendSocket(ELEVATOR_SEND_PORT)
	var elevator Elevator_s
	var sendMsg []byte
	for {

		elevator <- selfElevatorChan

		sendMsg := Json_encode_elevator(elevator)
		Udp_send_msg(elevatorSendSocket, sendMsg)

	}

}

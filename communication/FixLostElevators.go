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

func Update_new_elevator(
	passOrders chan chan ownVar.Orders_s,
	newElevatorFoundChan chan string,
	passLostElevators chan chan map[string]ownVar.Elevator_s) {
	//Start Of Function

	//Sockets
	newElevatorSocket := Set_up_udp_sendSocket(NEW_ELEVATOR_PORT)

	//Channels
	passOrdersChan := make(chan ownVar.Orders_s)
	passLostElevatorsChan := make(chan map[string]ownVar.Elevator_s)

	//Variables
	var address string
	var lostElevators map[string]ownVar.Elevator_s
	var orders ownVar.Orders_s
	newElev := ownVar.Elevator_s{}

	//Do
	for {
		select {
		case passLostElevators <- passLostElevatorsChan:
			lostElevators = <-passLostElevatorsChan
		case passOrders <- passOrdersChan:
			orders = <-passOrdersChan
		case address = <-newElevatorFoundChan:
			newElev.Ip = address
			newElev.Orders[BUTTON_CALL_DOWN] = orders.GlobalOrders[BUTTON_CALL_DOWN]
			newElev.Orders[BUTTON_CALL_UP] = orders.GlobalOrders[BUTTON_CALL_UP]
			if check_for_elevator_in_map(newElev, lostElevators) {
				newElev.Orders[BUTTON_COMMAND] = lostElevators[address].Orders[BUTTON_COMMAND]
			} else {
				newElev.Orders[BUTTON_COMMAND] = emptyCommandOrderList()
			}
			msg := Json_encode_elevator(newElev)
			for i := 0; i < 5; i++ {
				Udp_send_msg(newElevatorSocket, msg)
			}

		}
		<-time.After(25 * time.Millisecond)

	}
}

func Find_new_elevator(newElevatorFoundChan chan string) {
	discoverNewElevatorSocket := Set_up_udp_readSocket(NEW_ELEVATOR_SPAM_PORT)
	var address string
	for {
		_, address = Udp_receive_msg(discoverNewElevatorSocket)
		newElevatorFoundChan <- address
		<-time.After(25 * time.Millisecond)
	}

}

func Receive_msg_when_new(
	ReceiveNewMsgChan chan ownVar.Elevator_s) {
	//Start of function

	//Sockets
	newOrdersSocket := Set_up_udp_readSocket(NEW_ELEVATOR_PORT)

	for {

		msg, _ := Udp_receive_msg(newOrdersSocket)

		elev := Json_decode_elevator(msg)
		ReceiveNewMsgChan <- elev
	}

}

func Im_new_spam(exitImNewChan chan int) {
	imNewSocket := Set_up_udp_sendSocket(NEW_ELEVATOR_SPAM_PORT)
	myStr := "im new"
	sendMsg := []byte(myStr)
	finished := false
	for {
		select {
		case <-exitImNewChan:

			finished = true
			//imNewSocket.Close()
			break
		default:
			Udp_send_msg(imNewSocket, sendMsg)
		}
		if finished {
			break
		}
	}
	fmt.Println("Stop spamming")

}

func Handle_msg_when_new(
	exitImNewChan chan int,
	startElevatorProgram chan int,
	ordersWhenNewChan chan [N_BUTTONTYPES][N_FLOORS]int,
	ReceiveNewMsgChan chan ownVar.Elevator_s) {
	//Start of function

	//Variables
	var elev ownVar.Elevator_s
	ownIp := Get_own_ip()
	timer := time.Now()
	finished := false
	for {

		select {
		case elev = <-ReceiveNewMsgChan:
			if elev.Ip == ownIp {
				fmt.Println("elevip = ownip")
				ordersWhenNewChan <- elev.Orders
				exitImNewChan <- 1
				startElevatorProgram <- 1
				finished = true
				break
			}
		default:
			if time.Now().Sub(timer) > 2*time.Second {
				fmt.Println("timer out")
				exitImNewChan <- 1
				startElevatorProgram <- 1
				finished = true
				break
			}

		}
		if finished {
			break
		}

	}
	fmt.Println("finished handle msg")

}

func emptyCommandOrderList() [N_FLOORS]int {
	var list [N_FLOORS]int
	for i := 0; i < N_FLOORS; i++ {
		list[i] = 0
	}
	return list

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

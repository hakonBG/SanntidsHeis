package communication

import (
	"./../ownVar"
	"fmt"
	"net"
	"time"
)

func Handle_msg_when_network_down(
	exitImNewChan chan int,
	newGlobalOrdersChan chan [N_GLOBALBUTTONTYPES][N_FLOORS]int,
	ReceiveNewMsgChan chan ownVar.Elevator_s,
	ownIp string) {

	//Variables

	var elev ownVar.Elevator_s
	timer := time.Now()
	finished := false
	var newGlobalOrders [N_GLOBALBUTTONTYPES][N_FLOORS]int

	for {

		select {
		case elev = <-ReceiveNewMsgChan:
			fmt.Println("msg received")
			if elev.Ip == ownIp {
				fmt.Println("elevip = ownip")
				newGlobalOrders[BUTTON_CALL_UP] = elev.Orders[BUTTON_CALL_UP]
				newGlobalOrders[BUTTON_CALL_DOWN] = elev.Orders[BUTTON_CALL_DOWN]
				newGlobalOrdersChan <- newGlobalOrders

				exitImNewChan <- 1
				finished = true
				break
			}
		default:

			if time.Now().Sub(timer) > 2*time.Second {
				fmt.Println("timer out")
				exitImNewChan <- 1
				finished = true
				break
			}

		}

		if finished {
			break
		}

	}

}

func Check_network_connection_down(networkUpAgainChan chan bool) {

	//Sockets
	networkUpAddress, err := net.ResolveUDPAddr(CONN_TYPE_UDP, CONN_HOST+":"+CHECK_NETWORK_CONN_PORT)
	Check_error(err)
	networkUpSocket, err := net.DialUDP(CONN_TYPE_UDP, nil, networkUpAddress)
	Check_error(err)

	//Variables
	NetworkUp := true
	for {
		_, err = networkUpSocket.Write([]byte("hei"))
		if NetworkUp && (err != nil) {
			NetworkUp = false

		} else if (err == nil) && (NetworkUp == false) {
			NetworkUp = true
			networkUpAgainChan <- true
		}
		<-time.After(25 * time.Millisecond)

	}

}

func Get_global_orders_if_network_down(
	networkUpAgainChan chan bool,
	exitImNewChan chan int,
	newGlobalOrdersChan chan [N_GLOBALBUTTONTYPES][N_FLOORS]int,
	ReceiveNewMsgChan chan ownVar.Elevator_s,
	imNewSocket *net.UDPConn,
	newOrdersSocket *net.UDPConn,
	ownIp string) {
	//Start of function

	for {
		<-networkUpAgainChan

		go Im_new_spam(
			exitImNewChan,
			imNewSocket)

		go Receive_msg_when_new(
			ReceiveNewMsgChan,
			newOrdersSocket)

		go Handle_msg_when_network_down(
			exitImNewChan,
			newGlobalOrdersChan,
			ReceiveNewMsgChan,
			ownIp)

	}
}

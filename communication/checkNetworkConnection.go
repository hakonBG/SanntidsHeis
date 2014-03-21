package communication

import (
	"fmt"
	"net"
)

func Check_network_connection_up(networkUpChan chan bool) {
	networkUpAddress, _ := net.ResolveUDPAddr(CONN_TYPE_UDP, CONN_HOST+":"+CHECK_NETWORK_CONN_PORT)
	//Check_error(err)

	for {
		fmt.Println("Checking Connection")
		networkUpSocket, err := net.DialUDP(CONN_TYPE_UDP, nil, networkUpAddress)
		if err != nil {
			networkUpSocket.Close()

		} else {
			fmt.Println("err = nil")
			networkUpSocket.Close()
			networkUpChan <- true
			break

		}

	}

}

package communication

import (
	. "fmt"
	"net"
	"strings"

	//. "strings"
	//. "os"
)

const (
	CONN_HOST            = "129.241.187.255"
	ELEVATOR_STRUCT_PORT = "32373"

	CONN_TYPE = "udp4"
)

func check_error(err error) {
	if err != nil {
		Println("Error dialing", err.Error())
	}

}

func Set_up_udp_readSocket(port string) *net.UDPConn {
	Println("Setup Readsocket")

	udpReadAddress, err := net.ResolveUDPAddr(CONN_TYPE, CONN_HOST+":"+port)
	check_error(err)
	udpReadSocket, err := net.ListenUDP(CONN_TYPE, udpReadAddress)
	check_error(err)
	Println("read fin")
	return udpReadSocket

}

func Set_up_udp_sendSocket(port string) *net.UDPConn {
	Println("Setup Broadsocket")

	udpBroadcastAddress, err := net.ResolveUDPAddr(CONN_TYPE, CONN_HOST+":"+port)
	check_error(err)
	udpBroadcastSocket, err := net.DialUDP(CONN_TYPE, nil, udpBroadcastAddress)
	check_error(err)
	Println("BroadFin")
	return udpBroadcastSocket

}

func Udp_receive_msg(udpReadSocket *net.UDPConn) ([]byte, string) {

	msg := make([]byte, 1024)
	length, address, err := udpReadSocket.ReadFromUDP(msg)
	check_error(err)

	return msg[:length], strings.Split(address.String(), ":")[0]

}

func Udp_send_msg(udpBroadcastSocket *net.UDPConn, msg []byte) {

	_, err := udpBroadcastSocket.Write(msg)
	check_error(err)

}

package communication

import (
	. "fmt"
	"net"
	"strings"
)

const (
	CONN_HOST                = "129.241.187.255"
	ELEVATOR_STRUCT_PORT     = "32373"
	ADD_GLOBAL_ORDER_PORT    = "42718"
	REMOVE_GLOBAL_ORDER_PORT = "42714"
	NEW_ELEVATOR_SPAM_PORT   = "36279"
	NEW_ELEVATOR_PORT        = "35826"
	CHECK_NETWORK_CONN_PORT  = "38791"

	CONN_TYPE_UDP = "udp4"
)

func Check_error(err error) {
	if err != nil {
		Println("Error dialing", err.Error())

	}

}

func Set_up_udp_readSocket(port string) *net.UDPConn {
	Println("Setup Readsocket")

	udpReadAddress, err := net.ResolveUDPAddr(CONN_TYPE_UDP, CONN_HOST+":"+port)
	Check_error(err)
	udpReadSocket, err := net.ListenUDP(CONN_TYPE_UDP, udpReadAddress)
	Check_error(err)
	Println("read fin")
	return udpReadSocket

}

func Set_up_udp_sendSocket(port string) *net.UDPConn {
	Println("Setup Broadsocket")

	udpBroadcastAddress, err := net.ResolveUDPAddr(CONN_TYPE_UDP, CONN_HOST+":"+port)
	Check_error(err)
	udpBroadcastSocket, err := net.DialUDP(CONN_TYPE_UDP, nil, udpBroadcastAddress)
	Check_error(err)
	Println("BroadFin")
	return udpBroadcastSocket

}

func Udp_receive_msg(udpReadSocket *net.UDPConn) ([]byte, string) {

	msg := make([]byte, 1024)
	length, address, err := udpReadSocket.ReadFromUDP(msg)
	Check_error(err)

	return msg[:length], strings.Split(address.String(), ":")[0]

}

func Udp_send_msg(udpBroadcastSocket *net.UDPConn, msg []byte) {

	_, err := udpBroadcastSocket.Write(msg)
	Check_error(err)

}

func Get_own_ip() string {

	googleAddress, _ := net.ResolveTCPAddr("tcp", "www.google.com:80")
	googleConn, err := net.DialTCP("tcp", nil, googleAddress)
	for {
		Println("waitin for IP")
		if err == nil {
			break
		}
		googleConn, err = net.DialTCP("tcp", nil, googleAddress)
	}

	Ip := strings.Split(googleConn.LocalAddr().String(), ":")[0]
	googleConn.Close()
	return Ip

}

package UDP

import (
	. "fmt"
	"net"
	"runtime"
	"time"
	//. "strings"
	//. "os"
)

const (
	CONN_HOST      = "129.241.187.255"
	BROADCAST_PORT = "32333"
	READ_PORT      = "32330"
	CONN_TYPE      = "udp4"
)

func check_error(err error) {
	if err != nil {
		Println("Error dialing", err.Error())
	}

}

func set_up_udp_readsocket(udpReadChannel chan *net.UDPConn, port string) {
	Println("Setup Readsocket")

	udpReadAddress, err := net.ResolveUDPAddr(CONN_TYPE, CONN_HOST+":"+port)
	check_error(err)
	udpRead, err := net.ListenUDP(CONN_TYPE, udpReadAddress)
	check_error(err)
	Println("read fin")
	udpReadChannel <- udpRead

}

func set_up_udp_broadcastsocket(udpBroadcastChannel chan *net.UDPConn, port string) {
	Println("Setup Broadsocket")

	udpBroadcastAddress, err := net.ResolveUDPAddr(CONN_TYPE, CONN_HOST+":"+port)
	check_error(err)
	udpBroadcast, err := net.DialUDP(CONN_TYPE, nil, udpBroadcastAddress)
	check_error(err)
	udpBroadcastChannel <- udpBroadcast
	Println("BroadFin")

}

func udp_receive_msg(udpReadChannel chan *net.UDPConn, udpMsgChannel chan []byte) {

	udpTempConn := <-udpReadChannel
	msg := make([]byte, 1024)
	_, _, err := udpTempConn.ReadFromUDP(msg)
	check_error(err)
	Println("melding mottatt:")
	udpReadChannel <- udpTempConn
	select {
	case udpMsgChannel <- msg:
	}

}
func udp_send_msg(udpBroadcastChannel chan *net.UDPConn, msg []byte) {
	udpTempConn := <-udpBroadcastChannel
	_, err := udpTempConn.Write(msg)
	check_error(err)
	udpBroadcastChannel <- udpTempConn

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	udpReadChannel := make(chan *net.UDPConn, 1)
	udpBroadcastChannel := make(chan *net.UDPConn, 1)
	exit := make(chan int)
	//b := []byte("Hva skjer?")
	set_up_udp_readsocket(udpReadChannel)
	set_up_udp_broadcastsocket(udpBroadcastChannel)
	go udp_receive_msg(udpReadChannel)
	go udp_spam_routine(udpBroadcastChannel)
	<-exit

}

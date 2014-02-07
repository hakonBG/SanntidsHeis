package main

import (
	. "fmt"
	. "net"
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

func set_up_udp_readsocket(udpReadChannel chan *UDPConn) {
	Println("Setup Readsocket")

	udpReadAddress, err := ResolveUDPAddr(CONN_TYPE, CONN_HOST+":"+READ_PORT)
	check_error(err)
	udpRead, err := ListenUDP(CONN_TYPE, udpReadAddress)
	check_error(err)
	Println("read fin")
	udpReadChannel <- udpRead

}

func set_up_udp_broadcastsocket(udpBroadcastChannel chan *UDPConn) {
	Println("Setup Broadsocket")

	udpBroadcastAddress, err := ResolveUDPAddr(CONN_TYPE, CONN_HOST+":"+READ_PORT)
	check_error(err)
	udpBroadcast, err := DialUDP(CONN_TYPE, nil, udpBroadcastAddress)
	check_error(err)
	udpBroadcastChannel <- udpBroadcast
	Println("BroadFin")

}

func udp_receive_msg(udpReadChannel chan *UDPConn) {
	for {

		udpTempConn := <-udpReadChannel
		Println("Feil 1")
		msg := make([]byte, 1024)
		Println("Venter pa melding")
		_, _, err := udpTempConn.ReadFromUDP(msg)
		check_error(err)

		Println("melding mottatt:")
		Printf("%s \n", msg)
		udpReadChannel <- udpTempConn
	}

}
func udp_send_msg(udpBroadcastChannel chan *UDPConn) {
	udpTempConn := <-udpBroadcastChannel
	Println("UDPbroadcast, " + udpTempConn.RemoteAddr().String())
	msg := []byte("Fy faen det funker\x00")
	_, err := udpTempConn.Write(msg)
	check_error(err)
	Println("Melding sendt")
	udpBroadcastChannel <- udpTempConn

}

func udp_spam_routine(udpBroadcastChannel chan *UDPConn) {

	time1 := time.Now()
	for {
		time2 := time.Now()
		if time2.Second()-time1.Second() >= 1 {
			udp_send_msg(udpBroadcastChannel)
			time1 = time.Now()
		}
	}

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	udpReadChannel := make(chan *UDPConn, 1)
	udpBroadcastChannel := make(chan *UDPConn, 1)
	exit := make(chan int)
	//b := []byte("Hva skjer?")
	set_up_udp_readsocket(udpReadChannel)
	set_up_udp_broadcastsocket(udpBroadcastChannel)
	go udp_receive_msg(udpReadChannel)
	go udp_spam_routine(udpBroadcastChannel)
	<-exit

}

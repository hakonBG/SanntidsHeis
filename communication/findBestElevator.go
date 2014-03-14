package communication

import (
	"./../driver"
	"./../ownVar"
	//"fmt"
	"net"
	"strings"
)

const (
	N_FLOORS            = ownVar.N_FLOORS
	N_BUTTONTYPES       = ownVar.N_BUTTONTYPES
	N_GLOBALBUTTONTYPES = ownVar.N_GLOBALBUTTONTYPES
)

func Find_best_elevator(
	passOrders chan chan ownVar.Orders_s,
	passElevators chan chan map[string]ownVar.Elevator_s,
	addOrderChan chan ownVar.Order_call_s,
	removeOrderChan chan ownVar.Order_call_s,
	assignedOrdersChan chan [N_GLOBALBUTTONTYPES][N_FLOORS]string) {
	//Start of function

	var assignedOrders [N_GLOBALBUTTONTYPES][N_FLOORS]string
	var orders ownVar.Orders_s
	var bestElevator ownVar.Elevator_s
	orderChan := make(chan ownVar.Orders_s)
	elevatorChan := make(chan map[string]ownVar.Elevator_s)
	elevators := make(map[string]ownVar.Elevator_s)
	var orderCall ownVar.Order_call_s

	orderCall.OrderType = ownVar.GLOBAL

	ownIp := get_own_ip()
	for {
		select {
		case passElevators <- elevatorChan:
			elevators = <-elevatorChan
		case passOrders <- orderChan:
			orders = <-orderChan

		default:

			for i := 0; i < N_GLOBALBUTTONTYPES; i++ {

				if i == BUTTON_CALL_UP {
					orderCall.ButtonType = driver.BUTTON_CALL_UP
				} else if i == BUTTON_CALL_DOWN {
					orderCall.ButtonType = driver.BUTTON_CALL_DOWN
				}
				for j := 0; j < N_FLOORS; j++ {

					if orders.GlobalOrders[i][j] == 1 {

						orderCall.Floor = j
						orderCall.OrderType = ownVar.GLOBAL
						bestElevator = find_min_elev(elevators, orderCall)
						assignedOrders[i][j] = bestElevator.Ip
						if (bestElevator.Ip == ownIp) && (orders.LocalOrders[i][j] == 0) {
							orderCall.OrderType = ownVar.LOCAL
							addOrderChan <- orderCall
						} else if (bestElevator.Ip != ownIp) && (orders.LocalOrders[i][j] == 1) {
							orderCall.OrderType = ownVar.LOCAL
							removeOrderChan <- orderCall

						}
					} else {
						assignedOrders[i][j] = "tom"
					}

				}

			}
			assignedOrdersChan <- assignedOrders

		}

	}

}

func get_own_ip() string {
	googleAddress, _ := net.ResolveTCPAddr("tcp", "www.google.com:80")
	googleConn, _ := net.DialTCP("tcp", nil, googleAddress)
	Ip := strings.Split(googleConn.LocalAddr().String(), ":")[0]
	googleConn.Close()
	return Ip

}

func elev_cost(elev ownVar.Elevator_s, order ownVar.Order_call_s) int {
	value := 0
	if order.ButtonType == driver.BUTTON_CALL_UP {
		if elev.Direction == ownVar.UP {
			if !elev.Moving && elev.CurrentFloor == order.Floor {
				return 0

			} else if elev.CurrentFloor < order.Floor {
				for i := elev.CurrentFloor; i < order.Floor; i++ {
					value++
					if elev.Orders[driver.BUTTON_COMMAND][i] == 1 {
						value++

					}
				}
			} else {
				value = (N_FLOORS-1)*2 - (elev.CurrentFloor - order.Floor)
				for i := 0; i < N_FLOORS; i++ {
					if elev.Orders[BUTTON_COMMAND][i] == 1 {
						value++
					}
				}
			}

		}
	} else if order.ButtonType == driver.BUTTON_CALL_DOWN {
		if elev.Direction == ownVar.DOWN {
			if !elev.Moving && elev.CurrentFloor == order.Floor {
				return 0
			} else if elev.CurrentFloor > order.Floor {
				for i := elev.CurrentFloor; i > order.Floor; i-- {
					value++
					if elev.Orders[driver.BUTTON_COMMAND][i] == 1 {
						value++
					}

				}
			} else {
				value = (N_FLOORS-1)*2 - (order.Floor - elev.CurrentFloor)
				for i := 0; i < N_FLOORS; i++ {
					if elev.Orders[BUTTON_COMMAND][i] == 1 {
						value++
					}
				}
			}
		}
	}
	return value

}

func find_min_elev(elevators map[string]ownVar.Elevator_s, order ownVar.Order_call_s) ownVar.Elevator_s {
	minCost := 10000
	var elevCost int
	var minElev ownVar.Elevator_s

	for _, elev := range elevators {

		elevCost = elev_cost(elev, order)
		if elevCost < minCost {
			minCost = elevCost
			minElev = elev
		} else if elevCost == minCost {
			if elev.Ip < minElev.Ip {
				minElev = elev
			}
		}
	}
	return minElev
}

package communication

import (
	"./../driver"
	"./../ownVar"
	"fmt"
	"time"
)

const (
	QUARTER_SECOND = 250000000 //250 000 000
)

func Merge_global_orders(
	passOrders chan chan ownVar.Orders_s,
	passElevators chan chan map[string]ownVar.Elevator_s,
	assignedOrdersChan chan [N_GLOBALBUTTONTYPES][N_FLOORS]string,
	removeOrderChan chan ownVar.Order_call_s,
	addOrderChan chan ownVar.Order_call_s) {

	//Start of function

	//Channels
	elevators := make(map[string]ownVar.Elevator_s)
	passOrdersChan := make(chan ownVar.Orders_s)
	passElevatorsChan := make(chan map[string]ownVar.Elevator_s)

	//Variables
	var orders ownVar.Orders_s
	var assignedOrders [N_GLOBALBUTTONTYPES][N_FLOORS]string
	var orderCall ownVar.Order_call_s
	orderCall.OrderType = ownVar.GLOBAL
	timer := time.Now()
	ownIp := Get_own_ip()

	//Do
	for {
		select {
		case passElevators <- passElevatorsChan:
			elevators = <-passElevatorsChan
		case passOrders <- passOrdersChan:
			orders = <-passOrdersChan

		case assignedOrders = <-assignedOrdersChan:
		default:

		}

		//Do every quarter second

		if time.Now().Sub(timer) > 1000*time.Millisecond {

			for i := 0; i < N_GLOBALBUTTONTYPES; i++ {
				if i == BUTTON_CALL_UP {
					orderCall.ButtonType = driver.BUTTON_CALL_UP
				} else if i == BUTTON_CALL_DOWN {
					orderCall.ButtonType = driver.BUTTON_CALL_DOWN
				}
				for j := 0; j < N_FLOORS; j++ {
					orderCall.Floor = j
					for _, elev := range elevators {

						if (orders.GlobalOrders[i][j] == 0) && (elev.Orders[i][j] == 1) && (elev.Ip != ownIp) {
							fmt.Println("Merge Add")
							orders.GlobalOrders[i][j] = 1
							addOrderChan <- orderCall

						}
					}
				}
			}

			for i := 0; i < N_GLOBALBUTTONTYPES; i++ {
				if i == BUTTON_CALL_UP {
					orderCall.ButtonType = driver.BUTTON_CALL_UP
				} else if i == BUTTON_CALL_DOWN {
					orderCall.ButtonType = driver.BUTTON_CALL_DOWN
				}
				for j := 0; j < N_FLOORS; j++ {
					orderCall.Floor = j
					elevIp := assignedOrders[i][j]
					if (elevIp != "tom") && (elevIp != ownIp) {
						elev := elevators[elevIp]
						if elev.Orders[i][j] == 0 {
							assignedOrders[i][j] = "tom"
							orders.GlobalOrders[i][j] = 0

							removeOrderChan <- orderCall

							fmt.Println("merge remove")

						}
					}

				}
			}

			timer = time.Now()

		}
	}
}

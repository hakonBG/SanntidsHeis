package communication

const (
	N_FLOORS      = 4
	N_BUTTONTYPES = 3
)

type Order_call_s struct {
	orderType  order_type_t
	buttonType driver.Elev_button_type_t
	floor      int
}

func Find_best_elevator() {

}

func costFunction(elevators map[string]Elevator_s, orderCall logic.Order_call_s) string {
	n := len(elevators)
	c := 10000
	var ip string
	var c_temp int

	for i := 0; i < n; i++ {
		if elevators[i].Status == IDLE && elevators[i].NextFloor == -1 {
			c_temp = Abs(elevators[i].CurrentFloor - orderCall.floor)
			if c_temp <= c {
				c = c_temp
				ip = elevators[i].Ip
			}
		}
		if elevators[i].Status == MOVING_UP && elevators[i].CurrentFloor < orderCall.floor && orderCall.floor <= elevators[i].NextFloor && orderCall.buttonType == driver.BUTTON_CALL_UP {
			c_temp = Abs(elevators[i].CurrentFloor - orderCall.floor)
			if c_temp < c {
				c = c_temp
				ip = elevators[i].Ip
			}
		}
		if elevators[i].Status == MOVING_DOWN && elevators[i].NextFloor <= orderCall.floor && orderCall.floor < elevators[i].CurrentFloor && orderCall.buttonType == driver.BUTTON_CALL_DOWN {
			c_temp = elevators[i].CurrentFloor - orderCall.floor
			if c_temp < c {
				c = c_temp
				ip = elevators[i].Ip
			}
		}
	}

	if c != 10000 {
		return ip
	}
}

func elev_cost(elev map[string]Elevator_s, order logic.Order_call_s) int {
	value := 0
	if order.buttonType == driver.BUTTON_CALL_UP {
		if elev.Direction == logic.UP {
			if !elev.Moving && elev.CurrentFloor == order.floor {
				return 0

			} else if elev.CurrentFloor < order.floor {
				for i := elev.CurrentFloor; i < order.floor; i++ {
					value++
					if elev.Orders[driver.BUTTON_CALL_UP][i] == 1 || elev.Orders[driver.BUTTON_COMMAND][i] {
						value++

					}
				}
			} else {
				value = (N_FLOORS-1)*2 - (elev.CurrentFloor - order.floor)
				for i := 0; i < N_FLOORS; i++ {
					if (elev.Orders[BUTTON_CALL_UP][i] == 1) || (elev.Orders[BUTTON_COMMAND] == 1) {
						value++
					}
					if elev.Orders[BUTTON_CALL_DOWN][i] == 1 {
						value++
					}
				}
			}

		}
	} else if order.buttonType == driver.BUTTON_CALL_DOWN {
		if elev.Direction == logic.DOWN {
			if !elev.Moving && elev.CurrentFloor == order.floor {
				return 0
			} else if elev.CurrentFloor > order.floor {
				for i := elev.CurrentFloor; i > order.floor; i-- {
					value++
					if elev.Orders[driver.BUTTON_CALL_DOWN][i] == 1 || elev.Orders[driver.BUTTON_COMMAND][i] {
						value++
					}
					if elev.Orders[BUTTON_CALL_UP][i] == 1 {
						value++
					}
				}
			}
		}
	}

}

func min_cost(elev []Elevator_s, order logic.Order_call_s) Elevator_s {
	minCost := 10000
	var elevCost int
	var minElev Elevator_s
	for i := 0; i < len(elev); i++ {
		elevCost = elev_cost(elev[i], order)
		if elevCost <= minCost {
			minCost = elevCost
			minElev = elev[i]
		}
	}
	return minElev
}

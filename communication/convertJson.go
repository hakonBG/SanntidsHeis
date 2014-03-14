package communication

import ()

func Json_encode_elevator(elev Elevator_s) []byte {
	elevS, _ := json.Marshal(elev)
	return []byte(elevS)
}

func Json_decode_elevator(elevB []byte) Elevator_s {
	elev := Elevator_s{}
	json.Unmarshal(elevB, &elev)
	return elev
}

func Json_encode_order(order logic.OrderCall_s) []byte {
	orderB, _ := json.Marshal(order)
	return [1024]byte(orderB)
}

func Json_decode_order(orderB [1024]byte) logic.OrderCall_s {
	order := logic.OrderCall_s{}
	json.Unmarshal(orderB, &order)
	return order
}

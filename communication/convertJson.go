package communication

import (
	"./../ownVar"
	"encoding/json"
)

func Json_encode_elevator(elev ownVar.Elevator_s) []byte {
	elevS, _ := json.Marshal(elev)
	return []byte(elevS)
}

func Json_decode_elevator(elevB []byte) ownVar.Elevator_s {
	elev := ownVar.Elevator_s{}
	json.Unmarshal(elevB, &elev)
	return elev
}

func Json_encode_order(order ownVar.Order_call_s) []byte {
	orderB, _ := json.Marshal(order)
	return []byte(orderB)
}

func Json_decode_order(orderB []byte) ownVar.Order_call_s {
	order := ownVar.Order_call_s{}
	json.Unmarshal(orderB, &order)
	return order
}


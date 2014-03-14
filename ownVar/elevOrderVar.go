package ownVar

import (
	"./../driver"
)

const (
	N_FLOORS            = 4
	N_BUTTONTYPES       = 3
	N_GLOBALBUTTONTYPES = 2
	LAMP_ON             = 1
	LAMP_OFF            = 0
)

type Order_type_t int

const (
	GLOBAL Order_type_t = iota
	LOCAL
)

type Order_call_s struct {
	OrderType  Order_type_t
	ButtonType driver.Elev_button_type_t
	Floor      int
}

type Orders_s struct {
	GlobalOrders [N_GLOBALBUTTONTYPES][N_FLOORS]int
	LocalOrders  [N_BUTTONTYPES][N_FLOORS]int
}

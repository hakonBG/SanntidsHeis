package ownVar

import (
	"time"
)

type Direction_t int

const (
	UP Direction_t = iota
	DOWN
	IDLE
)

type Elevator_s struct {
	Direction    Direction_t
	Moving       bool
	Orders       [N_BUTTONTYPES][N_FLOORS]int
	LastTime     time.Time
	Ip           string
	NextFloor    int
	CurrentFloor int
}

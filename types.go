package main

// Property represents different entity properties
type Property int

const (
	ADDER Property = iota
	DELETER
	MOVABLE
	OUTPUT_LEFT
	OUTPUT_RIGHT
	OUTPUT_TOP
	OUTPUT_BOTTOM
)

type EntityID string

type Move struct {
	EntityID EntityID
	Coord    Coord
}

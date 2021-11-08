package main

import (
	"fmt"
)

type elevator struct {
	Direction     int
	Floor         int
	Capacity      int
	Contents      []int
	History       []int
	MaxFloorInDir int
}

type building struct {
	El             elevator
	Queues         [][]int
	Floors         int
	PeopleInQueues int
}

func NewBuilding(q [][]int, cap int) building {
	bldg := building{}
	bldg.Queues = q
	el := elevator{}
	//elevator always starts at 0
	el.History = []int{0}
	el.Capacity = cap
	//elevator will always start going up
	el.Direction = 1
	//assume queue will have someone to pickup at or above floor 0
	el.MaxFloorInDir = 0
	bldg.El = el
	bldg.Floors = len(q)
	bldg.PeopleInQueues = 0
	for i := 0; i < len(q); i++ {
		if len(q[i]) > 0 {
			bldg.PeopleInQueues += len(q[i])
		}
	}
	bldg.getMaxFloor()
	return bldg
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func (b *building) getMaxFloor() {
	for i := b.El.Floor + b.El.Direction; i < b.Floors && i >= 0; i += b.El.Direction {
		if len(b.Queues[i]) > 0 {
			b.El.MaxFloorInDir = i
		}
	}
}

func (b *building) moveElevator() {
	stopAtFloor := false
	elPop := len(b.El.Contents)
	qPop := len(b.Queues[b.El.Floor])
	//Let people off from contents
	if elPop > 0 {
		for i := 0; i < elPop; i++ {
			if b.El.Contents[i] == b.El.Floor {
				b.El.Contents = append(b.El.Contents[:i], b.El.Contents[i+1:]...)
				i--
				elPop--
				stopAtFloor = true
			}
		}
	}
	//Let people on who are going in the same direction
	if qPop > 0 {
		for i := 0; i < qPop; i++ {
			if (b.Queues[b.El.Floor][i]-b.El.Floor)*b.El.Direction > 0 {
				stopAtFloor = true
				if elPop < b.El.Capacity {
					//add to elevator
					b.El.Contents = append(b.El.Contents, b.Queues[b.El.Floor][i])
					elPop++
					//remove from queue
					b.Queues[b.El.Floor] = append(b.Queues[b.El.Floor][:i], b.Queues[b.El.Floor][i+1:]...)
					b.PeopleInQueues--
					qPop--
					i--
				}
			}
		}
	}
	// if anyone got on or off, add to path
	if stopAtFloor && b.El.Floor != b.El.History[len(b.El.History)-1] {
		b.El.History = append(b.El.History, b.El.Floor)
	}
	// if difference between floor and maxfloor times direction >= 0 then not yet at maxfloor
	// if at or beyond maxfloor and elevator empty, switch directions
	if (b.El.Floor-b.El.MaxFloorInDir)*b.El.Direction >= 0 && elPop == 0 {
		b.El.Floor += b.El.Direction
		b.El.Direction *= -1
		//recalculate max floor with new direction
		b.getMaxFloor()
	}
	// go to next floor in direction
	b.El.Floor = b.El.Floor + b.El.Direction
	// if more people need to be picked up or people remain on elevator, recurse
	if b.PeopleInQueues > 0 || elPop > 0 {
		b.moveElevator()
	} else {
		// if elevator both empty and queues empty
		// check if the elevator is already at floor 0. If so, do nothing, if not, go to 0
		if b.El.History[len(b.El.History)-1] != 0 {
			b.El.History = append(b.El.History, 0)
		}
	}
}

func main() {
	q := [][]int{
		{},     // G
		{},     // 1
		{1, 1}, // 2
		{},     // 3
		{},     // 4
		{},     // 5
		{},     // 6
	}
	bldg := NewBuilding(q, 5)
	bldg.moveElevator()
	fmt.Println(bldg.El.History)
}

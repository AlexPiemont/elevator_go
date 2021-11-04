package main

import (
	"fmt"
	"math"
)

type elevator struct {
	Direction int
	Floor     int
	Capacity  int
	Contents  []int
	History   []int
}

type building struct {
	El     elevator
	Queues [][]int
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
	bldg.El = el
	return bldg
}

func (b *building) moveElevator() {
	stopAtFloor := false
	//Let people off from contents
	if len(b.El.Contents) > 0 {
		for i := 0; i < len(b.El.Contents); i++ {
			if b.El.Contents[i] == b.El.Floor {
				b.El.Contents = append(b.El.Contents[:i], b.El.Contents[i+1:]...)
				i--
				stopAtFloor = true
			}
		}
	}
	//Let people on who are going in the same direction
	if len(b.Queues[b.El.Floor]) > 0 {
		for i := 0; i < len(b.Queues[b.El.Floor]); i++ {
			//if the difference between desired and actual floor * direction is positive,
			//then the person is going the same direction
			//e.g. person going to 4 from 3 difference is 1. If direction is 1 then result is 1, positive
			// person going to 3 from 4 difference is -1. if direction is -1 (down) then result is 1, positive
			// mismatch in sign between difference and direction means negative result means not going same dir
			if (b.Queues[b.El.Floor][i]-b.El.Floor)*b.El.Direction > 0 {
				if len(b.El.Contents) < b.El.Capacity {
					stopAtFloor = true
					//add to elevator
					b.El.Contents = append(b.El.Contents, b.Queues[b.El.Floor][i])
					//remove from queue
					b.Queues[b.El.Floor] = append(b.Queues[b.El.Floor][:i], b.Queues[b.El.Floor][i+1:]...)
					i--
				}
			}
		}
	}
	// if anyone got on or off, add to path
	if stopAtFloor {
		b.El.History = append(b.El.History, b.El.Floor)
	}
	//Check floors in direction for people
	//using two flags here allows to only iterate through whole queue once per elevator iteration
	//one checks if there are more people waiting at floors in the current direction
	//one checks if the queues are empty entirely, and ready for the elevator to stop
	peopleFoundInDir := false
	peopleFoundInQueue := false
	for i := 0; i < len(b.Queues); i++ {
		if len(b.Queues[i]) > 0 {
			peopleFoundInQueue = true
			//if (b.El.Direction < 0 && i < b.El.Floor + b.El.Direction) || (b.El.Direction > 0 && i > b.El.Floor + b.El.Direction){
			// if the floor with people (i) is in the direction of the elevator from El.Floor
			// if the abs of difference between the current floor and the found floor is greater than the difference between the next floor and found floor
			// if on 4 going down and found on 2 then 2 > 1
			// if on 2 going up and found on 6 then -4 > -3
			if math.Abs(float64(b.El.Floor-i)) > math.Abs(float64(b.El.Floor+b.El.Direction-i)) {
				peopleFoundInDir = true
				break
			}
		}
	}
	if !peopleFoundInDir && len(b.El.Contents) == 0 {
		//if no people in direction && elevator empty switch direction
		//before switching direction, increment floor in direction so that iteration
		// doesn't skip current floor
		// Essentially, if this is not done, the elevator might stop at a floor with people
		// who are headed in the wrong direction and therefore not pick them up, then switch direction,
		// then go to the opposite direction without picking up these people. This line allows
		// the elevator to redo the drop-off/pick-up logic at the current floor with the new direction
		b.El.Floor += b.El.Direction
		b.El.Direction *= -1
	}
	// go to next floor in direction
	b.El.Floor = b.El.Floor + b.El.Direction
	// if more people need to be picked up or people remain on elevator, recurse
	if peopleFoundInQueue || len(b.El.Contents) > 0 {
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

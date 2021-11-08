package main

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
	el.History = []int{0}
	el.Capacity = cap
	el.Direction = 1
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

func (b *building) getMaxFloor() {
	for i := b.El.Floor + b.El.Direction; i < b.Floors && i >= 0; i += b.El.Direction {
		if len(b.Queues[i]) > 0 {
			b.El.MaxFloorInDir = i
		}
	}
}

func (e *elevator) registerStop() {
	if e.Floor != e.History[len(e.History)-1] {
		e.History = append(e.History, e.Floor)
	}
}

func (e *elevator) unload() {
	for i := 0; i < len(e.Contents); i++ {
		if e.Contents[i] == e.Floor {
			e.Contents = append(e.Contents[:i], e.Contents[i+1:]...)
			i--
			e.registerStop()
		}
	}
}

func (b *building) loadElevator() {
	for i := 0; i < len(b.Queues[b.El.Floor]); i++ {
		if (b.Queues[b.El.Floor][i]-b.El.Floor)*b.El.Direction > 0 {
			b.El.registerStop()
			if len(b.El.Contents) < b.El.Capacity {
				//add to elevator
				b.El.Contents = append(b.El.Contents, b.Queues[b.El.Floor][i])
				//remove from queue
				b.Queues[b.El.Floor] = append(b.Queues[b.El.Floor][:i], b.Queues[b.El.Floor][i+1:]...)
				b.PeopleInQueues--
				i--
			}
		}
	}
}
func (e *elevator) needSwitch() bool {
	return (e.Floor-e.MaxFloorInDir)*e.Direction >= 0 && len(e.Contents) == 0
}

func (b *building) moveElevator() {
	b.El.unload()
	b.loadElevator()
	if b.El.needSwitch() {
		b.El.Floor += b.El.Direction
		b.El.Direction *= -1
		//recalculate max floor with new direction
		b.getMaxFloor()
	}
	b.El.Floor = b.El.Floor + b.El.Direction
	if b.PeopleInQueues > 0 || len(b.El.Contents) > 0 {
		b.moveElevator()
	} else {
		b.El.Floor = 0
		b.El.registerStop()
	}

}

func main() {

}

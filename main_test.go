package main

import (
	"reflect"
	"testing"
)

func TestUp(t *testing.T) {
	queues := [][]int{
		{},        // G
		{},        // 1
		{5, 5, 5}, // 2
		{},        // 3
		{},        // 4
		{},        // 5
		{},        // 6
	}
	bldg := NewBuilding(queues, 5)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 2, 5, 0}) {
		t.Fatalf(`Expectated: {0, 2, 5, 0} Got: %v`, bldg.El.History)
	}
}
func TestDown(t *testing.T) {
	queues := [][]int{
		{},     // G
		{},     // 1
		{1, 1}, // 2
		{},     // 3
		{},     // 4
		{},     // 5
		{},     // 6
	}
	bldg := NewBuilding(queues, 5)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 2, 1, 0}) {
		t.Fatalf(`Expectated: {0, 2, 1, 0} Got: %v`, bldg.El.History)
	}
}
func TestUpAndUp(t *testing.T) {
	queues := [][]int{
		{},  // G
		{3}, // 1
		{4}, // 2
		{},  // 3
		{5}, // 4
		{},  // 5
		{},  // 6
	}
	bldg := NewBuilding(queues, 5)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 1, 2, 3, 4, 5, 0}) {
		t.Fatalf(`Expectated: {0, 1, 2, 3, 4, 5, 0} Got: %v`, bldg.El.History)
	}
}
func TestDownAndDown(t *testing.T) {
	queues := [][]int{
		{},  // G
		{0}, // 1
		{},  // 2
		{},  // 3
		{2}, // 4
		{3}, // 5
		{},  // 6
	}
	bldg := NewBuilding(queues, 5)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 5, 4, 3, 2, 1, 0}) {
		t.Fatalf(`Expectated: {0, 5, 4, 3, 2, 1, 0}, Got: %v`, bldg.El.History)
	}
}
func TestDownAndDownWithCap(t *testing.T) {
	queues := [][]int{
		{},        // G
		{0},       // 1
		{},        // 2
		{},        // 3
		{2, 2, 2}, // 4
		{3, 3},    // 5
		{},        // 6
	}
	bldg := NewBuilding(queues, 3)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 5, 4, 3, 2, 1, 0, 4, 2, 0}) {
		t.Fatalf(`Expectated: {0, 5, 4, 3, 2, 1, 0, 4, 2, 0}, Got: %v`, bldg.El.History)
	}
}
func TestIrina(t *testing.T) {
	queues := [][]int{
		{3, 4, 5}, // G
		{},        // 1
		{4},       // 2
		{0},       // 3
		{5},       // 4
		{1, 0, 0}, // 5
	}
	bldg := NewBuilding(queues, 2)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 2, 3, 4, 5, 3, 1, 0, 2, 4, 5, 3, 0}) {
		t.Fatalf(`Expectated:{0, 2, 3, 4, 5, 3, 1, 0, 2, 4, 5, 3, 0}, Got: %v`, bldg.El.History)
	}
}

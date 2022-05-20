package astar

import "testing"

func TestPathFinding1(t *testing.T) {
	nodes := CreateNodes(5, 5)

	start := nodes.Access(0, 0)
	goal := nodes.Access(4, 4)

	path := nodes.FindPath(start, goal)

	if len(path) != 5 {
		t.Error(path)
		t.Fail()
	}
}

func TestPathFinding2(t *testing.T) {
	nodes := CreateNodes(5, 5)

	start := nodes.Access(0, 0)
	goal := nodes.Access(0, 2)

	path := nodes.FindPath(start, goal)

	if len(path) != 3 {
		t.Error(path)
		t.Fail()
	}
}

func TestMapLoading(t *testing.T) {
	const mapStr = "....." +
		"....." +
		"....." +
		"....." +
		"....."

	nodes, err := ReadMap(mapStr)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	start := nodes.Access(0, 0)
	goal := nodes.Access(4, 4)

	path := nodes.FindPath(start, goal)

	if len(path) != 5 {
		t.Error(path)
		t.Fail()
	}
}

func TestObstacleInMap(t *testing.T) {
	const mapStr = "....." +
		"....." +
		"..#.." +
		"....." +
		"....."

	nodes, err := ReadMap(mapStr)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	start := nodes.Access(0, 0)
	goal := nodes.Access(4, 4)

	path := nodes.FindPath(start, goal)

	if len(path) != 6 {
		t.Error(path)
		t.Fail()
	}
}

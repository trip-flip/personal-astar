package astar

import "testing"

func TestPathFinding(t *testing.T) {
	nodes := CreateNodes(5, 5)

	start := nodes.Access(0, 0)
	goal := nodes.Access(4, 4)

	path := nodes.FindPath(start, goal)

	if len(path) != 5 {
		t.Error(path)
		t.Fail()
	}
}

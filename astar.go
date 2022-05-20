package astar

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

const (
	Walkable = '.'
	Obstacle = '#'
)

type Node struct {
	parent   *Node
	nodeType byte
	x        uint32
	y        uint32
	f        float64
	g        float64
	h        float64
}
type Nodes struct {
	nodes []Node
	xSize uint32
	ySize uint32
}
type NodeRefs []*Node
type Path []*Node

func CreateNodes(x, y uint32) Nodes {
	nodes := Nodes{
		nodes: make([]Node, x*y),
		xSize: x,
		ySize: y,
	}
	nArray := nodes.nodes
	x, y = 0, 0
	for i := range nArray {
		node := &nArray[i]
		node.x = x
		node.y = y
		node.nodeType = '.'

		if x == nodes.xSize-1 {
			x = 0
			y += 1
		} else {
			x += 1
		}
	}

	return nodes
}

// Only reads square maps currently
func ReadMap(m string) (Nodes, error) {
	nodes := Nodes{}
	mapLen := len(m)
	s := int(math.Sqrt(float64(mapLen)))
	if s*s != mapLen {
		return nodes, errors.New("provided map string is not a square grid")
	}

	size := uint32(s)
	nodes.nodes = make([]Node, mapLen)
	nodes.xSize = size
	nodes.ySize = size

	nArray := nodes.nodes
	var x, y uint32 = 0, 0
	for i := range nArray {
		node := &nArray[i]
		node.x = x
		node.y = y
		node.nodeType = m[i]

		if x == nodes.xSize-1 {
			x = 0
			y += 1
		} else {
			x += 1
		}
	}

	return nodes, nil
}

func (a NodeRefs) Len() int           { return len(a) }
func (a NodeRefs) Less(i, j int) bool { return a[i].f < a[j].f }
func (a NodeRefs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Returns a slice of Node addresses. If the slice is nil, path was not found.
func (nodes Nodes) FindPath(start *Node, goal *Node) Path {
	openNodes := make([]*Node, 0, 5)
	closedNodes := make(map[*Node]bool)
	openNodes = append(openNodes, start)

	for pathSize := 1; len(openNodes) != 0; pathSize++ {
		currentNode := openNodes[0]
		closedNodes[currentNode] = true
		openNodes = append(openNodes[:0], openNodes[1:]...) // Remove first element

		if currentNode.x == goal.x && currentNode.y == goal.y {
			path := make([]*Node, 0, pathSize)
			n := currentNode
			for n != nil {
				path = append(path, n)
				n = n.parent
			}
			return path
		}
		// For n in neighbors
		for _, n := range nodes.neighbors(currentNode.x, currentNode.y) {
			_, isClosed := closedNodes[n]

			if isClosed || n == nil || n.nodeType == Obstacle {
				if n != nil {
					fmt.Println(n.nodeType)
				}
				continue
			}

			n.parent = currentNode
			n.g = currentNode.g + distance(currentNode, n)
			n.h = distance(n, goal)
			n.f = n.g + n.h
			openNodes = append(openNodes, n)
		}

		sort.Sort(NodeRefs(openNodes))
	}

	return nil
}

func (nodes Nodes) Access(x, y uint32) *Node {
	i := y*nodes.xSize + x
	if int(i) < len(nodes.nodes) {
		return &nodes.nodes[y*nodes.xSize+x]
	} else {
		return nil
	}
}

func (nodes Nodes) neighbors(x, y uint32) [8]*Node {
	// In no particular order
	nodeRefs := [8]*Node{
		nodes.Access(x-1, y-1),
		nodes.Access(x-1, y+0),
		nodes.Access(x-1, y+1),
		nodes.Access(x+0, y-1),
		nodes.Access(x+0, y+1),
		nodes.Access(x+1, y-1),
		nodes.Access(x+1, y+0),
		nodes.Access(x+1, y+1),
	}
	return nodeRefs
}

func distance(n1 *Node, n2 *Node) float64 {
	xPow := math.Pow(float64(n2.x-n1.x), 2)
	yPow := math.Pow(float64(n2.y-n1.y), 2)
	return math.Sqrt(
		xPow + yPow,
	)
}

func (p Path) String() (str string) {
	pLen := len(p)
	for i := range p {
		n := p[i]
		str += fmt.Sprintf("[%v %v]", n.x, n.y)
		if i != pLen-1 {
			str += "<--"
		}
	}

	return str
}

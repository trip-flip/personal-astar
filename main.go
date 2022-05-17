package main

import "fmt"

func main() {
	nodes := CreateNodes(5, 5)

	start := nodes.Access(0, 0)
	goal := nodes.Access(4, 4)

	if start == nil || goal == nil {
		fmt.Println("Start or goal null in main")
	}

	path := nodes.FindPath(start, goal)
	fmt.Println(path)
}

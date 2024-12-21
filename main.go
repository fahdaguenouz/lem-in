package main

import (
	"fmt"
	"log"

	Mosdef "Mosdef/funcs"
)

func main() {
	Mosdef.CheckArgs()
	lines := Mosdef.ReadFile()
	start, end, antsNumber, graph := Mosdef.GetRooms(lines)
	var allPaths [][]string
	if len(graph) > 50 {
		allPaths = Mosdef.BreadthFirstSearch(graph, start, end)
	} else {
		allPaths = Mosdef.DepthFirstSearch(graph, start, end)
	}
	filteredPaths := Mosdef.FilterPaths(allPaths)
	fmt.Println(filteredPaths)
	antDistribution := Mosdef.DistributeAnts(filteredPaths, antsNumber)
	finalResult, moveCount := Mosdef.SimulateAntMovement(filteredPaths, antDistribution)

	if moveCount < 1 {
		log.Fatal("Error: Invalid data format")
	}
	fmt.Println(finalResult)
}

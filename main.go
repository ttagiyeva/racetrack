package main

import "fmt"

func main() {
	inputs := readFromConsole()
	for _, input := range inputs {
		vertex := NewVertex(0, input.start, Velocity{DX: 0, DY: 0}, nil)
		obstacles := CreateObstacles(input.obstacleCount, input.obstacles)
		result := vertex.BFS(input.finish, input.width, input.height, obstacles)

		if result != nil {
			fmt.Printf("Optimal solution takes %d hops.\n", result.value)
			HopperJourney(result, input.start, input.finish, input.width, input.height, input.obstacles)
		} else {
			fmt.Println("No solution found.")
		}
	}
}

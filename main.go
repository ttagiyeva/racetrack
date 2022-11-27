package main

import "fmt"

func main() {
	inputs := readFromConsole()
	for _, input := range inputs {
		vertex := NewVertex(0, input.start, Velocity{DX: 0, DY: 0})
		obstacles := CreateObstacles(input.obstacleCount, input.obstacles)
		hops := vertex.BFS(input.finish, input.width, input.height, obstacles)

		if hops > 0 {
			fmt.Printf("Optimal solution takes %d hops.\n", hops)
		} else {
			fmt.Println("No solution found.")
		}
	}
}

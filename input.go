package main

import "fmt"

// Input saves inputes added from the console
type Input struct {
	circleCount   int
	width         int
	height        int
	start         Point
	finish        Point
	obstacleCount int
	obstacles     [][2][2]int
}

// readFromConsole reads inputs from the console
// lines indicate:
// - the number of circles
// - the width and height of the grid
// - the start and finish points
// - the number of obstacles
// - the obstacles themselves
//   - start and finish points of the obstacle
func readFromConsole() []Input {
	inputs := make([]Input, 0)
	var input Input
	fmt.Scan(&input.circleCount)

	for i := 0; i < input.circleCount; i++ {
		fmt.Scan(&input.width, &input.height)
		fmt.Scan(&input.start.x, &input.start.y, &input.finish.x, &input.finish.y)
		fmt.Scan(&input.obstacleCount)

		input.obstacles = make([][2][2]int, input.obstacleCount)
		for i := 0; i < input.obstacleCount; i++ {
			fmt.Scan(&input.obstacles[i][0][0], &input.obstacles[i][0][1], &input.obstacles[i][1][0], &input.obstacles[i][1][1])
		}
		inputs = append(inputs, input)
	}

	return inputs
}

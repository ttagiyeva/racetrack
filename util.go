package main

import (
	"fmt"
	"strings"
)

type Grid [][]int

var grid Grid
var directions = map[string]string{
	"up":        "↑",
	"upright":   "↗",
	"right":     "→",
	"downright": "↘",
	"down":      "↓",
	"downleft":  "↙",
	"left":      "←",
	"upleft":    "↖",
	"":          "↻",
}

func InitGrid(width, height int, obstaclesCoors [][2][2]int) {
	grid = make(Grid, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	for k := 0; k < len(obstaclesCoors); k++ {
		for i := obstaclesCoors[k][0][0]; i <= obstaclesCoors[k][0][1]; i++ {
			for j := obstaclesCoors[k][1][0]; j <= obstaclesCoors[k][1][1]; j++ {
				grid[j][i] = 1
			}
		}
	}
}

func HopperJourney(h *Vertex, start, finish Point, width, height int, obstaclesCoors [][2][2]int) {
	InitGrid(width, height, obstaclesCoors)

	fmt.Println("### GIVEN ###")
	fmt.Printf("start: %v, finish: %v\n", start, finish)

	marks := map[Point]string{start: "🏡", finish: "🏁"}
	PrintGrid(grid, marks)
	history := h.History()[1:]

	for i, vertex := range history {
		fmt.Printf("### STEP %d ####\n", i+1)
		fmt.Printf("velocity: %v, point: %v\n", vertex.speed, vertex.coor)

		marks[vertex.coor] = "👻" // choose your theme 🐸🐵🦠👀🚖🏂⛄
		PrintGrid(grid, marks)

		if i < len(history)-1 {
			marks[vertex.coor] = velocityEmoji(history[i+1].speed) // "👣"
		}

		println()
	}
}

func PrintGrid(g Grid, marks map[Point]string) {
	border := strings.Repeat("—", len(g[0])*5+1)
	println(border)
	for y := range g {
		print("|")
		for x := range g[y] {
			var valToPrint interface{}
			p := Point{x, y}
			markVal, ok := marks[p]
			if ok {
				valToPrint = markVal
			} else {
				cell := g[p.y][p.x]
				if cell == 0 {
					valToPrint = "  "
				} else {
					valToPrint = "🚫"
				}
			}
			fmt.Printf(" %v |", valToPrint)
		}
		println()
		println(border)
	}
}

func velocityEmoji(v Velocity) string {
	ver, hor := "", ""
	if v.DY > 0 {
		ver = "down"
	} else if v.DY < 0 {
		ver = "up"
	}
	if v.DX > 0 {
		hor = "right"
	} else if v.DX < 0 {
		hor = "left"
	}
	return directions[ver+hor] + " "
}

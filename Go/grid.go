package main

import (
	"math/rand"
	"strings"
	"time"
)

type Grid struct {
	grid                                   [][]bool
	height                                 int
	width                                  int
	liveSlotMinLiveNeighboursToKeepAlive   int
	liveSlotMaxLiveNeighboursToKeepAlive   int
	deadSlotMinLiveNeighboursToBringToLife int
	deadSlotMaxLiveNeighboursToBringToLife int
}

func createRandomGrid(height int, width int, threshold float32) [][]bool {
	rand.Seed(time.Now().UnixNano())
	ans := make([][]bool, height)
	for i := 0; i < height; i++ {
		ans[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			ans[i][j] = rand.Float32() < threshold
		}
	}
	return ans
}

func (g Grid) NextStep() *Grid {
	newGrid := make([][]bool, g.height)
	for i := 0; i < g.height; i++ {
		newGrid[i] = make([]bool, g.width)
	}

	nextGrid := Grid{
		grid:                                   newGrid,
		height:                                 g.height,
		width:                                  g.width,
		liveSlotMinLiveNeighboursToKeepAlive:   g.liveSlotMinLiveNeighboursToKeepAlive,
		liveSlotMaxLiveNeighboursToKeepAlive:   g.liveSlotMaxLiveNeighboursToKeepAlive,
		deadSlotMinLiveNeighboursToBringToLife: g.deadSlotMinLiveNeighboursToBringToLife,
		deadSlotMaxLiveNeighboursToBringToLife: g.deadSlotMaxLiveNeighboursToBringToLife,
	}

	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			nextGrid.grid[i][j] = SlotNextValue(g.grid[i][j], &g, i, j)

		}
	}

	return &nextGrid
}

func SlotNextValue(currentSlotValue bool, g *Grid, i int, j int) bool {
	liveNeighbours := GetNumberOfLiveNeighbours(g, i, j)
	// live slots handling
	if currentSlotValue {
		if liveNeighbours >= g.liveSlotMinLiveNeighboursToKeepAlive && liveNeighbours <= g.liveSlotMaxLiveNeighboursToKeepAlive {
			return true
		}
		return false
	}
	// dead slots handling
	if liveNeighbours >= g.deadSlotMinLiveNeighboursToBringToLife && liveNeighbours <= g.deadSlotMaxLiveNeighboursToBringToLife {
		return true
	}
	return false

}

func getNeighbour(g *Grid, i int, j int) bool {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	return g.grid[i][j]
}

func GetNumberOfLiveNeighbours(g *Grid, i int, j int) int {
	ans := 0

	upperLeft := getNeighbour(g, i-1, j-1)
	if upperLeft {
		ans += 1
	}

	upperMiddle := getNeighbour(g, i-1, j)
	if upperMiddle {
		ans += 1
	}

	upperRight := getNeighbour(g, i-1, j+1)
	if upperRight {
		ans += 1
	}

	middleLeft := getNeighbour(g, i, j-1)
	if middleLeft {
		ans += 1
	}

	middleRight := getNeighbour(g, i, j+1)
	if middleRight {
		ans += 1
	}

	lowerLeft := getNeighbour(g, i+1, j-1)
	if lowerLeft {
		ans += 1
	}

	lowerMiddle := getNeighbour(g, i+1, j)
	if lowerMiddle {
		ans += 1
	}

	lowerRight := getNeighbour(g, i+1, j+1)
	if lowerRight {
		ans += 1
	}
	return ans
}

func InitializeGrid(
	height int,
	width int,
	threshold float32,
	liveSlotMinLiveNeighboursToKeepAlive int,
	liveSlotMaxLiveNeighboursToKeepAlive int,
	deadSlotMinLiveNeighboursToBringToLife int,
	deadSlotMaxLiveNeighboursToBringToLife int) *Grid {
	grid := Grid{
		grid:                                   createRandomGrid(height, width, threshold),
		height:                                 height,
		width:                                  width,
		liveSlotMinLiveNeighboursToKeepAlive:   liveSlotMinLiveNeighboursToKeepAlive,
		liveSlotMaxLiveNeighboursToKeepAlive:   liveSlotMaxLiveNeighboursToKeepAlive,
		deadSlotMinLiveNeighboursToBringToLife: deadSlotMinLiveNeighboursToBringToLife,
		deadSlotMaxLiveNeighboursToBringToLife: deadSlotMaxLiveNeighboursToBringToLife,
	}
	return &grid
}

func ToStringArray(arr []bool) []string {
	var valuesText []string

	for i := range arr {
		number := arr[i]
		var text string
		if number {
			text = "*"
		} else {
			text = " "
		}
		valuesText = append(valuesText, text)
	}
	return valuesText
}

func (g Grid) String() string {
	var sb strings.Builder
	for _, gridRow := range g.grid {
		sb.WriteString(strings.Join(ToStringArray(gridRow), " "))
		sb.WriteString("\n")
	}
	return sb.String()
}

func flatten(g *Grid) []int {
	var bools []bool
	var ans []int
	for _, v1 := range g.grid {
		for _, b := range v1 {
			bools = append(bools, b)
		}
	}

	for _, b := range bools {
		s := 0
		if b {
			s = 1
		}
		ans = append(ans, s)
	}

	return ans
}

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var g = CreateRandomGrid(10,10, 0.3)


type Grid struct {
	grid [][]bool
	height int
	width int
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

func (oldGrid Grid) NextStep() *Grid{
	newGrid := make([][]bool, oldGrid.height)
	for i := 0; i < oldGrid.height; i++ {
		newGrid[i] = make([]bool, oldGrid.width)
	}

	nextGrid := Grid{
		grid:   newGrid,
		height: oldGrid.height,
		width:  oldGrid.width,
	}

	for i := 0; i < oldGrid.height; i++{
		for j := 0; j < oldGrid.width; j++ {
			nextGrid.grid[i][j] = SlotNextValue(oldGrid.grid[i][j], &oldGrid, i, j)

		}
	}

	return &nextGrid
}

func SlotNextValue(currentSlotValue bool, g *Grid, i int, j int) bool {
	liveNeighbours := GetNumberOfLiveNeighbours(g, i, j)
	if currentSlotValue {
		if liveNeighbours == 2 || liveNeighbours == 3 {
			return true
		}
		return false
	}
	if liveNeighbours == 3 {
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

func CreateRandomGrid(height int, width int, threshold float32) *Grid{
	grid := Grid{
		grid:   createRandomGrid(height, width, threshold),
		height: height,
		width:  width,
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
	for _, gridRow := range g.grid{
		sb.WriteString(strings.Join(ToStringArray(gridRow), " "))
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {
	http.HandleFunc("/next", getNextStepHttp)
	http.HandleFunc("/init", initHttp)
	http.ListenAndServe(":8090", nil)
}

func getNextStepHttp(writer http.ResponseWriter, request *http.Request) {
		g = g.NextStep()
		flat := flatten(g)
		bytes, _ := json.Marshal(flat)
		fmt.Fprint(writer, string(bytes))
}

func initHttp(writer http.ResponseWriter, request *http.Request) {

		g = CreateRandomGrid(10,10, 0.3)
		flat := flatten(g)
		bytes, _ := json.Marshal(flat)
		fmt.Fprint(writer, string(bytes))
}

func flatten(g *Grid) []int {
	var bools []bool
	var ans []int
	for _, v1 := range g.grid {
		for _, b := range v1 {
			bools = append(bools, b)
		}
	}

	for _, b := range bools{
		s := 0
		if b {
			s = 1
		}
		ans = append(ans, s)
	}

	return ans
}
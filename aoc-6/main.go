package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Maze struct {
	Rows int
	Cols int
	Map  []string
}

type Player struct {
	RowPos           int
	ColPos           int
	Direction        string
	Moves            int // distict moves
	VisitedPositions map[int][]int
}

func (p *Player) Init(m *Maze) {
	p.Moves = 0
	p.VisitedPositions = make(map[int][]int)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if m.Map[i][j] == '^' {
				p.RowPos = i
				p.ColPos = j
				p.Direction = "^"
				p.Moves++
				p.VisitedPositions[i] = []int{j}
			}
		}
	}
}

func (m *Maze) IsObstacle(row, col int) bool {
	return m.Map[row][col] == '#'
}

func (p *Player) String() string {
	return fmt.Sprintf("Player at (%d, %d) facing %s, moves: %d", p.RowPos, p.ColPos, p.Direction, p.Moves)
}

func (p *Player) TurnRight() {
	switch p.Direction {
	case "^":
		p.Direction = ">"
	case ">":
		p.Direction = "v"
	case "v":
		p.Direction = "<"
	case "<":
		p.Direction = "^"
	}
}

func (p *Player) UpdateVisitedPositions(x, y int) {
	yPositions, ok := p.VisitedPositions[x]
	if ok {
		if !slices.Contains(yPositions, y) {
			p.VisitedPositions[x] = append(p.VisitedPositions[x], y)
		}
	} else {
		p.VisitedPositions[x] = []int{y}
	}
}

func (p *Player) CanMove(m *Maze) bool {
	return p.RowPos > 0 && p.RowPos < m.Rows && p.ColPos > 0 && p.ColPos < m.Cols
}

func (p *Player) Move(m *Maze) {

	var x, y int

	switch p.Direction {
	case "^":
		x = p.RowPos - 1
		y = p.ColPos
	case ">":
		x = p.RowPos
		y = p.ColPos + 1
	case "v":
		x = p.RowPos + 1
		y = p.ColPos
	case "<":
		x = p.RowPos
		y = p.ColPos - 1
	}

	if !m.IsObstacle(x, y) {
		p.RowPos = x
		p.ColPos = y
		p.UpdateVisitedPositions(x, y)
	} else {
		p.TurnRight()
	}
}

func main() {

	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	mazeMap := make([]string, 0)

	for scanner.Scan() {
		mazeMap = append(mazeMap, scanner.Text())
	}

	rows, cols := len(mazeMap), len(mazeMap[0])
	maze := Maze{Rows: rows, Cols: cols, Map: mazeMap}
	p := Player{}
	p.Init(&maze)

	for p.CanMove(&maze) {
		p.Move(&maze)
	}

	sum := 0
	for _, v := range p.VisitedPositions {
		sum += len(v)
	}

	fmt.Println("Number of visited positions: ", sum)
}

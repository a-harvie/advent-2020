package lifeboi

import (
	"fmt"
	"strings"
)

// Lifeboi is conway-esque
type Lifeboi struct {
	grid   [][]int
	height int
	width  int
}

// Barren will never be full
// Empty is not currently full but could be
// Full is not currently empty but could be
const (
	Barren = -1
	Empty  = 0
	Full   = 1
)

// BoiMe will give you a Lifeboi if you treat it right
func BoiMe(input []string, legend map[string]int) (*Lifeboi, error) {
	g, err := parseGrid(input, legend)
	if err != nil {
		return nil, err
	}
	l := Lifeboi{
		grid:   g,
		height: len(g),
		width:  len(g[0]),
	}
	return &l, nil
}

// Step makes it go
func (b *Lifeboi) Step(tooCrowded int, countMethod string) {
	newGrid := make([][]int, len(b.grid))
	for i, line := range b.grid {
		newLine := make([]int, len(line))
		for j, cell := range line {
			skipBarren := countMethod != "neighbour"
			neighbours := b.countViewable(i, j, skipBarren)

			newLine[j] = cell
			if cell == Empty && neighbours == 0 {
				newLine[j] = Full
			}
			if cell == Full && neighbours >= tooCrowded {
				newLine[j] = Empty
			}
		}
		newGrid[i] = newLine
	}
	b.grid = newGrid
}

// GetPopulation gives you a number
func (b *Lifeboi) GetPopulation() int {
	pop := 0
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			pop += b.getCell(i, j)
		}
	}

	return pop
}

func (b *Lifeboi) countViewable(line, col int, skipBarren bool) int {
	viewable := 0

	viewable += b.findNextViewable(line, col, -1, 0, skipBarren)  // u
	viewable += b.findNextViewable(line, col, -1, 1, skipBarren)  // ur
	viewable += b.findNextViewable(line, col, 0, 1, skipBarren)   // r
	viewable += b.findNextViewable(line, col, 1, 1, skipBarren)   // dr
	viewable += b.findNextViewable(line, col, 1, 0, skipBarren)   // d
	viewable += b.findNextViewable(line, col, 1, -1, skipBarren)  // dl
	viewable += b.findNextViewable(line, col, 0, -1, skipBarren)  // l
	viewable += b.findNextViewable(line, col, -1, -1, skipBarren) // ul

	return viewable
}

func (b *Lifeboi) findNextViewable(line, col, v, h int, skipBarren bool) int {
	for {
		line += v
		col += h
		if line < 0 || col < 0 || line == b.height || col == b.width {
			break
		}
		switch b.grid[line][col] {
		case Barren:
			if skipBarren {
				continue
			}
			return Empty
		default:
			return b.grid[line][col]
		}

	}
	return Empty
}

func (b *Lifeboi) getCell(line, col int) int {
	if b.grid[line][col] >= 0 {
		return b.grid[line][col]
	}
	return 0
}

// Render is gonna show you something
func (b *Lifeboi) Render(tiles map[int]string) string {
	out := make([]string, len(b.grid))
	for i := 0; i < len(b.grid); i++ {
		line := make([]string, len(b.grid[i]))
		for j := 0; j < len(b.grid[i]); j++ {
			line[j] = tiles[b.grid[i][j]]
		}
		out[i] = strings.Join(line, "")
	}

	return strings.Join(out, "\n")
}

func parseGrid(input []string, legend map[string]int) ([][]int, error) {
	g := make([][]int, len(input))
	for i, line := range input {
		gridLine := make([]int, len(line))
		cells := strings.Split(line, "")
		for j, cell := range cells {
			c, ok := legend[cell]
			if !ok {
				return [][]int{}, fmt.Errorf("Couldn't find %s in legend", cell)
			}
			gridLine[j] = c
		}
		g[i] = gridLine
	}
	return g, nil
}

package nlifeboi

import "fmt"

// Thrlifeboi is conway-esque, but thicc
type Thrlifeboi struct {
	grid ThrGrid
}

// Forlifeboi is conway-esque, but hyperthicc
type Forlifeboi struct {
	grid ForGrid
}

// OneGrid is a 1D grid
type OneGrid []int

// TwoGrid is a 2D grid
type TwoGrid [][]int

// ThrGrid is a 3D grid, but spelled badly so that it lines up
type ThrGrid [][][]int

// ForGrid is a 4D grid, but spelled extra badly so that it also lines up
type ForGrid [][][][]int

// BoiMe will give you a Thrlifeboi, if you treat it right
func BoiMe(initialGrid ThrGrid) *Thrlifeboi {
	b := Thrlifeboi{
		grid: initialGrid,
	}
	return &b
}

// HyperBoiMe will give you a Forlifeboi, if you treat it right
func HyperBoiMe(initialGrid ForGrid) *Forlifeboi {
	b := Forlifeboi{
		grid: initialGrid,
	}
	return &b
}

// Step determines who lives or dies
func (b *Thrlifeboi) Step() {
	newGrid := make(ThrGrid, len(b.grid))
	count := 0
	for i := range b.grid {
		newPlane := make(TwoGrid, len(b.grid[i]))
		for j := range b.grid[i] {
			newRow := make(OneGrid, len(b.grid[i][j]))
			for k, active := range b.grid[i][j] {
				n := b.countNeighbours(i, j, k)
				newRow[k] = activeNext(active, n)
				count++
			}
			newPlane[j] = newRow
		}
		newGrid[i] = newPlane
	}
	fmt.Println("Step did ", count)
	b.grid = newGrid
}

func (b *Thrlifeboi) countNeighbours(z, y, x int) int {
	count := 0
	for i := z - 1; i <= z+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := x - 1; k <= x+1; k++ {
				if i < 0 || i == len(b.grid) ||
					j < 0 || j == len(b.grid[0]) ||
					k < 0 || k == len(b.grid[0][0]) ||
					(i == z && j == y && k == x) {
					continue
				}
				count += b.grid[i][j][k]
			}
		}
	}
	return count
}

// Step determines who lives or dies
func (b *Forlifeboi) Step() {
	newGrid := make(ForGrid, len(b.grid))
	for i := range b.grid {
		newCube := make(ThrGrid, len(b.grid[i]))
		// newGrid[i] = make(ThrGrid, len(b.grid[i]))
		for j := range b.grid[i] {
			newPlane := make(TwoGrid, len(b.grid[i][j]))
			// newGrid[i][j] = make(TwoGrid, len(b.grid[i][j]))
			for k := range b.grid[i][j] {
				newRow := make(OneGrid, len(b.grid[i][j][k]))
				// newGrid[i][j][k] = make(OneGrid, len(b.grid[i][j][k]))
				for l, active := range b.grid[i][j][k] {
					n := b.countNeighbours(i, j, k, l)
					// newGrid[i][j][k][l] = activeNext(active, n)
					newRow[l] = activeNext(active, n)
				}
				newPlane[k] = newRow
			}
			newCube[j] = newPlane
		}
		newGrid[i] = newCube
	}
	b.grid = newGrid
}

func (b *Forlifeboi) countNeighbours(w, z, y, x int) int {
	count := 0
	for i := w - 1; i <= w+1; i++ {
		for j := z - 1; j <= z+1; j++ {
			for k := y - 1; k <= y+1; k++ {
				for l := x - 1; l <= x+1; l++ {
					if i < 0 || i == len(b.grid) ||
						j < 0 || j == len(b.grid[i]) ||
						k < 0 || k == len(b.grid[i][j]) ||
						l < 0 || l == len(b.grid[i][j][k]) ||
						(i == w && j == z && k == y && l == x) {
						continue
					}
					count += b.grid[i][j][k][l]
				}
			}
		}
	}
	return count
}

func activeNext(active, neighours int) int {
	activeNext := 0
	if active > 0 && (neighours == 2 || neighours == 3) {
		activeNext = 1
	}
	if active == 0 && neighours == 3 {
		activeNext = 1
	}
	return activeNext
}

// GetGrid will grid you good
func (b *Thrlifeboi) GetGrid() ThrGrid {
	return b.grid
}

// GetPopulation will give you a number
func (b *Thrlifeboi) GetPopulation() int {
	count := 0
	for i := range b.grid {
		for j := range b.grid[i] {
			for k := range b.grid[i][j] {
				count += b.grid[i][j][k]
			}
		}
	}
	return count
}

// GetGrid will grid you good
func (b *Forlifeboi) GetGrid() ForGrid {
	return b.grid
}

// GetPopulation will give you a number
func (b *Forlifeboi) GetPopulation() int {
	count := 0
	for i := range b.grid {
		for j := range b.grid[i] {
			for k := range b.grid[i][j] {
				for l := range b.grid[i][j][k] {
					count += b.grid[i][j][k][l]
				}
			}
		}
	}
	return count
}

package utils

// Point represents a 2D coordinate
type Point struct {
	X, Y int
}

// Common direction vectors
var (
	Up        = Point{0, -1}
	Down      = Point{0, 1}
	Left      = Point{-1, 0}
	Right     = Point{1, 0}
	UpLeft    = Point{-1, -1}
	UpRight   = Point{1, -1}
	DownLeft  = Point{-1, 1}
	DownRight = Point{1, 1}

	// Cardinal directions (4-way)
	Cardinals = []Point{Up, Right, Down, Left}

	// Diagonal directions (4-way)
	Diagonals = []Point{UpRight, DownRight, DownLeft, UpLeft}

	// All 8 directions
	AllDirs = []Point{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
)

// Add returns p + other
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

// Sub returns p - other
func (p Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

// Scale returns p * n
func (p Point) Scale(n int) Point {
	return Point{p.X * n, p.Y * n}
}

// Manhattan returns the Manhattan distance from p to other
func (p Point) Manhattan(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors4 returns the 4 cardinal neighbors
func (p Point) Neighbors4() []Point {
	return []Point{
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y},
	}
}

// Neighbors8 returns all 8 neighbors
func (p Point) Neighbors8() []Point {
	return []Point{
		{p.X, p.Y - 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y + 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X - 1, p.Y - 1},
	}
}

// RotateRight rotates the direction 90 degrees clockwise
func (p Point) RotateRight() Point {
	return Point{-p.Y, p.X}
}

// RotateLeft rotates the direction 90 degrees counter-clockwise
func (p Point) RotateLeft() Point {
	return Point{p.Y, -p.X}
}

// InBounds checks if point is within bounds [0, width) x [0, height)
func (p Point) InBounds(width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

// GridBounds returns width and height of a 2D grid
func GridBounds[T any](grid [][]T) (width, height int) {
	if len(grid) == 0 {
		return 0, 0
	}
	return len(grid[0]), len(grid)
}

// GridAt safely gets a value from a grid, returning zero value if out of bounds
func GridAt[T any](grid [][]T, p Point) (T, bool) {
	var zero T
	if p.Y < 0 || p.Y >= len(grid) {
		return zero, false
	}
	if p.X < 0 || p.X >= len(grid[p.Y]) {
		return zero, false
	}
	return grid[p.Y][p.X], true
}

// FindInGrid finds all positions of a value in a byte grid
func FindInGrid(grid [][]byte, target byte) []Point {
	var result []Point
	for y, row := range grid {
		for x, cell := range row {
			if cell == target {
				result = append(result, Point{x, y})
			}
		}
	}
	return result
}

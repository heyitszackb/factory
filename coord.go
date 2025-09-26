package dreamer

type Coord struct {
	Row int
	Col int
}

func NewCoord(row int, col int) Coord {
	return Coord{row, col}
}

func (c *Coord) IsAtRightOf(coord Coord) bool {
	return coord.Row == c.Row && c.Col == coord.Col+1
}

func (c *Coord) IsAtLeftOf(coord Coord) bool {
	return coord.Row == c.Row && c.Col == coord.Col-1
}

func (c *Coord) IsAtTopOf(coord Coord) bool {
	return coord.Row == c.Row-1 && c.Col == coord.Col
}

func (c *Coord) IsAtBottomOf(coord Coord) bool {
	return coord.Row == c.Row+1 && c.Col == coord.Col
}

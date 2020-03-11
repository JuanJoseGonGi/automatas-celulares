package models

// Cell is a struct to represent a cell
type Cell struct {
	Index     int
	State     bool
	PrevState bool
}

// NewCell creates a new cell
func NewCell(index int) Cell {
	return Cell{
		Index:     index,
		State:     false,
		PrevState: false,
	}
}

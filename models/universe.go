package models

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// Universe is a struct to represent the universe
type Universe struct {
	Rule  int
	Rules []bool
	Image *image.Gray
	Data  [][]uint8
	Cells []*Cell
	Step  int
}

func decodeRule(rule int) []bool {
	rule64 := int64(rule)

	bin := strconv.FormatInt(rule64, 2)
	binWithZeros := fmt.Sprintf("%08s", bin)
	splittedBin := strings.Split(binWithZeros, "")

	rules := []bool{}
	for _, bit := range splittedBin {
		if bit == "1" {
			rules = append(rules, true)
			continue
		}
		rules = append(rules, false)
	}

	return rules
}

// NewUniverse creates a new instance of Universe
func NewUniverse(rule int, width int, height int) Universe {
	return Universe{
		Rule:  rule,
		Rules: decodeRule(rule),
		Image: image.NewGray(image.Rect(0, 0, width, height)),
		Step:  0,
	}
}

// AddCell adds a new cell to the Universe
func (universe *Universe) AddCell() {
	cell := NewCell(len(universe.Cells))
	universe.Cells = append(universe.Cells, &cell)
}

// Update updates the cells states
func (universe *Universe) Update() {
	for _, cell := range universe.Cells {
		cell.PrevState = cell.State

		previousCellState := false
		if cell.Index > 0 {
			previousCellState = universe.Cells[cell.Index-1].PrevState
		}

		currentCellState := cell.State

		nextCellState := false
		if cell.Index < len(universe.Cells)-1 {
			nextCellState = universe.Cells[cell.Index+1].State
		}

		if previousCellState && currentCellState && nextCellState {
			cell.State = universe.Rules[0]
		} else if previousCellState && currentCellState && !nextCellState {
			cell.State = universe.Rules[1]
		} else if previousCellState && !currentCellState && nextCellState {
			cell.State = universe.Rules[2]
		} else if previousCellState && !currentCellState && !nextCellState {
			cell.State = universe.Rules[3]
		} else if !previousCellState && currentCellState && nextCellState {
			cell.State = universe.Rules[4]
		} else if !previousCellState && currentCellState && !nextCellState {
			cell.State = universe.Rules[5]
		} else if !previousCellState && !currentCellState && nextCellState {
			cell.State = universe.Rules[6]
		} else if !previousCellState && !currentCellState && !nextCellState {
			cell.State = universe.Rules[7]
		}
	}
}

// DrawRow draws the current step row of the cell
func (universe *Universe) DrawRow() {
	if universe.Step == 0 {
		middle := len(universe.Cells) / 2
		universe.Cells[middle].State = true
	}

	for index, cell := range universe.Cells {
		if cell.State {
			universe.Image.SetGray(index, universe.Step, color.Gray{uint8(255)})
		}
	}

	universe.Step++
}

// Save saves the generated image to the disk
func (universe *Universe) Save() error {
	file, err := os.OpenFile(fmt.Sprintf("RULE-%d.png", universe.Rule), os.O_CREATE|os.O_WRONLY, 0644) // #nosec
	if err != nil {
		return err
	}

	err = png.Encode(file, universe.Image)
	if err != nil {
		return err
	}
	return nil
}

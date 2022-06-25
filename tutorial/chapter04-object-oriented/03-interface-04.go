package main

import (
	"fmt"
	"time"
)

/**
 * @author       weimenghua
 * @time         2024-07-23 16:50
 * @description
 */

func NewDataCells(names []string, createdAts []time.Time) []DataCell {
	cells := make([]DataCell, len(names))
	for i, name := range names {
		cells[i] = &dataCell{
			name:      name + "_tmp",
			createdAt: createdAts[i],
		}
	}
	return cells
}

type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

type dataCell struct {
	name      string
	createdAt time.Time
}

func (c *dataCell) GetCreation() time.Time {
	return c.createdAt
}

func (c *dataCell) GetName() string {
	return c.name
}

func main() {
	var cells []DataCell = NewDataCells([]string{"cell1", "cell2", "cell3"}, []time.Time{time.Now(), time.Now().Add(time.Hour), time.Now().Add(2 * time.Hour)})
	for _, cell := range cells {
		fmt.Println(cell.GetName(), cell.GetCreation())
	}
}

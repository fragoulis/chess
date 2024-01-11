package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Black = color.Gray16{0x5500}
)

type BoardCell struct {
	Scene

	// Piece holds the piece that currently occupies the cell.
	Piece *Piece

	// Color of the cell when rendered
	Color color.Color

	// The position of the cell on the board
	Row, Column int

	// The drawing target shape
	Image *ebiten.Image

	// The cell dimentions
	Size int

	Hidden bool
}

func NewBoardCell(id int, row, column int, size int) *BoardCell {
	return &BoardCell{
		Scene: Scene{
			ID: fmt.Sprintf("cell_%d", id),
		},
		Row:    row,
		Column: column,
		Image:  ebiten.NewImage(size, size),
		Color:  color.RGBA{0xbb, 0xad, 0xa0, 0xff},
		Size:   size,
		Hidden: true,
	}
}

func (c *BoardCell) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	opts.GeoM.Translate(
		float64((c.Size+14)*c.Column),
		float64((c.Size+14)*c.Row),
	)

	if !c.Hidden {
		c.Image.Fill(c.Color)
		screen.DrawImage(c.Image, &opts)
	}

	c.Scene.Draw(screen, opts)
}

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

	opts ebiten.DrawImageOptions

	Board *Board
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

	// Save the options in order for us to use them later and
	// translate the image position back to screenspace.
	c.opts = opts

	if !c.Hidden {
		c.Image.Fill(c.Color)
		screen.DrawImage(c.Image, &opts)
	}

	c.Scene.Draw(screen, opts)
}

func (c *BoardCell) Update() error {
	if c.Piece != nil && c.Board.Selected == nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Grab the cursor absolute position.
		mx, my := ebiten.CursorPosition()
		point := image.Pt(mx, my)

		// Translate the cell position to screenspace (absolute).
		x0, y0 := c.opts.GeoM.Apply(
			float64(c.Image.Bounds().Min.X),
			float64(c.Image.Bounds().Min.Y),
		)
		x1, y1 := c.opts.GeoM.Apply(
			float64(c.Image.Bounds().Max.X),
			float64(c.Image.Bounds().Max.X),
		)
		bbox := image.Rect(int(x0), int(y0), int(x1), int(y1))

		if point.In(bbox) {
			c.SetDrawOrder(99)
			c.Piece.Dragged = true
			c.Board.Selected = c.Piece
			fmt.Printf("Clicked cell %s => %s\n", c.ID, point)
		}
	} else if c.Piece != nil && c.Piece.Dragged && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		c.SetDrawOrder(0)
		c.Piece.Dragged = false
		c.Board.Selected = nil
	}

	return c.Scene.Update()
}

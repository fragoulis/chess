package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Piece struct {
	Scene

	// The cell at which the piece resides
	Cell *BoardCell

	// The piece image
	Image *ebiten.Image

	// Whether the piece is currently selected
	Dragged bool

	Hidden bool

	opts ebiten.DrawImageOptions
}

func NewPiece(id string, img *ebiten.Image) *Piece {
	return &Piece{
		Scene: Scene{
			ID: id,
		},
		Image: img,
	}
}

func (p *Piece) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	if p.Hidden {
		return
	}

	if p.Dragged {
		// Grab the cursor absolute position.
		mx, my := ebiten.CursorPosition()
		// point := image.Pt(mx, my)

		x := float64(mx) - float64(p.Cell.Size)*0.5
		y := float64(my) - float64(p.Cell.Size)*0.5

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(p.Image, op)

		return
	}

	// Save the options in order for us to use them later and
	// translate the image position back to screenspace.
	p.opts = opts

	screen.DrawImage(p.Image, &opts)
}

// func (p *Piece) MoveToCell(cell *BoardCell) {
// 	// Empty currently occupied cell
// 	p.Cell.Piece = nil

// 	// Move to new cell
// 	cell.Piece = p
// 	p.Cell = cell
// }

// func (p *Piece) CanMoveToCell(cell *BoardCell) bool {
// 	if cell.Piece != nil {
// 		return false
// 	}

// 	switch p.ID {
// 	case DQ, LQ:
// 		return false
// 	case DK, LK:
// 		return ((cell.PosX == p.PosX()+1) ||
// 			(cell.PosX == p.PosX()-1) ||
// 			(cell.PosY == p.PosY()+1) ||
// 			(cell.PosY == p.PosY()-1))
// 	}

// 	return true
// }

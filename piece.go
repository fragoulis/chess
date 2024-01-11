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
	// Dragged bool

	Hidden bool
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

	// x := offsetDrawX + (tileSize * p.PosX())
	// y := tileSize * p.PosY()
	// if p.Dragged {
	// 	mx, my := ebiten.CursorPosition()
	// 	x = mx - tileSize/2
	// 	y = my - tileSize/2
	// }

	// op.GeoM.Translate(400, 400)
	// op.GeoM.Scale(targetWidth/originalWidth, targetHeight/originalHeight)
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

package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DarkQueenImage = iota
	DarkKingImage
	DarkRookImage
	DarkKnightImage
	DarkBishopImage
	DarkPawnImage
	LightQueenImage
	LightKingImage
	LightRookImage
	LightKnightImage
	LightBishopImage
	LightPawnImage
)

const (
	E = iota
	DR
	DKn
	DB
	DQ
	DK
	DP
	LR
	LKn
	LB
	LQ
	LK
	LP
)

var (
	PieceImageIndex = map[int]int{
		DQ:  DarkQueenImage,
		DK:  DarkKingImage,
		DR:  DarkRookImage,
		DKn: DarkKnightImage,
		DB:  DarkBishopImage,
		DP:  DarkPawnImage,
		LQ:  LightQueenImage,
		LK:  LightKingImage,
		LR:  LightRookImage,
		LKn: LightKnightImage,
		LB:  LightBishopImage,
		LP:  LightPawnImage,
	}
)

type Piece struct {
	ID      int
	Cell    *BoardCell
	Image   *ebiten.Image
	Dragged bool
	Bbox    image.Rectangle
	Alive   bool
}

func NewPiece(id int, img *ebiten.Image) *Piece {
	return &Piece{
		ID:    id,
		Image: img,
		Alive: true,
	}
}

func (p *Piece) PosX() int {
	return p.Cell.PosX
}

func (p *Piece) PosY() int {
	return p.Cell.PosY
}

func (p *Piece) Draw(screen *ebiten.Image) {
	if !p.Alive {
		return
	}

	x := offsetDrawX + (tileSize * p.PosX())
	y := tileSize * p.PosY()

	p.Bbox = image.Rect(x, y, x+tileSize, y+tileSize)

	if p.Dragged {
		mx, my := ebiten.CursorPosition()
		x = mx - tileSize/2
		y = my - tileSize/2
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(p.Image, op)
}

func (p *Piece) MoveToCell(cell *BoardCell) {
	// Empty currently occupied cell
	p.Cell.Piece = nil

	// Move to new cell
	cell.Piece = p
	p.Cell = cell
}

func (p *Piece) CanMoveToCell(cell *BoardCell) bool {
	if cell.Piece != nil {
		return false
	}

	switch p.ID {
	case DQ, LQ:
		return false
	case DK, LK:
		return ((cell.PosX == p.PosX()+1) ||
			(cell.PosX == p.PosX()-1) ||
			(cell.PosY == p.PosY()+1) ||
			(cell.PosY == p.PosY()-1))
	}

	return true
}

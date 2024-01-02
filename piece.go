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

var (
	PieceImageIndex = map[int]int{
		DarkQueen:    DarkQueenImage,
		DarkKing:     DarkKingImage,
		DarkRookL:    DarkRookImage,
		DarkRookR:    DarkRookImage,
		DarkKnightL:  DarkKnightImage,
		DarkKnightR:  DarkKnightImage,
		DarkBishopL:  DarkBishopImage,
		DarkBishopR:  DarkBishopImage,
		DarkPawn1:    DarkPawnImage,
		DarkPawn2:    DarkPawnImage,
		DarkPawn3:    DarkPawnImage,
		DarkPawn4:    DarkPawnImage,
		DarkPawn5:    DarkPawnImage,
		DarkPawn6:    DarkPawnImage,
		DarkPawn7:    DarkPawnImage,
		DarkPawn8:    DarkPawnImage,
		LightQueen:   LightQueenImage,
		LightKing:    LightKingImage,
		LightRookL:   LightRookImage,
		LightRookR:   LightRookImage,
		LightKnightL: LightKnightImage,
		LightKnightR: LightKnightImage,
		LightBishopL: LightBishopImage,
		LightBishopR: LightBishopImage,
		LightPawn1:   LightPawnImage,
		LightPawn2:   LightPawnImage,
		LightPawn3:   LightPawnImage,
		LightPawn4:   LightPawnImage,
		LightPawn5:   LightPawnImage,
		LightPawn6:   LightPawnImage,
		LightPawn7:   LightPawnImage,
		LightPawn8:   LightPawnImage,
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
	case DarkQueen, LightQueen:
		return false
	case DarkKing, LightKing:
		return ((cell.PosX == p.PosX()+1) ||
			(cell.PosX == p.PosX()-1) ||
			(cell.PosY == p.PosY()+1) ||
			(cell.PosY == p.PosY()-1))
	}

	return true
}

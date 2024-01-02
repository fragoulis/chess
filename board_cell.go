package main

type BoardCell struct {
	ID    int
	Piece *Piece
	PosX  int
	PosY  int
}

func NewBoardCell(id int, p *Piece, x, y int) *BoardCell {
	cell := &BoardCell{
		ID:    id,
		Piece: p,
		PosX:  x,
		PosY:  y,
	}

	if p != nil {
		p.Cell = cell
	}

	return cell
}

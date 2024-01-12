package main

import (
	"bytes"
	"image"
	"math"

	"github.com/go-errors/errors"
	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
	_ "image/png"
)

const (
	DarkQueenImageIndex = iota
	DarkKingImageIndex
	DarkRookImageIndex
	DarkKnightImageIndex
	DarkBishopImageIndex
	DarkPawnImageIndex
	LightQueenImageIndex
	LightKingImageIndex
	LightRookImageIndex
	LightKnightImageIndex
	LightBishopImageIndex
	LightPawnImageIndex
)

var (
	PieceNameByIndex = [12]string{
		"DarkQueen",
		"DarkKing",
		"DarkRook",
		"DarkKnight",
		"DarkBishop",
		"DarkPawn",
		"LightQueen",
		"LightKing",
		"LightRook",
		"LightKnight",
		"LightBishop",
		"LightPawn",
	}

	PieceIndexByName = map[string]int{
		"DarkQueen":   DarkQueenImageIndex,
		"DarkKing":    DarkKingImageIndex,
		"DarkRook":    DarkRookImageIndex,
		"DarkKnight":  DarkKnightImageIndex,
		"DarkBishop":  DarkBishopImageIndex,
		"DarkPawn":    DarkPawnImageIndex,
		"LightQueen":  LightQueenImageIndex,
		"LightKing":   LightKingImageIndex,
		"LightRook":   LightRookImageIndex,
		"LightKnight": LightKnightImageIndex,
		"LightBishop": LightBishopImageIndex,
		"LightPawn":   LightPawnImageIndex,
	}
)

var (
	// Black = color.Gray16{0x5500}
	// White = color.Gray16{0xffff}

	//go:embed board.png
	BoardPng []byte
)

type Board struct {
	Scene

	Image       *ebiten.Image
	Cells       [64]*BoardCell
	PieceImages [12]*ebiten.Image
	Pieces      []*Piece

	// The piece that is currently being dragged.
	Selected *Piece
}

func NewBoard() (*Board, error) {
	board := &Board{
		Scene: Scene{
			ID: "board",
		},
	}

	err := board.loadImage()
	if err != nil {
		return nil, err
	}

	board.initializeCells()

	err = board.loadPieceSet()
	if err != nil {
		return nil, err
	}

	board.placePiecesOnBoard()

	return board, nil
}

func (b *Board) initializeCells() {
	for i := range b.Cells {
		row := int(math.Floor(float64(i) / 8.0))
		column := i % 8

		cell := NewBoardCell(i, row, column, 50)
		b.Register(cell)
		b.Cells[i] = cell
	}
}

func (b *Board) loadImage() error {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(BoardPng))
	if err != nil {
		return errors.New(err)
	}

	b.Image = ebiten.NewImageFromImage(img)

	return nil
}

func (b *Board) loadPieceSet() error {
	pieceSet, err := NewPieceSet()
	if err != nil {
		return err
	}

	for i := range b.PieceImages {
		img, err := pieceSet.LoadPieceImage(i)
		if err != nil {
			return err
		}
		b.PieceImages[i] = img
	}

	return nil
}

func (b *Board) placePiecesOnBoard() {
	b.placePiece("DarkRook", 0+0*8)
	b.placePiece("DarkKnight", 1+0*8)
	b.placePiece("DarkBishop", 2+0*8)
	b.placePiece("DarkQueen", 3+0*8)
	b.placePiece("DarkKing", 4+0*8)
	b.placePiece("DarkBishop", 5+0*8)
	b.placePiece("DarkKnight", 6+0*8)
	b.placePiece("DarkRook", 7+0*8)
	b.placePiece("DarkPawn", 0+1*8)
	b.placePiece("DarkPawn", 1+1*8)
	b.placePiece("DarkPawn", 2+1*8)
	b.placePiece("DarkPawn", 3+1*8)
	b.placePiece("DarkPawn", 4+1*8)
	b.placePiece("DarkPawn", 5+1*8)
	b.placePiece("DarkPawn", 6+1*8)
	b.placePiece("DarkPawn", 7+1*8)

	b.placePiece("LightRook", 0+7*8)
	b.placePiece("LightKnight", 1+7*8)
	b.placePiece("LightBishop", 2+7*8)
	b.placePiece("LightQueen", 3+7*8)
	b.placePiece("LightKing", 4+7*8)
	b.placePiece("LightBishop", 5+7*8)
	b.placePiece("LightKnight", 6+7*8)
	b.placePiece("LightRook", 7+7*8)
	b.placePiece("LightPawn", 0+6*8)
	b.placePiece("LightPawn", 1+6*8)
	b.placePiece("LightPawn", 2+6*8)
	b.placePiece("LightPawn", 3+6*8)
	b.placePiece("LightPawn", 4+6*8)
	b.placePiece("LightPawn", 5+6*8)
	b.placePiece("LightPawn", 6+6*8)
	b.placePiece("LightPawn", 7+6*8)
}

func (b *Board) placePiece(pieceImageName string, cellIndex int) {
	// Get piece image by image id
	pieceImage := b.PieceImages[PieceIndexByName[pieceImageName]]

	piece := NewPiece(pieceImageName, pieceImage)
	b.Pieces = append(b.Pieces, piece)

	cell := b.Cells[cellIndex]
	cell.Piece = piece
	piece.Cell = cell
	cell.Register(piece)
}

func (b *Board) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	padding := 0.5 * (float64(screen.Bounds().Dx()) - float64(b.Image.Bounds().Dx()))

	opts.GeoM.Translate(padding, 0)
	screen.DrawImage(b.Image, &opts)

	// Move drawing cursor within the board
	// to draw the cells and pieces.
	opts.GeoM.Translate(45, 45)

	b.Scene.Draw(screen, opts)
}

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

	Image  *ebiten.Image
	Cells  [64]*BoardCell
	Pieces [12]*Piece

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

	for i := range b.Pieces {
		img, err := pieceSet.LoadPieceImage(0)
		if err != nil {
			return err
		}

		id := PieceNameByIndex[i]
		piece := NewPiece(id, img)
		b.Pieces[i] = piece

		// Do not register pieces with the board but with
		// the individual cells instead.
		// b.Register(piece)
	}

	return nil
}

func (b *Board) placePiecesOnBoard() {
	b.placePiece("DarkRook", 0)
	b.placePiece("LightRook", 1)
}

func (b *Board) placePiece(id string, cellIndex int) {
	piece := b.Pieces[PieceIndexByName[id]]
	cell := b.Cells[cellIndex]
	cell.Piece = piece
	piece.Cell = cell
	cell.Register(piece)
}

func (b *Board) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	// originalWidth := float64(b.Image.Bounds().Dx())
	// originalHeight := float64(b.Image.Bounds().Dy())
	// targetWidth := float64(screen.Bounds().Dx()) - padding
	// targetHeight := float64(screen.Bounds().Dy())
	padding := 0.5 * (float64(screen.Bounds().Dx()) - float64(b.Image.Bounds().Dx()))

	// opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(padding, 0)
	// opts.GeoM.Scale(targetWidth/originalWidth, targetHeight/originalHeight)
	screen.DrawImage(b.Image, &opts)

	// Move drawing cursor within the board
	// to draw the cells and pieces.
	opts.GeoM.Translate(50, 50)

	b.Scene.Draw(screen, opts)
}

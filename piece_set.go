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
	PieceImageWidth  = 60
	PieceImageHeight = 60
)

var (
	//go:embed pieces.png
	PiecesPng []byte
)

type PieceSet struct {
	Image *ebiten.Image
}

func NewPieceSet() (*PieceSet, error) {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(PiecesPng))
	if err != nil {
		return nil, errors.New(err)
	}

	return &PieceSet{
		Image: ebiten.NewImageFromImage(img),
	}, nil
}

func (s *PieceSet) LoadPieceImage(i int) (*ebiten.Image, error) {
	if i < 0 || i > 12 {
		return nil, errors.New("Out of bounds piece index")
	}

	x := i % 6
	y := int(math.Floor(float64(i) / 6.0))

	area := image.Rect(
		PieceImageWidth*x, // top left
		PieceImageHeight*y,
		PieceImageWidth*(x+1), // bottom right
		PieceImageHeight*(y+1),
	)

	// The index represents the position of the piece in the image.
	return s.Image.SubImage(area).(*ebiten.Image), nil
}

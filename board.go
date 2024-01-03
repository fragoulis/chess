package main

import (
	"bytes"
	"image"
	"image/color"

	"github.com/go-errors/errors"
	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
	_ "image/png"
)

var (
	Black = color.Gray16{0x5500}
	White = color.Gray16{0xffff}

	//go:embed board.png
	PiecesPng []byte
)

type Board struct {
	Image *ebiten.Image
}

func NewBoard() (*Board, error) {
	b := &Board{}

	err := b.loadImage()
	if err != nil {
		return nil, err
	}

	return b, err
}

func (b *Board) loadImage() error {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(PiecesPng))
	if err != nil {
		return errors.New(err)
	}

	b.Image = ebiten.NewImageFromImage(img)

	return nil
}

func (b *Board) Render(screen *Screen) {
	originalWidth := float64(b.Image.Bounds().Dx())
	originalHeight := float64(b.Image.Bounds().Dy())
	padding := 0.2 * float64(screen.Width)
	targetWidth := float64(screen.Width) - padding
	targetHeight := float64(screen.Height)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(padding, 0)
	op.GeoM.Scale(targetWidth/originalWidth, targetHeight/originalHeight)
	screen.Image.DrawImage(b.Image, op)
}

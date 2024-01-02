package main

import (
	"bytes"
	"image"
	"image/color"
	"sort"

	"github.com/go-errors/errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	_ "embed"
	_ "image/png"
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
	Black = color.Gray16{0x5500}
	White = color.Gray16{0xffff}

	//go:embed chess.png
	PiecesPng []byte
)

type Board struct {
	Cells    [64]*BoardCell
	Pieces   []*Piece
	Dragging *Piece

	TileSize    int
	OffsetDrawX int

	PieceImages [12]*ebiten.Image
}

func NewBoard() (*Board, error) {
	b := &Board{
		TileSize:    tileSize,
		OffsetDrawX: offsetDrawX,
	}

	err := b.loadPieceImages()
	if err != nil {
		return nil, err
	}

	return b, err
}

// func (b *Board) At(x, y int) (int, int) {
// 	return ((x - b.OffsetDrawX) / b.TileSize), (y / b.TileSize)
// }

func (b *Board) At(x, y int) *BoardCell {
	cellX, cellY := ((x - b.OffsetDrawX) / tileSize), (y / tileSize)

	if cellX < 0 {
		cellX = 0
	} else if cellX > 7 {
		cellX = 7
	}

	if cellY < 0 {
		cellY = 0
	} else if cellY > 7 {
		cellY = 7
	}

	i := cellY*8 + cellX
	return b.Cells[i]
}

func (b *Board) loadPieceImages() error {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(PiecesPng))
	if err != nil {
		return errors.New(err)
	}

	tilesImage := ebiten.NewImageFromImage(img)

	sz := 60 // tile size
	for i := 0; i < 2; i++ {
		for j := 0; j < 6; j++ {
			area := image.Rect(sz*j, sz*i, sz*(j+1), sz*(i+1))
			b.PieceImages[i*6+j] = tilesImage.SubImage(area).(*ebiten.Image)
		}
	}

	return nil
}

func (b *Board) Initialize(setup [64]int) error {
	for i := 0; i < 64; i++ {
		cellID := setup[i]
		cellX := i % 8
		cellY := i / 8
		var piece *Piece

		// fmt.Printf("cell:(%d,%d) id:%d\n", cellX, cellY, cellID)

		if cellID > 0 {
			piece = NewPiece(cellID, b.PieceImages[PieceImageIndex[cellID]])
			b.Pieces = append(b.Pieces, piece)
		}

		b.Cells[i] = NewBoardCell(i, piece, cellX, cellY)
	}

	return nil
}

func (b *Board) DrawAll(screen *ebiten.Image) {
	b.DrawBoard(screen)

	// if b.Dragging != nil {
	// 	b.DrawPath(screen, b.Dragging)
	// }

	sort.SliceStable(b.Pieces, func(i, j int) bool {
		return !b.Pieces[i].Dragged
	})

	for _, p := range b.Pieces {
		p.Draw(screen)
	}
}

// func (b *Board) DrawPath(screen *ebiten.Image, p *Piece) {
// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX()+1)*b.TileSize),
// 		float64((p.PosY()+1)*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX()+1)*b.TileSize),
// 		float64((p.PosY()-1)*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX()-1)*b.TileSize),
// 		float64((p.PosY()-1)*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX()-1)*b.TileSize),
// 		float64((p.PosY()+1)*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX()+1)*b.TileSize),
// 		float64((p.PosY())*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX()-1)*b.TileSize),
// 		float64((p.PosY())*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX())*b.TileSize),
// 		float64((p.PosY()+1)*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)

// 	ebitenutil.DrawRect(
// 		screen,
// 		float64(b.OffsetDrawX+(p.PosX())*b.TileSize),
// 		float64((p.PosY()-1)*b.TileSize),
// 		float64(b.TileSize),
// 		float64(b.TileSize),
// 		color.RGBA{255, 100, 100, 180},
// 	)
// }

func (b *Board) DrawBoard(screen *ebiten.Image) {
	dark := color.RGBA{50, 100, 50, 255}
	light := color.RGBA{255, 255, 255, 255}
	w := b.TileSize
	h := b.TileSize

	for i := 0; i < 64; i++ {
		x := (i % 8) * b.TileSize
		y := (i / 8) * b.TileSize

		c := dark
		if (i/8)%2 == 0 {
			if i%2 == 0 {
				c = light
			}
		} else {
			if i%2 != 0 {
				c = light
			}
		}

		vector.DrawFilledRect(
			screen,
			float32(x+b.OffsetDrawX),
			float32(y),
			float32(w),
			float32(h),
			c,
			true,
		)
	}
}

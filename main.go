package main

import (
	"fmt"
	"image"
	"os"

	"github.com/go-errors/errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Empty = iota
	DarkQueen
	DarkKing
	DarkRookL
	DarkRookR
	DarkKnightL
	DarkKnightR
	DarkBishopL
	DarkBishopR
	DarkPawn1
	DarkPawn2
	DarkPawn3
	DarkPawn4
	DarkPawn5
	DarkPawn6
	DarkPawn7
	DarkPawn8
	LightQueen
	LightKing
	LightRookL
	LightRookR
	LightKnightL
	LightKnightR
	LightBishopL
	LightBishopR
	LightPawn1
	LightPawn2
	LightPawn3
	LightPawn4
	LightPawn5
	LightPawn6
	LightPawn7
	LightPawn8
)

var (
	InitialBoardSetup = [64]int{
		DR, DKn, DB, DQ, DK, DB, DKn, DR,
		DP, DP, DP, DP, DP, DP, DP, DP,
		E, E, E, E, E, E, E, E,
		E, E, E, E, E, E, E, E,
		E, E, E, E, E, E, E, E,
		E, E, E, E, E, E, E, E,
		LP, LP, LP, LP, LP, LP, LP, LP,
		LR, LKn, LB, LQ, LK, LB, LKn, LR,
	}
)

//
//
//

const (
	tileSize     = 60
	screenWidth  = tileSize * 12
	screenHeight = tileSize * 8
	offsetDrawX  = tileSize * 2
)

func Fatal(err error) {
	if err == nil {
		return
	}

	// if errors.Is(err, errors.Error) {
	// 	fmt.Println(err.(*errors.Error).ErrorStack())
	// } else {
	// 	log.Fatal(err)
	// }

	fmt.Println(err.(*errors.Error).ErrorStack())
	os.Exit(1)
}

type Game struct {
	Board *Board
}

func (g *Game) Update() error {
	if g.Board.Dragging == nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		point := image.Pt(mx, my)

		for _, piece := range g.Board.Pieces {
			if !piece.Alive {
				continue
			}

			piece.Dragged = point.In(piece.Bbox)
			if piece.Dragged {
				g.Board.Dragging = piece
			}
		}
	}

	if g.Board.Dragging != nil && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		cell := g.Board.At(mx, my)

		if g.Board.Dragging.CanMoveToCell(cell) {
			g.Board.Dragging.MoveToCell(cell)
		}

		g.Board.Dragging.Dragged = false
		g.Board.Dragging = nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Board.DrawAll(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Chess")

	board, err := NewBoard()
	Fatal(err)

	board.Initialize(InitialBoardSetup)
	Fatal(err)

	err = ebiten.RunGame(&Game{Board: board})
	Fatal(err)
}

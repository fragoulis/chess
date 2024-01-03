package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Screen struct {
	Width  int
	Height int
	Image  *ebiten.Image
}

type Game struct {
	Screen   *Screen
	Renderer *Renderer
	Board    *Board
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen.Image = screen
	g.Renderer.Render(g.Screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	screen := &Screen{
		Width:  800,
		Height: 600,
	}

	renderer := NewRenderer()

	board, err := NewBoard()
	Fatal(err)

	renderer.Register(board)

	ebiten.SetWindowSize(screen.Width, screen.Height)
	ebiten.SetWindowTitle("Chess")
	// ebiten.SetWindowPosition(1400, 800)
	err = ebiten.RunGameWithOptions(
		&Game{
			Screen:   screen,
			Renderer: renderer,
			Board:    board,
		},
		&ebiten.RunGameOptions{
			InitUnfocused: true,
		},
	)
	Fatal(err)
}

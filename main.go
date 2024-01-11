package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Scene
}

func NewGame() (*Game, error) {
	game := &Game{}

	board, err := NewBoard()
	if err != nil {
		return nil, err
	}

	game.Register(board)

	return game, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen, ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game, err := NewGame()
	Fatal(err)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Chess")
	// ebiten.SetWindowPosition(1400, 800)
	err = ebiten.RunGameWithOptions(
		game,
		&ebiten.RunGameOptions{
			InitUnfocused: true,
		},
	)
	Fatal(err)
}

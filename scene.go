package main

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Updatable interface {
	Update() error
}

type Drawable interface {
	Draw(*ebiten.Image, ebiten.DrawImageOptions)
	DrawOrder() int
	SetDrawOrder(o int)
}

type Point struct {
	X, Y float64
}

type Scene struct {
	Components []interface{}

	// Do not draw
	// Hidden bool

	// Do not update
	Disabled bool

	// Drawing offset
	// Offset Point

	// Unique ID for each scene component
	ID string

	drawOrder int
}

func (s *Scene) Register(component interface{}) {
	s.Components = append(s.Components, component)
}

func (s *Scene) Update() error {
	if s.Disabled {
		return nil
	}

	for _, c := range s.Components {
		if u, ok := c.(Updatable); ok {
			if err := u.Update(); err != nil {
				return err
			}
		}
	}

	return nil
}

type Drawables []Drawable

func (a Drawables) Len() int           { return len(a) }
func (a Drawables) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Drawables) Less(i, j int) bool { return a[i].DrawOrder() < a[j].DrawOrder() }

func (s *Scene) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	// if s.Hidden {
	// 	return
	// }

	// Copy opts
	// newOpts := opts

	// Reset GeoM to an identity matrix
	// newOpts.GeoM.Reset()

	// Translate scenespace to worldspace (translate by s.Offset)
	// newOpts.GeoM.Translate(s.Offset.X, s.Offset.Y)

	// Reapply the original opts.GeoM to translate into screenspace or whatever
	// newOpts.GeoM.Concat(opts.GeoM)

	drawables := make(Drawables, 0, len(s.Components))
	for _, c := range s.Components {
		if d, ok := c.(Drawable); ok {
			drawables = append(drawables, d)
		}
	}

	// Draw the subcomponents
	sort.Sort(drawables)

	for _, d := range drawables {
		d.Draw(screen, opts)
	}
}

func (s *Scene) DrawOrder() int {
	return s.drawOrder
}

func (s *Scene) SetDrawOrder(o int) {
	s.drawOrder = o
}

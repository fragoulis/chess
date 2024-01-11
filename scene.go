package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Updatable interface {
	Update() error
}

type Drawable interface {
	Draw(*ebiten.Image, ebiten.DrawImageOptions)
}

type Point struct {
	X, Y float64
}

type Scene struct {
	Components []interface{}

	// Do not draw
	Hidden bool

	// Do not update
	Disabled bool

	// Drawing offset
	Offset Point

	// Unique ID for each scene component
	ID string
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

func (s *Scene) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	if s.Hidden {
		return
	}

	// Copy opts
	// newOpts := opts

	// Reset GeoM to an identity matrix
	// newOpts.GeoM.Reset()

	// Translate scenespace to worldspace (translate by s.Offset)
	// newOpts.GeoM.Translate(s.Offset.X, s.Offset.Y)

	// Reapply the original opts.GeoM to translate into screenspace or whatever
	// newOpts.GeoM.Concat(opts.GeoM)

	// Draw the subcomponents
	for _, c := range s.Components {
		if d, ok := c.(Drawable); ok {
			d.Draw(screen, opts)
		}

	}
}

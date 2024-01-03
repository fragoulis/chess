package main

type Renderable interface {
	Render(screen *Screen)
}

type Renderer struct {
	renderables map[int]Renderable
}

func NewRenderer() *Renderer {
	return &Renderer{
		renderables: make(map[int]Renderable),
	}
}

func (r *Renderer) Render(screen *Screen) {
	for _, renderable := range r.renderables {
		renderable.Render(screen)
	}
}

func (r *Renderer) Register(renderable Renderable) {
	order := len(r.renderables)
	r.renderables[order] = renderable
}

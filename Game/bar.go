package Game

import (
	"github.com/vova616/garageEngine/engine"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Bar struct {
	engine.BaseComponent
	value float32
	width float32
}

func NewBar(width float32) *Bar {
	return &Bar{engine.NewComponent(), 0, width}
}
func (s *Bar) Start() {
}
func (s *Bar) GetValue() float32 {
	return s.value
}
func (s *Bar) SetValue(val float32) {
	s.value = val
	sc := s.GameObject().Transform().Scale()
	sc.X = s.width * s.value
	s.GameObject().Transform().SetScale(sc)
}

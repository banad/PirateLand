package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"math"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Enemy struct {
	Engine.BaseComponent

	frames    int
	Hp        float32
	MaxHp     float32
	HpB       *Bar
	OnGround  bool
	Attack    bool
	able      bool
	jump      bool
	hit       bool
	speed     float32
	width     float32
	height    float32
	LastFloor *Engine.GameObject
}

func NewEnemy(Hp *Bar) *Enemy {
	return &Enemy{Engine.NewComponent(), 0, 0, 100, Hp, false, false, true, false, false, 60, 0, 0, nil}

}
func (s *Enemy) Start() {
	s.GameObject().Sprite.SetAnimation("enemy_jump")

	s.GameObject().Physics.Body.SetMoment(Engine.Inf)
	s.width = s.Transform().WorldScale().X
	s.height = s.Transform().WorldScale().Y
}
func (s *Enemy) Update() {
	ph := s.GameObject().Physics.Body
	s.GameObject().Sprite.SetAlign(Engine.AlignTopCenter)
	if float32(math.Abs(float64(ph.Velocity().X))) > 3 {
		if s.GameObject().Sprite.CurrentAnimation() == "enemy_stand" {
			s.GameObject().Sprite.SetAnimation("enemy_walk")
		}
	} else if !s.Attack {
		s.GameObject().Sprite.SetAnimation("enemy_stand")
	}
	if s.OnGround == false {
		s.frames++
	} else {
		s.frames = 0
	}
	d := s.Transform().WorldPosition()
	s.HpB.Transform().SetWorldPosition(d.Add(Engine.NewVector2(0, 30)))
	if s.able {
		if plComp.Transform().WorldPosition().X > s.Transform().WorldPosition().X {
			ph.AddForce(s.speed, 0)
			s.GameObject().Transform().SetScalef(s.width, s.height)

		} else {
			ph.AddForce(-s.speed, 0)
			s.GameObject().Transform().SetScalef(-s.width, s.height)
		}
		d = plComp.Transform().WorldPosition()
		if d.Distance(s.Transform().WorldPosition()) < 50 {
			s.Attack = true

			s.GameObject().Sprite.SetAnimation("enemy_attack")
			s.GameObject().Sprite.AnimationEndCallback = func(sprite *Engine.Sprite) {
				s.Attack = false
				s.GameObject().Sprite.SetAnimation("enemy_stand")
			}

		}
		if s.jump && s.OnGround {
			s.GameObject().Physics.Body.AddForce(0, 100)
		}
	}
	if s.hit {
		s.GameObject().Sprite.SetAnimation("enemy_hit")
	}
	if (s.GameObject().Sprite.CurrentAnimation() == "enemy_stand" || s.GameObject().Sprite.CurrentAnimation() == "enemy_walk") && !s.OnGround && s.frames > 15 {
		s.GameObject().Sprite.SetAnimation("enemy_jump")
	}

}
func (s *Enemy) OnCollisionPostSolve(arbiter Engine.Arbiter) {
	if arbiter.GameObjectB().Tag != "lader" && arbiter.GameObjectA().Tag != "lader" {
		count := 0
		for _, con := range arbiter.Contacts {
			if arbiter.Normal(con).Y < -0.9 {
				count++

			}
		}
		if count >= 1 {
			if s.GameObject().Sprite.CurrentAnimation() == "enemy_jump" {
				s.GameObject().Sprite.SetAnimation("enemy_stand")
			}
			s.LastFloor = arbiter.GameObjectB()

			s.OnGround = true
		}
		count = 0
		for _, con := range arbiter.Contacts {
			if math.Abs(float64(arbiter.Normal(con).X)) > 0.9 {
				count++

			}
		}
		if count >= 1 {
			s.jump = true
		}
		if s.Attack && arbiter.GameObjectB().Tag == "player" {
			if plComp.hitable {
				plComp.Hit()
			}
		}
	}

}
func (s *Enemy) FixedUpdate() {
	s.OnGround = false
	s.LastFloor = nil
	s.jump = false
}

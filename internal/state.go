package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type State struct {
	Player      Player
	Enemies     EnemyList
	Projectiles ProjectileList
}

type Afterimage = [20]rl.Vector2
type Player struct {
	Alive       bool
	Position    rl.Vector2
	Velocity    rl.Vector2
	Radius      float32
	ZPos        float32
	ZVel        float32
	Afterimages Afterimage
}

type EnemyList = [500]Enemy
type Enemy struct {
	Alive       bool
	Position    rl.Vector2
	Radius      float32
	Personality float32
}

type ProjectileList = [2000]Projectile
type Projectile struct {
	Alive    bool
	Position rl.Vector2
	Velocity rl.Vector2
	Hostile  bool
	Radius   float32
}

func DefaultState() State {
	return State{
		Player: Player{
			Alive:    true,
			Position: rl.Vector2{X: 0, Y: 0},
			Radius:   64,
		},
		Enemies:     EnemyList{},
		Projectiles: ProjectileList{},
	}
}

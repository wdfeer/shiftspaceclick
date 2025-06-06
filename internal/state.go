package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type State struct {
	Player      Player
	Enemies     EnemyList
	Projectiles ProjectileList
}

type Player struct {
	Alive    bool
	Position rl.Vector2
	Velocity rl.Vector2
	Size     float32
}

type EnemyList = [99]Enemy
type Enemy struct {
	Alive    bool
	Position rl.Vector2
	Size     float32
}

type ProjectileList = [999]Projectile
type Projectile struct {
	Alive    bool
	Position rl.Vector2
	Velocity rl.Vector2
	Hostile  bool
	Size     float32
}

func DefaultState() State {
	return State{
		Player: Player{
			Alive:    true,
			Position: rl.Vector2{X: 0, Y: 0},
			Size:     64,
		},
		Enemies:     EnemyList{},
		Projectiles: ProjectileList{},
	}
}

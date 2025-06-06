package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type State struct {
	Player      Player
	Enemies     EnemyList
	Projectiles ProjectileList
}

type Player struct {
	Position rl.Vector2
	Velocity rl.Vector2
}

type EnemyList = [99]Enemy
type Enemy struct {
	Alive    bool
	Position rl.Vector2
}

type ProjectileList = [999]Projectile
type Projectile struct {
	Alive    bool
	Position rl.Vector2
	Velocity rl.Vector2
	Hostile  bool
}

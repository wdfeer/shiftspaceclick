package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (state State) Render() {
	rl.ClearBackground(rl.Black)

	state.Player.render()

	for _, e := range state.Enemies {
		e.render()
	}

	for _, p := range state.Projectiles {
		p.render()
	}
}

func (player Player) render() {
	rl.DrawCircleV(player.Position, player.Size/2, rl.RayWhite)
}

func (enemy Enemy) render() {
	rl.DrawCircleV(enemy.Position, enemy.Size/2, rl.Maroon)
}

func (projectile Projectile) render() {
	var color rl.Color
	if projectile.Hostile {
		color = rl.Red
	} else {
		color = rl.White
	}
	rl.DrawCircleV(projectile.Position, projectile.Size/2, color)
}

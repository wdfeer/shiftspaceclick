package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (state State) Render() {
	camera := rl.Camera2D{}
	camera.Target = state.Player.Position

	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(camera)

	if state.Player.Alive {
		state.Player.render()
	}

	for _, e := range state.Enemies {
		if e.Alive {
			e.render()
		}
	}

	for _, p := range state.Projectiles {
		if p.Alive {
			p.render()
		}
	}

	rl.EndMode2D()
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

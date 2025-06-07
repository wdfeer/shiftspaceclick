package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (state State) Render() {
	camera := rl.Camera2D{}
	camera.Target = state.Player.Position
	camera.Offset = rl.Vector2{X: 800, Y: 450}
	camera.Zoom = 0.5

	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(camera)

	for i := -25; i <= 25; i++ {
		spacing := 1000 * int32(i)
		rl.DrawLine(-9999, spacing, 9999, spacing, rl.RayWhite)
		rl.DrawLine(spacing, -9999, spacing, 9999, rl.RayWhite)
	}

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

	state.renderUI()
}

func (player Player) render() {
	shadowOffset := rl.Vector2{X: 3, Y: 3}
	rl.DrawCircleV(rl.Vector2Add(player.Position, shadowOffset), player.Radius, rl.ColorTint(rl.Gray, rl.RayWhite))
	rl.DrawCircleV(player.Position, player.Radius, rl.RayWhite)
}

func (enemy Enemy) render() {
	rl.DrawCircleV(rl.Vector2Add(enemy.Position, rl.Vector2{X: 3, Y: 3}), enemy.Radius, rl.ColorTint(rl.Gray, rl.Maroon))
	rl.DrawCircleV(enemy.Position, enemy.Radius, rl.Maroon)
}

func (projectile Projectile) render() {
	var color rl.Color
	if projectile.Hostile {
		color = rl.Red
	} else {
		color = rl.White
	}
	rl.DrawCircleV(rl.Vector2Add(projectile.Position, rl.Vector2{X: 2, Y: 2}), projectile.Radius, rl.ColorTint(rl.Gray, color))
	rl.DrawCircleV(projectile.Position, projectile.Radius, color)
}

func (state State) renderUI() {
	if !state.Player.Alive {
		rl.DrawText("You died!", 10, 10, 20, rl.White)
	}
}

package internal

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (state State) Render() {
	camera := rl.Camera2D{}
	camera.Target = state.Player.Position
	camera.Offset = rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
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

func getShadowOffset(distance float32) rl.Vector2 {
	return rl.Vector2Scale(rl.Vector2Rotate(rl.Vector2{X: 1, Y: 0}, math.Pi/6), distance)
}

func (player Player) render() {
	for i, p := range player.Afterimage {
		if p != rl.Vector2Zero() {
			opacity := float32(math.Pow(1-float64(i)/float64(len(player.Afterimage)), 3))
			rl.DrawCircleV(p, player.Radius, rl.ColorAlpha(rl.RayWhite, opacity))
		}
	}

	shadowOffset := getShadowOffset(max(1, player.ZPos*50))
	rl.DrawCircleV(rl.Vector2Add(player.Position, shadowOffset), player.Radius, rl.ColorTint(rl.Gray, rl.RayWhite))
	rl.DrawCircleV(player.Position, player.Radius, rl.RayWhite)
}

func (enemy Enemy) render() {
	for i, p := range enemy.Afterimage {
		if p != rl.Vector2Zero() {
			opacity := float32(math.Pow(1-float64(i)/float64(len(enemy.Afterimage)), 3))
			rl.DrawCircleV(p, enemy.Radius, rl.ColorAlpha(rl.Maroon, opacity))
		}
	}

	rl.DrawCircleV(rl.Vector2Add(enemy.Position, getShadowOffset(5)), enemy.Radius, rl.ColorTint(rl.Gray, rl.Maroon))
	rl.DrawCircleV(enemy.Position, enemy.Radius, rl.Maroon)
}

func (projectile Projectile) render() {
	var color rl.Color
	var shadowOffset float32 = 3
	if projectile.Hostile {
		color = rl.Red
	} else {
		color = rl.White
		if projectile.Radius == 24 {
			shadowOffset = 8
		}
	}

	for i, p := range projectile.Afterimage {
		if p != rl.Vector2Zero() {
			opacity := float32(math.Pow(1-float64(i)/float64(len(projectile.Afterimage)), 3))
			rl.DrawCircleV(p, projectile.Radius, rl.ColorAlpha(color, opacity))
		}
	}

	rl.DrawCircleV(rl.Vector2Add(projectile.Position, getShadowOffset(shadowOffset)), projectile.Radius, rl.ColorTint(rl.Gray, color))
	rl.DrawCircleV(projectile.Position, projectile.Radius, color)
}

func (state State) renderUI() {
	if state.Player.Position.Y == 0 && state.Player.Position.X == 0 {
		rl.DrawText("Movement: WASD", 660, 600, 40, rl.Yellow)
		rl.DrawText("Dash: LSHIFT", 760, 650, 40, rl.Yellow)
		rl.DrawText("Jump: SPACE", 860, 700, 40, rl.Yellow)
		rl.DrawText("Shoot: LCLICK", 960, 750, 40, rl.Yellow)
		rl.DrawText("Rocket: Jump + Shoot", 1060, 800, 40, rl.Yellow)
	}
	rl.DrawText(fmt.Sprintf("Speed: %d", int32(rl.Vector2Length(state.Player.Velocity))), 10, 10, 40, rl.White)
}

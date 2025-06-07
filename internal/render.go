package internal

import (
	"fmt"
	"math"

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

func getShadowOffset(distance float32) rl.Vector2 {
	return rl.Vector2Scale(rl.Vector2Rotate(rl.Vector2{X: 1, Y: 0}, math.Pi/6), distance)
}

func (player Player) render() {
	shadowOffset := getShadowOffset(max(1, player.ZPos*50))
	rl.DrawCircleV(rl.Vector2Add(player.Position, shadowOffset), player.Radius, rl.ColorTint(rl.Gray, rl.RayWhite))
	rl.DrawCircleV(player.Position, player.Radius, rl.RayWhite)
}

func (enemy Enemy) render() {
	rl.DrawCircleV(rl.Vector2Add(enemy.Position, getShadowOffset(5)), enemy.Radius, rl.ColorTint(rl.Gray, rl.Maroon))
	rl.DrawCircleV(enemy.Position, enemy.Radius, rl.Maroon)
}

func (projectile Projectile) render() {
	var color rl.Color
	if projectile.Hostile {
		color = rl.Red
	} else {
		color = rl.White
	}
	rl.DrawCircleV(rl.Vector2Add(projectile.Position, getShadowOffset(3)), projectile.Radius, rl.ColorTint(rl.Gray, color))
	rl.DrawCircleV(projectile.Position, projectile.Radius, color)
}

func (state State) renderUI() {
	if !state.Player.Alive {
		rl.DrawText("You died!", 10, 10, 50, rl.White)
	} else {
		if state.Player.Position.Y == 0 && state.Player.Position.X == 0 {
			rl.DrawText("Movement: WASD", 660, 600, 40, rl.Yellow)
			rl.DrawText("Dash: LSHIFT", 760, 650, 40, rl.Yellow)
			rl.DrawText("Jump: SPACE", 860, 700, 40, rl.Yellow)
			rl.DrawText("Shoot: LCLICK", 960, 750, 40, rl.Yellow)
		}
		rl.DrawText(fmt.Sprintf("Speed: %d", int32(rl.Vector2Length(state.Player.Velocity))), 10, 10, 40, rl.White)
	}
}

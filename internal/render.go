package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (state State) Render() {
	camera := rl.Camera2D{}
	camera.Target = state.Player.Position
	camera.Offset = rl.Vector2{X: 800, Y: 450}
	camera.Zoom = 0.5

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

	state.renderUI()
}

func (player Player) render() {
	rl.DrawCircleV(player.Position, player.Radius, rl.RayWhite)
}

func (enemy Enemy) render() {
	rl.DrawCircleV(enemy.Position, enemy.Radius, rl.Maroon)
}

func (projectile Projectile) render() {
	var color rl.Color
	if projectile.Hostile {
		color = rl.Red
	} else {
		color = rl.White
	}
	rl.DrawCircleV(projectile.Position, projectile.Radius, color)
}

func (state State) renderUI() {
	if !state.Player.Alive {
		rl.DrawText("You died!", 10, 10, 20, rl.White)
	}
}

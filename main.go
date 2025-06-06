package main

import (
	"vektorrush/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "vektorrush")
	defer rl.CloseWindow()

	rl.SetTargetFPS(1000)

	state := internal.State{
		Player: internal.Player{
			Alive:    true,
			Position: rl.Vector2{X: 0, Y: 0},
			Size:     64,
		},
	}
	for !rl.WindowShouldClose() {
		state = state.Update()
		rl.BeginDrawing()
		state.Render()
		rl.EndDrawing()
	}
}

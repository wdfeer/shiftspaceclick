package main

import (
	"vektorrush/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "vektorrush")
	defer rl.CloseWindow()

	rl.SetTargetFPS(500)

	rl.SetWindowState(rl.FlagWindowResizable + rl.FlagBorderlessWindowedMode)

	state := internal.DefaultState()
	for !rl.WindowShouldClose() {
		state = state.Update()
		rl.BeginDrawing()
		state.Render()
		rl.EndDrawing()
	}
}

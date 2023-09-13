package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetTraceLogCallback(func(msgType int, text string) {})
	rl.InitWindow(800, 450, "Warp Riders")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Warp Riders Time", 190, 200, 20, rl.Black)

		rl.EndDrawing()
	}
}

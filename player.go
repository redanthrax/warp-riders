package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position  rl.Vector2
	Sprite    rl.Texture2D
	Frame     rl.Rectangle
	IsWalking bool
}

func LoadPlayer(g *Game) {
	g.Player.Position = rl.NewVector2(0, 0)
	g.Player.Frame = rl.NewRectangle(0, 420, spriteSize, spriteSize)
}

func DrawPlayer(g *Game) {
	rl.DrawText(fmt.Sprintf("X: %f, Y: %f", g.Player.Position.X, g.Player.Position.Y), 0, screenHeight-20, 10, rl.Black)
	rl.DrawText(fmt.Sprintf("Walking: %v", g.Player.IsWalking), 0, screenHeight-40, 10, rl.Black)
	dst := rl.NewRectangle(g.Player.Position.X, g.Player.Position.Y, spriteSize*2, spriteSize*2)
	rl.DrawTexturePro(g.Player.Sprite,
		g.Player.Frame, dst, rl.Vector2Zero(), 0, rl.White)
}

func AnimatePlayer(g *Game) {
	g.FramesCounter++
	if g.Player.IsWalking {
		if g.CurrentFrame > 3 {
			g.CurrentFrame = 0
		}

		if g.FramesCounter > g.FramesSpeed {
			g.FramesCounter = 0
		}

		g.Player.Frame.X = float32(g.CurrentFrame) * spriteSize
	}
}

func PlayerControls(g *Game) {
	if rl.IsKeyDown(rl.KeyQ) {
		g.WindowShouldClose = true
	}

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		g.Player.IsWalking = true
		g.Player.Position.X--
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		g.Player.IsWalking = true
		g.Player.Position.X++
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		g.Player.IsWalking = true
		g.Player.Position.Y--
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		g.Player.IsWalking = true
		g.Player.Position.Y++
	}
}

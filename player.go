package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position     rl.Vector2
	Sprite       rl.Texture2D
	Frame        rl.Rectangle
	IsWalking    bool
	LastPosition rl.Vector2
}

func LoadPlayer(g *Game) {
	g.Player.Position = rl.NewVector2(0, 0)
	g.Player.Frame = rl.NewRectangle(0, 420, spriteSize, spriteSize)
}

func DrawPlayer(g *Game) {
	dst := rl.NewRectangle(g.Player.Position.X, g.Player.Position.Y, spriteSize*3, spriteSize*3)
	rl.DrawTexturePro(g.Player.Sprite, g.Player.Frame, dst, rl.Vector2Zero(), 0,
		rl.White)
}

func AnimatePlayer(g *Game) {
	if g.Player.IsWalking {
		g.FramesCounter++
		g.Player.Frame.Y = 420
		if g.CurrentFrame > 4 {
			g.CurrentFrame = 1
		}

		if g.FramesCounter > g.FramesSpeed {
			g.FramesCounter = 1
		}

		g.Player.Frame.X = (float32(g.CurrentFrame) - 1) * spriteSize
		//4 walking frames
		wf := g.FramesSpeed / 4
		if wf*g.CurrentFrame <= g.FramesCounter {
			g.CurrentFrame++
		}

		g.Player.LastPosition = g.Player.Position

	} else {
		g.Player.Frame.Y = 132
		g.Player.Frame.X = 0
		g.FramesCounter = 1
		g.CurrentFrame = 0
	}

	if g.Player.LastPosition == g.Player.Position {
		g.Player.IsWalking = false
	}
}

func DrawOnPlayer(g *Game, text string) {
	rl.DrawText(text, int32(g.Player.Position.X), int32(g.Player.Position.Y), 10, rl.Black)
}

func PlayerControls(g *Game) {
	if rl.IsKeyDown(rl.KeyQ) {
		g.WindowShouldClose = true
	}

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		g.Player.IsWalking = true
		g.Player.Position.X--
		g.Player.Frame.Width = -spriteSize
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		g.Player.IsWalking = true
		g.Player.Position.X++
		g.Player.Frame.Width = spriteSize
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

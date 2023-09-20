package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadPlayer(g *Game) {
	g.Player.Position = rl.NewVector2(0, 0)
	g.Player.FrameRec = rl.NewRectangle(0, 420, spriteSize, spriteSize)
}

func DrawPlayer(g *Game) {
	//rl.DrawTextureRec(g.Player.Sprite, g.Player.FrameRec, g.Player.Position, rl.RayWhite)
	//DrawTexturePro(texture, { 0.0f, 0.0f, texture.width, texture.height },
	//{ screenWidth / 2.0f, screenHeight / 2.0f, texture.width, texture.height }, {texture.width / 2, texture.height / 2}, 0.0f, WHITE);

	posX := screenWidth / 2
	posY := screenHeight / 2
	dst := rl.NewRectangle(float32(posX), float32(posY), g.Player.FrameRec.Width*2, g.Player.FrameRec.Height*2)
	org := rl.NewVector2(float32(g.Player.Sprite.Width/2), float32(g.Player.Sprite.Height/2))
	rl.DrawTexturePro(g.Player.Sprite, g.Player.FrameRec, dst, org, 0, rl.White)

	rl.DrawText(fmt.Sprintf("X: %f, Y: %f", g.Player.Position.X, g.Player.Position.Y), 0, screenHeight-20, 10, rl.Black)
}

func PlayerControls(g *Game) {
	if rl.IsKeyDown(rl.KeyQ) {
		g.WindowShouldClose = true
	}

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		g.Player.Position.X--
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		g.Player.Position.X++
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		g.Player.Position.Y--
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {

		g.Player.Position.Y++
	}
}

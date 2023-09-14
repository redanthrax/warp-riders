package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	spriteSize   = 32
)

type Player struct {
	Position rl.Vector2
	Sprite   rl.Texture2D
	FrameRec rl.Rectangle
}

type Game struct {
	GameOver bool
	Dead     bool
	Pause    bool

	FramesCounter int32
	FramesSpeed   int32
	CurrentFrame  int32

	WindowShouldClose bool

	Player Player
}

func NewGame() (g Game) {
	g.Init()
	return
}

func main() {
	game := NewGame()
	game.GameOver = true

	rl.InitWindow(screenWidth, screenHeight, "Warp Riders")

	rl.InitAudioDevice()

	game.Load()

	rl.SetTargetFPS(60)

	for !game.WindowShouldClose {
		game.Update()
		game.Draw()
	}

	game.Unload()
	rl.CloseAudioDevice()
	rl.CloseWindow()
	os.Exit(0)
}

func (g *Game) Init() {
	g.Player = Player{
		Position: rl.NewVector2(0, 0),
		FrameRec: rl.NewRectangle(0, 0, spriteSize, spriteSize),
	}

	g.FramesCounter = 0
	g.WindowShouldClose = false
	g.GameOver = false
	g.Pause = false
}

func (g *Game) Load() {
	g.Player.Sprite = rl.LoadTexture("assets/character/walk/walk.png")
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Player.Sprite)
}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}

	if !g.GameOver {
		if rl.IsKeyPressed(rl.KeyQ) {
			g.WindowShouldClose = true
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			g.Player.Position.X--
		}

		if rl.IsKeyDown(rl.KeyRight) {
			g.Player.Position.X++
		}

		if !g.Pause {
			//g.FramesCounter++
			//if g.FramesCounter >= 8 {
			//	g.FramesCounter = 0
			//	g.Player.FrameRec.X = spriteSize
			//} else {
			//	g.Player.FrameRec.X = 0
			//}
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.Init()
		}
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	if !g.GameOver {
		rl.DrawTextureRec(g.Player.Sprite, g.Player.FrameRec, g.Player.Position, rl.RayWhite)
		rl.DrawText("LETS GO", 0, 0, 10, rl.Black)
	} else {
		rl.DrawTextureRec(g.Player.Sprite, g.Player.FrameRec, g.Player.Position, rl.RayWhite)
		rl.DrawText("Enter to start", 0, 0, 10, rl.Black)
	}

	rl.EndDrawing()
}

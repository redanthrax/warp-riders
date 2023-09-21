package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 768
	spriteSize   = 32
	gameTitle    = "Warp Riders"
)

func NewGame() (g Game) {
	g.Init()
	return
}

func main() {
	game := NewGame()
	game.MainMenu = true
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.InitAudioDevice()
	game.Load()
	monitor := rl.GetCurrentMonitor()
	refreshRate := rl.GetMonitorRefreshRate(monitor)
	game.FramesSpeed = int32(refreshRate)
	rl.SetTargetFPS(int32(refreshRate))
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
	g.Player = Player{}
	g.MusicTracks = map[string]MusicTrack{}
	g.Sounds = make(map[string]rl.Sound)
	g.FramesCounter = 0
	g.WindowShouldClose = false
	g.MainMenu = false
	g.Pause = false
	g.Menu = Menu{
		Selected: 0,
		Options:  []string{"New Game", "Settings", "Exit"},
	}
}

func (g *Game) Load() {
	g.Player.Sprite = rl.LoadTexture("assets/characters/buddie0 sprite sheet x1.png")
	LoadMusic(g)
	LoadSounds(g)
	LoadPlayer(g)
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Player.Sprite)
	rl.UnloadMusicStream(g.MusicTracks[g.CurrentTrack].music)
	rl.CloseAudioDevice()
}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}

	UpdateMusic(g)
	if !g.MainMenu {
		if !g.Pause {
			PlayerControls(g)
			AnimatePlayer(g)
		}
	} else {
		g.CurrentTrack = "menu"
		MenuOptions(g)
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)
	rl.DrawText(fmt.Sprintf("Frames: %d", rl.GetFPS()), 0, 0, 10, rl.Black)
	if !g.MainMenu {
		DrawPlayer(g)
	} else {
		g.DrawMenu()
	}

	rl.EndDrawing()
}

func (g *Game) DrawMenu() {
	//draw title
	titleWidth := rl.MeasureText(gameTitle, 30)
	titlePosX := (rl.GetScreenWidth() / 2) - (int(titleWidth) / 2)
	rl.DrawText(gameTitle, int32(titlePosX), 100, 30, rl.Black)

	//draw options
	for i := 0; i < len(g.Menu.Options); i++ {
		optionText := g.Menu.Options[i]
		if g.Menu.Selected == i {
			optionText = fmt.Sprintf(">%s<", g.Menu.Options[i])
		}

		optWidth := rl.MeasureText(optionText, 20)
		optPosX := (rl.GetScreenWidth() / 2) - (int(optWidth) / 2)
		rl.DrawText(optionText, int32(optPosX), int32(150+(i*40)), 20, rl.Red)
	}
}

package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func MenuOptions(g *Game) {
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		g.Menu.Selected--
		if g.Menu.Selected < 0 {
			g.Menu.Selected = len(g.Menu.Options) - 1
		}

		PlaySound(g, "menu")
	}

	if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		g.Menu.Selected++
		if g.Menu.Selected > (len(g.Menu.Options) - 1) {
			g.Menu.Selected = 0
		}

		PlaySound(g, "menu")
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		switch g.Menu.Options[g.Menu.Selected] {
		case "New Game":
			g.MainMenu = false
			PlaySound(g, "confirm")
		case "Exit":
			g.WindowShouldClose = true
		}
	}
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

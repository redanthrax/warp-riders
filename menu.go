package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

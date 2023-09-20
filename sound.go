package main

import rl "github.com/gen2brain/raylib-go/raylib"

func LoadMusic(g *Game) {
	menuMusic := MusicTrack{
		music: rl.LoadMusicStream("assets/music/Race To Mars.mp3"),
	}

	g.MusicTracks["menu"] = menuMusic
}

func LoadSounds(g *Game) {
	g.Sounds["menu"] = rl.LoadSound("assets/sounds/Bleep_03.wav")
	g.Sounds["confirm"] = rl.LoadSound("assets/sounds/Confirm_06.wav")
}

func UpdateMusic(g *Game) {
	rl.UpdateMusicStream(g.MusicTracks[g.CurrentTrack].music)
	if !rl.IsMusicStreamPlaying(g.MusicTracks[g.CurrentTrack].music) {
		rl.PlayMusicStream(g.MusicTracks[g.CurrentTrack].music)
	}
}

func PlaySound(g *Game, sound string) {
	rl.PlaySound(g.Sounds[sound])
}

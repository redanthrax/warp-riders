package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	MainMenu          bool
	Pause             bool
	FramesCounter     int32
	FramesSpeed       int32
	CurrentFrame      int32
	WindowShouldClose bool
	Player            Player
	MusicTracks       map[string]MusicTrack
	CurrentTrack      string
	Menu              Menu
	Sounds            map[string]rl.Sound
}

type MusicList struct {
	Tracks []MusicTrack
}

type MusicTrack struct {
	music     rl.Music
	label     string
	volume    float32
	maxVolume float64 `default:"0.5"`
}

type Player struct {
	Position rl.Vector2
	Sprite   rl.Texture2D
	FrameRec rl.Rectangle
}

type Menu struct {
	Selected int
	Options  []string
}

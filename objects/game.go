package objects

import (
	"somegame/engine"
)

var level *Level

func onNewFrame() {
	if level != nil {
		level.OnNewFrame()
	}
}

func onReady() {
	level = CreateLevel()
}

func CreateGame() {
	engine.OpenWindow(onReady, onNewFrame)
}

func loadLevel(l *Level) {
	level = l
}

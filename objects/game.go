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
	level = CreateLevel([]GameObject{CreatePlayer(triangle, 0.1)})
}

func CreateGame() {
	engine.OpenWindow(onReady, onNewFrame)
}

func loadLevel(l *Level) {
	level = l
}

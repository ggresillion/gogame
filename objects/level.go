package objects

import (
	"somegame/common"
)

var (
	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
)

// Level - represents a loaded game level
type Level struct {
	GameObjects []common.GameObject
}

// CreateLevel - level constructor
func CreateLevel() *Level {
	var level = &Level{[]common.GameObject{}}
	level.GameObjects = append(level.GameObjects, CreatePlayer(level, triangle, 1))
	return level
}

// RegisterGameObject - register a new game object in this level instance
func (l Level) RegisterGameObject(o common.GameObject) {
	l.GameObjects = append(l.GameObjects, o)
}

func renderGameObject(o common.GameObject) {
}

func (l Level) OnNewFrame() {
	for _, obj := range l.GameObjects {
		obj.OnNewFrame()
	}
}

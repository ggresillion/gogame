package objects

import (
	"somegame/engine"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Movement - current movement of the player
type Movement struct {
	X int
	Y int
}

// Player - player object
type Player struct {
	Vertices []float32
	Movement Movement
	Speed    float32
}

// OnNewFrame - triggered on every new frame
func (p Player) OnNewFrame() {
	p.Move(float32(p.Movement.X)*p.Speed, float32(p.Movement.Y)*p.Speed)
	engine.RenderGameObject(p.Vertices)
}

// CreatePlayer - instantiate a new Player
func CreatePlayer(vertices []float32, size float32) Player {
	for i, point := range vertices {
		vertices[i] = point * size
	}
	player := Player{vertices, Movement{0, 0}, 0.1}
	engine.RegisterOnKeyCallback(player.KeyCallback)
	return player
}

// Move - move a Player to the given directions
func (p Player) Move(x float32, y float32) {
	for i := 0; i < len(p.Vertices); i = i + 3 {
		p.Vertices[i] = p.Vertices[i] + x
		p.Vertices[i+1] = p.Vertices[i+1] + y
	}
}

// KeyCallback - callback triggered for every key press
func (p Player) KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyW && action == glfw.Press {
		p.Movement.X = 1
	}
	if key == glfw.KeyS && action == glfw.Press {
		p.Movement.X = -1
	}
	if key == glfw.KeyD && action == glfw.Press {
		p.Movement.Y = 1
	}
	if key == glfw.KeyA && action == glfw.Press {
		p.Movement.Y = -1
	}
	if key == glfw.KeyW && action == glfw.Release {
		p.Movement.X = 0
	}
	if key == glfw.KeyS && action == glfw.Release {
		p.Movement.X = 0
	}
	if key == glfw.KeyD && action == glfw.Release {
		p.Movement.Y = 0
	}
	if key == glfw.KeyA && action == glfw.Release {
		p.Movement.Y = 0
	}
}

package objects

import (
	"somegame/common"
	"somegame/engine"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Player - player object
type Player struct {
	level    *Level
	vertices []float32
	position common.Coordinates
	movement common.Coordinates
	speed    float32
	size     float32
}

func (p *Player) GetVertices() []float32 {
	return p.vertices
}

func (p *Player) GetPosition() common.Coordinates {
	return p.position
}

func (p *Player) GetSize() float32 {
	return p.size
}

// OnNewFrame - triggered on every new frame
func (p *Player) OnNewFrame() {
	p.Move(float32(p.movement.X)*p.speed, float32(p.movement.Y)*p.speed)
	engine.RenderGameObject(p)
}

// CreatePlayer - instantiate a new Player
func CreatePlayer(level *Level, vertices []float32, size float32) *Player {
	for i, point := range vertices {
		vertices[i] = point * size
	}
	player := Player{level, vertices, common.Coordinates{0.5, 0.5}, common.Coordinates{0, 0}, 0.01, 0.1}
	engine.RegisterOnKeyCallback(player.KeyCallback)
	return &player
}

// Move - move a Player to the given directions
func (p *Player) Move(x float32, y float32) {
	p.position.X = p.position.X + p.movement.X*p.speed
	p.position.Y = p.position.Y + p.movement.Y*p.speed
}

func (p *Player) Shoot() {

}

// KeyCallback - callback triggered for every key press
func (p *Player) KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyD && action == glfw.Press {
		p.movement.X = 1
	}
	if key == glfw.KeyA && action == glfw.Press {
		p.movement.X = -1
	}
	if key == glfw.KeyW && action == glfw.Press {
		p.movement.Y = 1
	}
	if key == glfw.KeyS && action == glfw.Press {
		p.movement.Y = -1
	}
	if key == glfw.KeyD && action == glfw.Release {
		p.movement.X = 0
	}
	if key == glfw.KeyA && action == glfw.Release {
		p.movement.X = 0
	}
	if key == glfw.KeyW && action == glfw.Release {
		p.movement.Y = 0
	}
	if key == glfw.KeyS && action == glfw.Release {
		p.movement.Y = 0
	}
}

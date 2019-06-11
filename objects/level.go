package objects

var (
	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
)

// Level - represents a loaded game level
type Level struct {
	GameObjects []GameObject
}

// CreateLevel - level constructor
func CreateLevel(objects []GameObject) *Level {
	return &Level{objects}
}

// RegisterGameObject - register a new game object in this level instance
func (l Level) RegisterGameObject(o GameObject) {
	l.GameObjects = append(l.GameObjects, o)
}

func renderGameObject(o GameObject) {
}

func (l Level) OnNewFrame() {
	for _, obj := range l.GameObjects {
		obj.OnNewFrame()
	}
}

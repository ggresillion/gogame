package objects

// GameObject - basic game object interface
type GameObject interface {
	GetVertices() []float32
	GetPosition() Coordinates
	OnNewFrame()
}

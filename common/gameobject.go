package common

// GameObject - basic game object interface
type GameObject interface {
	GetVertices() []float32
	GetPosition() Coordinates
	GetSize() float32
	OnNewFrame()
}

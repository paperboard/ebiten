package ebiten

type Camera3D struct {
	Fov  float64
	Near float64
	Far  float64
	X    float64
	Y    float64
	Z    float64
}

func NewCamera3D(fov float64, near float64, far float64) *Camera3D {
	return &Camera3D{
		Fov:  fov,
		Near: near,
		Far:  far,
		X:    0,
		Y:    0,
		Z:    0,
	}
}

func (c *Camera3D) Position(x float64, y float64, z float64) {
	c.X = x
	c.Y = y
	c.Z = z
}

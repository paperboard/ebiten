package ebiten

type Cuboid3D struct {
	Textures []*Image
	Length   int
	Width    int
	Height   int
	X        float64
	Y        float64
	Z        float64
}

func NewCuboid3D(length int, width int, height int) *Cuboid3D {
	return &Cuboid3D{
		Textures: []*Image{},
		Length:   length,
		Width:    width,
		Height:   height,
		X:        0,
		Y:        0,
		Z:        0,
	}
}

func (c *Cuboid3D) Position(x float64, y float64, z float64) {
	c.X = x
	c.Y = y
	c.Z = z
}

func (c *Cuboid3D) Faces(front *Image, back *Image, top *Image, bottom *Image, left *Image, right *Image) {
	c.Textures = append(c.Textures, front, back, top, bottom, left, right)
}

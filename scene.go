package ebiten

type Scene3D struct {
	Cuboids []*Cuboid3D
	Camera  *Camera3D
}

func NewScene3D(camera3d *Camera3D) *Scene3D {
	return &Scene3D{
		Cuboids: []*Cuboid3D{},
		Camera:  camera3d,
	}
}

func (s *Scene3D) Add(cuboid *Cuboid3D) {
	s.Cuboids = append(s.Cuboids, cuboid)
}

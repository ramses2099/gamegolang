package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	maxAngle     = 256
)

//interface Game Object
type GameObject interface {
	Update() error
	Draw(screen *ebiten.Image)
}

//struct Game Entity
type GameEntity struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	vx          int
	vy          int
	angle       int
	image       *ebiten.Image
}

func NewGameEntity(x, y, vx, vy, angle int, image *ebiten.Image) GameObject {
	w, h := image.Size()
	return &GameEntity{
		imageWidth:  w,
		imageHeight: h,
		x:           x,
		y:           y,
		vx:          vx,
		vy:          vy,
		angle:       angle,
		image:       image}
}

//function draw to screen
func (ge *GameEntity) Draw(screen *ebiten.Image) {
	w, h := ge.image.Size()
	op := ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(2 * math.Pi * float64(ge.angle) / maxAngle)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(float64(ge.x), float64(ge.y))
	screen.DrawImage(ge.image, &op)
}

//function update
func (ge *GameEntity) Update() error {
	ge.x += ge.vx
	ge.y += ge.vy
	if ge.x < 0 {
		ge.x = -ge.x
		ge.vx = -ge.vx
	} else if mx := screenWidth - ge.imageWidth; mx <= ge.x {
		ge.x = 2*mx - ge.x
		ge.vx = -ge.vx
	}
	if ge.y < 0 {
		ge.y = -ge.y
		ge.vy = -ge.vy
	} else if my := screenHeight - ge.imageHeight; my <= ge.y {
		ge.y = 2*my - ge.y
		ge.vy = -ge.vy
	}
	ge.angle++
	if ge.angle == maxAngle {
		ge.angle = 0
	}
	return nil
}

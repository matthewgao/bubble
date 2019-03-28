package shape

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Box struct {
	BH    int
	BW    int
	H     int
	W     int
	X     int
	Y     int
	Yb    int
	Xb    int
	Xo    int
	Yo    int
	color color.RGBA
}

func NewBox(bh, bw int, h, w int) *Box {
	return &Box{
		BH:    bh,
		BW:    bw,
		H:     h,
		W:     w,
		Xb:    int(rand.Int63n(5)),
		Yb:    int(rand.Int63n(5)),
		Xo:    1,
		Yo:    -1,
		X:     int(rand.Int63n(int64(bw))),
		Y:     int(rand.Int63n(int64(bh))),
		color: color.RGBA{100, 100, 100, 255},
	}
}

func NewBoxFix(bh, bw int, h, w int, x, y int) *Box {
	return &Box{
		BH:    bh,
		BW:    bw,
		H:     h,
		W:     w,
		Xb:    int(rand.Int63n(5)),
		Yb:    int(rand.Int63n(5)),
		Xo:    1,
		Yo:    -1,
		X:     x,
		Y:     y,
		color: color.RGBA{100, 100, 100, 255},
	}
}

func (this *Box) MoveOn(screen *ebiten.Image) {
	this.X += this.Xb * this.Xo
	this.Y += this.Yb * this.Yo
	this.Flip()
	this.DrawOn(screen)
}

func (this *Box) DrawOn(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(this.X), float64(this.Y), float64(this.W), float64(this.H), this.color)
}

func (this *Box) Flip() {
	//+1 to skip 0
	if this.X > this.BW {
		this.Xo = -1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	}

	if this.Y > this.BH {
		this.Yo = -1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	}

	if this.X < 0 {
		this.Xo = 1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	}

	if this.Y < 0 {
		this.Yo = 1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	}
}

func (this *Box) FlipOnCollision(target Box) {
	if target.X-this.X > 0 {
		this.Xo = -1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	} else {
		this.Xo = 1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	}
	if target.Y-this.Y > 0 {
		this.Yo = -1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	} else {
		this.Yo = 1
		this.Xb = int(rand.Int63n(5)) + 1
		this.Yb = int(rand.Int63n(5)) + 1
	}

	this.color.B += uint8(rand.Int63n(20))
	this.color.G += uint8(rand.Int63n(20))
	this.color.R += uint8(rand.Int63n(20))
	// this.color.B += int(rand.Int63n(5))
}

func (this *Box) RandColor() {
	this.color.B += uint8(rand.Int63n(200))
	this.color.G += uint8(rand.Int63n(200))
	this.color.R += uint8(rand.Int63n(200))
}

func (this *Box) IsCollided(target Box) bool {
	tx, ty := target.X+target.H/2, target.Y+target.W/2
	x, y := this.X+this.H/2, this.Y+this.W/2

	tr := target.W / 2
	if target.W > target.H {
		tr = target.H / 2
	}

	r := this.W / 2
	if this.W > this.H {
		r = this.H / 2
	}

	if (tx-x)*(tx-x)+(ty-y)*(ty-y) < (tr+r)*(tr+r) {
		this.FlipOnCollision(target)
		return true
	}
	return false
}

func (this *Box) IsExactCollided(target Box) bool {

	collisionX := this.X+this.W >= target.X &&
		target.X+target.W >= this.X
	// y轴方向碰撞？
	collisionY := this.Y+this.H >= target.Y &&
		target.Y+this.H >= this.Y
	// 只有两个轴向都有碰撞时才碰撞
	if collisionX && collisionY {
		this.FlipOnCollision(target)
		return true
	}
	return false
}

func (this *Box) DetectKeys() *Box {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		this.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		this.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		this.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		this.X -= 5
	}

	if this.X < 0 {
		this.X = 0
	}

	if this.Y < 0 {
		this.Y = 0
	}

	if this.X+this.W > this.BW {
		this.X = this.BW - this.W
	}

	if this.Y+this.H > this.BH {
		this.Y = this.BH - this.H
	}
	return this
}

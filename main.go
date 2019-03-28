package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/matthewgao/bubble/shape"
	"github.com/matthewgao/bubble/stage"
)

var p *shape.Box
var mp []*shape.Box
var bar *shape.Box

func update(screen *ebiten.Image) error {

	ebiten.SetCursorVisible(false)
	cx, cy := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d.%d %d.%d", p.X, p.Y, cx, cy))
	src, _ := ebiten.NewImage(20, 20, ebiten.FilterLinear)
	src.Fill(color.White)
	ebitenutil.DrawLine(screen, float64(cx), float64(cy), float64(cx)+0.1, float64(cy)+0.1, color.White)
	// ebitenutil.DrawRect(screen, float64(pp.X), float64(pp.Y), float64(pp.H), float64(p.W), color.White)
	p.DetectKeys()
	p.DrawOn(screen)
	for k, v := range mp {
		for ki, vi := range mp {
			if k == ki {
				continue
			}
			// v.IsCollided(*vi)
			v.IsExactCollided(*vi)
			// vi.IsCollided(v)
		}
	}

	for _, v := range mp {
		if v == p {
			continue
		}
		v.MoveOn(screen)
	}

	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Apply(float64(pp.X), float64(pp.Y))
	// r := image.Rect(-pp.X, -pp.Y, 10-pp.X, 10-pp.Y)
	// op.SourceRect = &r
	// screen.DrawImage(src, op)
	return nil
}

func updateGame(screen *ebiten.Image) error {
	ebiten.SetCursorVisible(false)
	cx, cy := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d.%d %d.%d", p.X, p.Y, cx, cy))
	src, _ := ebiten.NewImage(20, 20, ebiten.FilterLinear)
	src.Fill(color.White)
	ebitenutil.DrawLine(screen, float64(cx), float64(cy), float64(cx)+0.5, float64(cy)+0.5, color.White)

	bar.DetectKeys()
	bar.DrawOn(screen)
	p.MoveOn(screen)
	p.IsExactCollided(*bar)

	stage.FlushFrame(p, screen)
	return nil
}

func main() {
	p = shape.NewBox(240, 320, 10, 10)
	bar = shape.NewBox(240, 320, 10, 70)
	stage.StartStage()
	ebiten.SetWindowDecorated(true)
	ebiten.Run(updateGame, 320, 240, 2, "Hello world!")
}

// func main() {
// 	// p.X = 100
// 	// p.Y = 100
// 	mp = make([]*shape.Box, 0)
// 	p = shape.NewBox(240, 320, 10, 100)
// 	for i := 0; i < 80; i++ {
// 		x := shape.NewBox(240, 320, 10, 10)
// 		// x.Init()
// 		mp = append(mp, x)
// 	}
// 	mp = append(mp, p)
// 	ebiten.SetWindowDecorated(true)
// 	ebiten.Run(update, 320, 240, 2, "Hello world!")
// }

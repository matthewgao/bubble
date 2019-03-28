package stage

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/matthewgao/bubble/shape"
)

//FIXME: not render each block for each frame
var BlockGrp []*shape.Box

func StartStage() {
	for j := 0; j < 5; j++ {
		for i := 0; i < 32; i++ {
			x := shape.NewBoxFix(240, 320, 10, 10, i*10, j*10)
			x.RandColor()
			BlockGrp = append(BlockGrp, x)
		}
	}
}

func FlushFrame(p *shape.Box, screen *ebiten.Image) {
	tmp := make([]*shape.Box, 0)
	for _, v := range BlockGrp {
		if p.IsExactCollided(*v) {
			continue
		}
		tmp = append(tmp, v)
	}
	BlockGrp = tmp
	for _, v := range BlockGrp {
		v.DrawOn(screen)
	}
}

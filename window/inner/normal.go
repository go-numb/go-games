package inner

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/freetype/truetype"
	gaming "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

type Normal struct {
	count  int
	font   font.Face
	Fields map[int]*gaming.Image
}

func NewNormal() *Normal {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	fontFace := truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	// ディレクトリ内を取得
	dir, err := ioutil.ReadDir("./img")
	if err != nil {
		return nil
	}

	m := make(map[int]*gaming.Image)
	for i := range dir {
		if dir[i].IsDir() {
			continue
		}
		if !strings.HasSuffix(dir[i].Name(), "-field.jpeg") {
			continue
		}

		f, err := os.Open(filepath.Join("./img", dir[i].Name()))
		if err != nil {
			continue
		}

		img, err := jpeg.Decode(f)
		if err != nil {
			continue
		}
		f.Close()
		field, _ := gaming.NewImageFromImage(img, gaming.FilterDefault)
		m[i] = field
	}

	return &Normal{
		count:  0,
		font:   fontFace,
		Fields: m,
	}
}

func (p *Normal) Update(screen *gaming.Image) error {
	p.count++
	if gaming.IsDrawingSkipped() {
		return nil
	}

	text.Draw(
		screen,
		fmt.Sprintf("draw count: %d", p.count),
		p.font,
		40+p.count,
		40,
		color.White)

	op := &gaming.DrawImageOptions{}
	screen.DrawImage(p.Fields[0], op)

	return nil
}

func (p *Normal) Layout(w, h int) (sw, sh int) {

	return w, h
}

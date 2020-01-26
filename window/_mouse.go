package main

import (
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"

	gaming "github.com/hajimehoshi/ebiten"
)

const (
	WIDTH  = 1280
	HEIGTH = 720
)

var (
	count = 0
)

type Game struct {
	db     *leveldb.DB
	Images map[int]*gaming.Image
}

func NewGame() *Game {
	ldb, err := leveldb.OpenFile("./ldb", nil)
	if err != nil {
		return nil
	}

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
		if !strings.HasSuffix(dir[i].Name(), "-char.png") {
			continue
		}

		f, err := os.Open(filepath.Join("./img", dir[i].Name()))
		if err != nil {
			continue
		}

		img, _, err := image.Decode(f)
		if err != nil {
			continue
		}
		f.Close()
		charctor, _ := gaming.NewImageFromImage(img, gaming.FilterDefault)
		m[i] = charctor
	}

	return &Game{
		db:     ldb,
		Images: m,
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func (g *Game) Update(screen *gaming.Image) error {
	count++
	if gaming.IsDrawingSkipped() {
		return nil
	}
	// w, h := g.Images[0].Size()
	op := &gaming.DrawImageOptions{}

	// // Move the image's center to the screen's upper-left corner.
	// // This is a preparation for rotating. When geometry matrices are applied,
	// // the origin point is the upper-left corner.
	// op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// // Rotate the image. As a result, the anchor point of this rotate is
	// // the center of the image.
	// op.GeoM.Rotate(float64(count%360) * 2 * math.Pi / 360)

	// // Move the image to the screen's center.
	// sw, sh := screen.Size()
	// op.GeoM.Translate(float64(sw)/2, float64(sh)/2)

	screen.DrawImage(g.Images[0], op)

	x, y := gaming.CursorPosition()
	var add int
	if gaming.IsKeyPressed(gaming.KeyF) {
		add -= 400
	}
	op.GeoM.Translate(float64(x+add), float64(y))
	screen.DrawImage(g.Images[1], op)
	return nil
}

func main() {
	g := NewGame()

	gaming.SetWindowSize(WIDTH, HEIGTH)
	gaming.SetMaxTPS(60)
	gaming.SetWindowTitle("First game")
	gaming.SetWindowResizable(true)
	// gaming.IsKeyPressed()
	// gaming.SetFullscreen(true)
	if err := gaming.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

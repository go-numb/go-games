package main

import (
	"github.com/go-numb/go-games/window/inner"
	gaming "github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"
)

func init() {

}

const (
	W = 2560
	H = 720
)

func main() {
	g := inner.NewNormal()
	gaming.SetWindowSize(W, H)
	gaming.SetWindowTitle("Rotate (Resizable Window Demo)")
	gaming.SetWindowResizable(true)
	if err := gaming.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

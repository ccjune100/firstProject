package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	_ "image/png"
)

type firstGame struct {
	player *ebiten.Image
	xloc   int
	yloc   int
	score  int
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	playerPict, _, err := ebitenutil.NewImageFromFile("quaxly.png")
	if err != nil {
		fmt.Println("Error Loading Quaxly")
	}
	ourGame := firstGame{player: playerPict,
		yloc: 480,
		xloc: 500}
	err = ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}

func (game *firstGame) Update() error {
	game.xloc += 1
	if game.xloc > 1000 {
		game.xloc = 0
	}
	return nil
}

func (game *firstGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Blanchedalmond)
	drawOps := ebiten.DrawImageOptions{}
	for i := 0; i < 3; i++ {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Scale(1, 1)
		// for xloc if you do +i*250 there will be i amount of copies following the original png
		// and will reset to left side of screen once the last copy is off-screen but if you have
		// -i*250, the movement resets once the original png is off-screen.
		drawOps.GeoM.Translate(float64(game.xloc+i*250), float64(game.yloc))
		screen.DrawImage(game.player, &drawOps)
	}
}

func (game firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight //by default, just return the current dimensions
}

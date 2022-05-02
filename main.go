package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ramses2099/gamegolang/game"
)

var (
	shipImage *ebiten.Image
)

// init function
func init() {
	var err error
	shipImage, _, err = ebitenutil.NewImageFromFile("image/ship_0004.png")
	if err != nil {
		log.Fatal(err)
	}

}

// Game implements ebiten.Game interface.
type Game struct {
	Entity game.GameObject
	op     ebiten.DrawImageOptions
	inited bool
}

func (g *Game) Init() {
	defer func() {
		g.inited = true
	}()

	g.Entity = game.NewGameEntity(20, 23, 0, 0, 90, shipImage)

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	if !g.inited {
		g.Init()
	}
	// Write your game's logical update.
	err := g.Entity.Update()
	if err != nil {
		fmt.Printf("error update %s", err.Error())
		return nil
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	//draw image
	screen.DrawImage(shipImage, nil)
	//other item
	g.Entity.Draw(screen)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game Ebiten")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

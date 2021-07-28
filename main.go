package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const screenWidth int = 800
const screenHeight int = 600
const playerWidth int = 50
const playerHeight int = 80
const bossWidth int = 250
const bossHeight int = 200

type Game struct {
	boss *ebiten.Image
	bossPosX float64
	bossPosY float64
	player *ebiten.Image
	playPosX float64
	playerPosY float64
	playerBullet *ebiten.Image
	playerBulletPosX float64
	playerBulletPosY float64

	bossDx float64
}


func (g *Game) Update() error {


	for _, k := range inpututil.PressedKeys() {
		if k == ebiten.KeyRight {
			g.playPosX += 3
		} else if k == ebiten.KeyLeft {
			g.playPosX -= 3
		} else if k == ebiten.KeyUp {
			g.playerPosY -= 3
		} else if k == ebiten.KeyDown {
			g.playerPosY += 3
		}
	}

	g.playerBulletPosY -= 3


	if g.playPosX + float64(playerWidth) >= float64(screenWidth) {
		g.playPosX = float64(screenWidth) - float64(playerWidth)
	}

	if g.playPosX <= 0 {
		g.playPosX = 0
	}

	if g.playerPosY + float64(playerHeight) >= float64(screenHeight) {
		g.playerPosY = float64(screenHeight) - float64(playerHeight)
	}

	if g.playerPosY <= float64(screenHeight / 2) {
		g.playerPosY = float64(screenHeight / 2)
	}

	if g.bossPosX + float64(bossWidth) >= float64(screenWidth) {
		g.bossDx = -3
	}

	if g.bossPosX <= 0 {
		g.bossDx = 3
	}

	g.bossPosX += g.bossDx



	if g.playerBulletPosY <= g.bossPosY + float64(bossHeight) &&
		g.playerBulletPosX <= g.bossPosX + float64(bossWidth) &&
		g.playerBulletPosX >= g.bossPosX {
		g.playerBulletPosX = g.playPosX + float64(playerWidth / 2) - float64(5 / 2)
		g.playerBulletPosY = g.playerPosY - 20
	}

	if g.playerBulletPosY <= 0 {
		g.playerBulletPosX = g.playPosX + float64(playerWidth / 2) - float64(5 / 2)
		g.playerBulletPosY = g.playerPosY - 20
	}



	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playPosX, g.playerPosY)
	screen.DrawImage(g.player, op)

	opBullet := &ebiten.DrawImageOptions{}
	opBullet.GeoM.Translate(g.playerBulletPosX, g.playerBulletPosY)
	screen.DrawImage(g.playerBullet, opBullet)

	bossOp := &ebiten.DrawImageOptions{}
	bossOp.GeoM.Translate(g.bossPosX, g.bossPosY)
	screen.DrawImage(g.boss, bossOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Spaceship")

	playImg, _, err := ebitenutil.NewImageFromFile("ship.png")
	if err != nil {
		log.Fatalf("Load image error: %v", err)
	}

	g := &Game{}
	g.player = playImg
	g.playPosX = float64(screenWidth / 2) - float64(playerWidth / 2)
	g.playerPosY = float64(screenHeight /  2)

	bulletImg := ebiten.NewImage(5, 10)
	bulletImg.Fill(color.Black)
	g.playerBullet = bulletImg
	g.playerBulletPosX = g.playPosX + float64(playerWidth / 2) - float64(5 / 2)
	g.playerBulletPosY = g.playerPosY - 20


	bossImg, _, err := ebitenutil.NewImageFromFile("boss.png")
	if err != nil {
		log.Fatalf("Load image error: %v", err)
	}
	g.boss= bossImg
	g.bossPosX = float64(screenWidth / 2) - float64(bossWidth / 2)
	g.bossDx = 3


	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Run game error: %v" ,err)
	}
}
package main

import (
	"bytes"
	"image"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct{}

var menu = newMenu(FPos{1280 - 32, 0})

var money_multiplier_indecator, _, _ = ebitenutil.NewImageFromFile("./art/ui/X2.png")

var cursor_img *ebiten.Image

func (g *Game) Update() error {
	if state == 0 {
		updateMainMenu(&mainMenu)
	} else if state == 1 {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			main_clicked = false
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !main_clicked {
			mx, my := ebiten.CursorPosition()
			mx /= 10
			my /= 10
			if fcollide(FPos{2, 2}, FPos{30, 68}, FPos{float64(mx), float64(my)}, FPos{1, 1}) {
				state = 4
				rounds = round_1
				selected_round = 1
			} else if fcollide(FPos{35, 2}, FPos{30, 68}, FPos{float64(mx), float64(my)}, FPos{1, 1}) {
				state = 4
				rounds = round_2
				selected_round = 2
			} else if fcollide(FPos{66, 2}, FPos{30, 68}, FPos{float64(mx), float64(my)}, FPos{1, 1}) {
				state = 4
				rounds = round_3
				selected_round = 3
			} else if fcollide(FPos{96, 2}, FPos{30, 68}, FPos{float64(mx), float64(my)}, FPos{1, 1}) {
				state = 4
				rounds = round_4
				selected_round = 4
			} else if fcollide(FPos{126, 70}, FPos{2, 2}, FPos{float64(mx), float64(my)}, FPos{0.1, 0.1}) {
				state = 4
				rounds = round_5
				selected_round = 5
			}
			main_clicked = true
		}
	} else if state == 2 {
		mapLogic()
		updateMenu(&menu)
	} else if state == 3 {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			main_clicked = false
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !main_clicked {
			mx, my := ebiten.CursorPosition()
			mx /= 10
			my /= 10
			if fcollide(FPos{44, 53}, FPos{39, 14}, FPos{float64(mx), float64(my)}, FPos{1, 1}) {
				state = 0
				main_clicked = true
			}
		}
	} else if state == -1 {
		os.Exit(0)
	} else if state == 4 {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			main_clicked = false
		}
		modeMenu()
	}
	return nil
}

var font *text.GoTextFaceSource

func (g *Game) Draw(s *ebiten.Image) {
	if state == 0 {
		drawMaineMenu(s, &mainMenu)
	} else if state == 1 {
		img, _, err := ebitenutil.NewImageFromFile("./art/ui/difficulty_select.png")
		if err != nil {
			log.Fatal(err)
		}
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(10, 10)
		s.DrawImage(img, &op)
	} else if state == 2 {
		if health > 0 {
			ebitenutil.DebugPrint(s, "how you seeing this")
			mapDrawing(s)
			drawMenu(s, menu)
			drawOutline(s)
			top := text.DrawOptions{}
			top.GeoM.Translate(10, 10)
			text.Draw(s, "money: "+strconv.Itoa(money), &text.GoTextFace{
				Source: font,
				Size:   48,
			}, &top)
			top.GeoM.Reset()
			top.GeoM.Translate(10, 47)
			text.Draw(s, "health: "+strconv.Itoa(health), &text.GoTextFace{
				Source: font,
				Size:   48,
			}, &top)
			top.GeoM.Reset()
			top.GeoM.Translate(10, 85)
			text.Draw(s, "round: "+strconv.Itoa(current_round), &text.GoTextFace{
				Source: font,
				Size:   48,
			}, &top)

			if money_multiplier != 1 {
				op := ebiten.DrawImageOptions{}
				op.GeoM.Scale(2, 2)
				op.GeoM.Translate(240, 7)
				s.DrawImage(money_multiplier_indecator, &op)
			}
		} else {
			mapDrawing(s)
		}
	} else if state == 3 {
		img, _, err := ebitenutil.NewImageFromFile("./art/ui/win_screen.png")
		if err != nil {
			log.Fatal(err)
		}
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(10, 10)
		s.DrawImage(img, &op)
	} else if state == 4 {
		modeMenuDraw(s)
	}
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	mx, my := ebiten.CursorPosition()
	op.GeoM.Translate(float64(mx), float64(my))
	s.DrawImage(cursor_img, &op)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1280, 720
}

func main() {
	file, err := os.Open("./art/enemy/enemy1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	template, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	data, _ := os.ReadFile("./art/ui/font.ttf")
	font, _ = text.NewGoTextFaceSource(bytes.NewReader(data))

	new_cursor_img, _, err := ebitenutil.NewImageFromFile("./art/ui/cursor.png")
	if err != nil {
		log.Fatal(err)
	}
	cursor_img = new_cursor_img

	ebiten.SetWindowIcon([]image.Image{template})
	ebiten.SetWindowTitle("gober tower defence")
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	setUpMaps()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

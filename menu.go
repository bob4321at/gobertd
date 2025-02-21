package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Menu struct {
	pos       FPos
	img       *ebiten.Image
	in_or_out bool
}

type Button struct {
	pos  FPos
	size FPos
	id   int
}

type MainMenu struct {
	img     *ebiten.Image
	buttons []Button
}

var buttons = map[int]Button{
	0:  {FPos{1280 - 192, 0}, FPos{64, 64}, 1},
	1:  {FPos{1280 - 192, 64}, FPos{64, 64}, 2},
	2:  {FPos{1280 - 192, 128}, FPos{64, 64}, 3},
	3:  {FPos{1280 - 192, 192}, FPos{64, 64}, 4},
	4:  {FPos{1280 - 192, 256}, FPos{64, 64}, 5},
	5:  {FPos{1280 - 192, 320}, FPos{64, 64}, 6},
	6:  {FPos{1280 - 192, 384}, FPos{64, 64}, 7},
	7:  {FPos{1280 - 192, 448}, FPos{64, 64}, 8},
	8:  {FPos{1280 - 192, 512}, FPos{64, 64}, 9},
	9:  {FPos{1280 - 192, 576}, FPos{64, 64}, 10},
	10: {FPos{1280 - 192, 640}, FPos{64, 64}, 11},
	11: {FPos{1280 - 192, 704}, FPos{64, 64}, 14},
	12: {FPos{1280 - 192, 772}, FPos{64, 64}, 15},
}

var state = 0
var main_clicked = false

var mainMenu = MainMenu{
	ebiten.NewImage(128, 72),
	[]Button{
		{FPos{40, 21}, FPos{48, 24}, 0},
		{FPos{41, 49}, FPos{48, 24}, 0},
	},
}

func init() {
	new_mode_menu_img, _, err := ebitenutil.NewImageFromFile("./art/ui/allmodes.png")
	if err != nil {
		panic(err)
	}
	mode_menu_img = new_mode_menu_img
	mainMenu_img, _, err := ebitenutil.NewImageFromFile("./art/ui/gobertdstartmenu.png")
	if err != nil {
		log.Fatal(err)
	}

	mainMenu.img = ebiten.NewImageFromImage(mainMenu_img)
}

func updateButton(call func()) {
	call()
}

func newMenu(pos FPos) Menu {
	img, _, _ := ebitenutil.NewImageFromFile("./art/ui/tower_menu.png")
	return Menu{pos, img, false}
}

func drawMaineMenu(s *ebiten.Image, m *MainMenu) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	s.DrawImage(m.img, &op)
}

func updateMainMenu(m *MainMenu) {
	for bn := 0; bn < len(m.buttons); bn++ {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !main_clicked {
			rmx, rmy := ebiten.CursorPosition()
			mx, my := float64(rmx)/10, float64(rmy)/10
			if fcollide(m.buttons[bn].pos, m.buttons[bn].size, FPos{mx, my}, FPos{1, 1}) && !main_clicked {
				main_clicked = true
				if bn == 0 {
					state = 1
					placed = true
				} else if bn == 1 {
					state = -1
				}
			}
		}
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			main_clicked = false
		}
	}
}

func updateMenu(m *Menu) {
	rmx, rmy := ebiten.CursorPosition()
	mx, my := float64(rmx), float64(rmy)
	if mx < 1280 && mx > 0 {
		if my < 720 && my > 0 {
			if m.pos.x <= 1280-250 {
				_, w_y := ebiten.Wheel()
				if m.pos.y+w_y*40 > -378*2 && m.pos.y+w_y*40 < 10 {
					m.pos.y += w_y * 40
				}
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if m.pos.x <= 1280-250 {
					for b := 0; b < len(buttons); b++ {
						if fcollide(FPos{buttons[b].pos.x, buttons[b].pos.y + m.pos.y}, buttons[b].size, FPos{mx, my}, FPos{1, 1}) {
							updateButton(func() {
								selected_tower = buttons[b].id
							})
						}
					}
				}
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
				m.in_or_out = true
				m.pos.y = 0
				m.pos.x = 1280 - 64
				can_place = true
				placed = true
			}
			if fcollide(m.pos, FPos{128, 65}, FPos{mx, my}, FPos{1, 1}) || fcollide(FPos{m.pos.x + 64, m.pos.y}, FPos{64, 368 * 2}, FPos{mx, my}, FPos{1, 1}) {
				can_place = false
				if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
					m.pos.x = 1280 - 256
					m.in_or_out = true
				}
			}
		}
	}
}

func drawMenu(s *ebiten.Image, m Menu) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(m.pos.x, m.pos.y)
	s.DrawImage(m.img, &op)
}

func modeMenu() {
	mx, my := ebiten.CursorPosition()
	rmx, rmy := float64(mx)/10, float64(my)/10

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !main_clicked {
		if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{38, 0}, FPos{61, 20}) {
			current_mode = 0
			state = 2
		} else if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{2, 23}, FPos{41, 25}) {
			current_mode = 1
			state = 2
		} else if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{104, 1}, FPos{23, 23}) {
			current_mode = 2
			state = 2
		} else if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{92, 25}, FPos{35, 42}) {
			current_mode = 3
			state = 2
		} else if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{0, 50}, FPos{49, 23}) {
			current_mode = 4
			state = 2
		} else if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{1, 0}, FPos{34, 20}) {
			current_mode = 5
			state = 2
		} else if fcollide(FPos{rmx, rmy}, FPos{1, 1}, FPos{51, 22}, FPos{41, 39}) {
			current_mode = 6
			state = 2
		}
		main_clicked = true
	}
}

var mode_menu_img = ebiten.NewImage(16, 16)

func modeMenuDraw(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	s.DrawImage(mode_menu_img, &op)
}

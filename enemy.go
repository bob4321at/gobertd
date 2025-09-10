package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var health_bar_img, _, _ = ebitenutil.NewImageFromFile("./art/ui/health.png")

func init() {
	enemy_img1, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy1.png")
	enemy_images[1] = enemy_img1
	enemy_img2, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy2.png")
	enemy_images[2] = enemy_img2
	enemy_img3, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy3.png")
	enemy_images[3] = enemy_img3
	enemy_img4, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy4.png")
	enemy_images[4] = enemy_img4
	enemy_img5, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy5.png")
	enemy_images[5] = enemy_img5
	enemy_img6, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy6.png")
	enemy_images[6] = enemy_img6
	enemy_img7, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy7.png")
	enemy_images[7] = enemy_img7
	enemy_img8, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy8.png")
	enemy_images[8] = enemy_img8
	enemy_img9, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy9.png")
	enemy_images[9] = enemy_img9
	enemy_img10, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy10.png")
	enemy_images[10] = enemy_img10
	enemy_img11, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy11.png")
	enemy_images[11] = enemy_img11
	enemy_img12, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy12.png")
	enemy_images[12] = enemy_img12
	enemy_img13, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy13.png")
	enemy_images[13] = enemy_img13
	enemy_img14, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy14.png")
	enemy_images[14] = enemy_img14
	enemy_img15, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy15.png")
	enemy_images[15] = enemy_img15
	enemy_img16, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy16.png")
	enemy_images[16] = enemy_img16
	enemy_img17, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy17.png")
	enemy_images[17] = enemy_img17
	enemy_img18, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy18.png")
	enemy_images[18] = enemy_img18
	enemy_img19, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy19.png")
	enemy_images[19] = enemy_img19
	enemy_img20, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy20.png")
	enemy_images[20] = enemy_img20
	enemy_img21, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy21.png")
	enemy_images[21] = enemy_img21
	enemy_img22, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy22.png")
	enemy_images[22] = enemy_img22
	enemy_img23, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy23.png")
	enemy_images[23] = enemy_img23
	enemy_img24, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy24.png")
	enemy_images[24] = enemy_img24
	enemy_img25, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy25.png")
	enemy_images[25] = enemy_img25
	enemy_img26, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy26.png")
	enemy_images[26] = enemy_img26
	enemy_img27, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy27.png")
	enemy_images[27] = enemy_img27
	enemy_img28, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy28.png")
	enemy_images[28] = enemy_img28
	enemy_img29, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy29.png")
	enemy_images[29] = enemy_img29
	enemy_img30, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy30.png")
	enemy_images[30] = enemy_img30
	enemy_img31, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy31.png")
	enemy_images[31] = enemy_img31
	enemy_img32, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy32.png")
	enemy_images[32] = enemy_img32
	enemy_img33, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy33.png")
	enemy_images[33] = enemy_img33
	enemy_img34, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy34.png")
	enemy_images[34] = enemy_img34
	enemy_img35, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy35.png")
	enemy_images[35] = enemy_img35
	enemy_img36, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy36.png")
	enemy_images[36] = enemy_img36
	enemy_img37, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy37.png")
	enemy_images[37] = enemy_img37
	enemy_img38, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy38.png")
	enemy_images[38] = enemy_img38
	enemy_img39, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy39.png")
	enemy_images[39] = enemy_img39
	enemy_img40, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy40.png")
	enemy_images[40] = enemy_img40
	enemy_img41, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy41.png")
	enemy_images[41] = enemy_img41
	enemy_img42, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy42.png")
	enemy_images[42] = enemy_img42
	enemy_img43, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy43.png")
	enemy_images[43] = enemy_img43
	enemy_img44, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy44.png")
	enemy_images[44] = enemy_img44
	enemy_img45, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy45.png")
	enemy_images[45] = enemy_img45
	enemy_img46, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy46.png")
	enemy_images[46] = enemy_img46
	enemy_img47, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy47.png")
	enemy_images[47] = enemy_img47
	enemy_img48, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy48.png")
	enemy_images[48] = enemy_img48
	enemy_img49, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy49.png")
	enemy_images[49] = enemy_img49
	enemy_img50, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy50.png")
	enemy_images[50] = enemy_img50
	enemy_img51, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy51.png")
	enemy_images[51] = enemy_img51
	enemy_img52, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy52.png")
	enemy_images[52] = enemy_img52
	enemy_img53, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy53.png")
	enemy_images[53] = enemy_img53
	enemy_img54, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy54.png")
	enemy_images[54] = enemy_img54
	enemy_img55, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy55.png")
	enemy_images[55] = enemy_img55
	enemy_img56, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy56.png")
	enemy_images[56] = enemy_img56
	enemy_img57, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy57.png")
	enemy_images[57] = enemy_img57
	enemy_img58, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy58.png")
	enemy_images[58] = enemy_img58
	enemy_img59, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy59.png")
	enemy_images[59] = enemy_img59
	enemy_img60, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy60.png")
	enemy_images[60] = enemy_img60
	enemy_img61, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy61.png")
	enemy_images[61] = enemy_img61
	enemy_img62, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy62.png")
	enemy_images[62] = enemy_img62
	enemy_img101, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy101.png")
	enemy_images[101] = enemy_img101
	enemy_img102, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy102.png")
	enemy_images[102] = enemy_img102
	enemy_img103, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy103.png")
	enemy_images[103] = enemy_img103
	enemy_img104, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy104.png")
	enemy_images[104] = enemy_img104
	enemy_img105, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy105.png")
	enemy_images[105] = enemy_img105
	enemy_img106, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy106.png")
	enemy_images[106] = enemy_img106
	enemy_img107, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy107.png")
	enemy_images[107] = enemy_img107
	enemy_img108, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy108.png")
	enemy_images[108] = enemy_img108
	enemy_img109, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy109.png")
	enemy_images[109] = enemy_img109
	enemy_img110, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy110.png")
	enemy_images[110] = enemy_img110
	enemy_img111, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy111.png")
	enemy_images[111] = enemy_img111
	enemy_img112, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy112.png")
	enemy_images[112] = enemy_img112
	enemy_img113, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy113.png")
	enemy_images[113] = enemy_img113
	enemy_img63, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy63.png")
	enemy_images[63] = enemy_img63
	enemy_img64, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy64.png")
	enemy_images[64] = enemy_img64
	enemy_image70, _, _ := ebitenutil.NewImageFromFile("./art/enemy/retroenemy1.png")
	enemy_images[70] = enemy_image70
	enemy_image71, _, _ := ebitenutil.NewImageFromFile("./art/enemy/retroenemy2.png")
	enemy_images[71] = enemy_image71
	enemy_image72, _, _ := ebitenutil.NewImageFromFile("./art/enemy/retroenemy3.png")
	enemy_images[72] = enemy_image72
	enemy_image73, _, _ := ebitenutil.NewImageFromFile("./art/enemy/retroenemy4.png")
	enemy_images[73] = enemy_image73
	enemy_image74, _, _ := ebitenutil.NewImageFromFile("./art/enemy/retroenemy5.png")
	enemy_images[74] = enemy_image74
	enemy_img114, _, _ := ebitenutil.NewImageFromFile("./art/enemy/enemy114.png")
	enemy_images[114] = enemy_img114
}

func newEnemy(id int32, pos FPos) (e Enemy) {
	e_img := enemy_images[id]
	if id != 0 {
		e = Enemy{e_img, id, pos, 0, enemy_speed[id], int32(enemy_health[id]), int32(enemy_health[id]), false, 0}
	} else {
		e = Enemy{e_img, id, pos, 0, enemy_speed[id], int32(enemy_health[id]), int32(enemy_health[id]), false, -100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000}
	}
	return e
}

func updateEnemy(path []Pos, e *Enemy, h *int) {
	if e.img != enemy_images[0] && e.point != -1 {
		c_point := e.point

		dir := rad2deg(math.Atan2(e.pos.x-float64(path[c_point].x), e.pos.y-float64(path[c_point].y)))
		dir -= 90

		e.pos.x += math.Cos(deg2rad(dir)) * e.speed
		e.pos.y -= math.Sin(deg2rad(dir)) * e.speed

		if e.point >= 1 {
			e.travelled -= e.speed
		}

		if int32(e.pos.x) >= path[c_point].x-2 && int32(e.pos.x) <= path[c_point].x+2 {
			if int32(e.pos.y) >= path[c_point].y-2 && int32(e.pos.y) <= path[c_point].y+2 {
				if e.point+1 < int32(len(path)) {
					e.point += 1
				}
			}
		}
		if e.point >= int32(len(path)-1) && !e.money_given {
			*h -= int(e.health)
			e.img = enemy_images[0]
			e.money_given = true
			e.r_health = -1
			e.health = -1
		}

		if e.r_health < 1 {
			e.img = ebiten.NewImage(1, 1)
		}
	}
}

func drawEnemy(s *ebiten.Image, e Enemy) {
	if e.img != enemy_images[0] && !e.money_given {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(2, 2)
		op.GeoM.Translate(float64(e.pos.x)*10-32, float64(e.pos.y)*10-32)
		s.DrawImage(e.img, &op)
		op.GeoM.Reset()
		op.GeoM.Scale(float64(e.r_health)*10, 5)
		op.GeoM.Translate(float64(e.pos.x)*10-16, float64(e.pos.y)*10-32)
		s.DrawImage(health_bar_img, &op)
	}
}

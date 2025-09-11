package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type level struct {
	img     *ebiten.Image
	id      int32
	path    []Pos
	enemies []Enemy
	tower   []TowerFamily
}

var gober_secret_keys = map[int]bool{
	0: false,
	1: false,
	2: false,
	3: false,
}
var r_gober_secret_keys = map[int]ebiten.Key{
	0: ebiten.KeyG,
	1: ebiten.KeyO,
	2: ebiten.KeyB,
	3: ebiten.KeyE,
}

var current_round = 0

var current_mode = 0

var damage_divider = 1

var can_upgrade = true

var thanos_timer = -1

var selected_round = -1

var levels = map[int32]level{
	1: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 62}, {0, 62}, {36, 59}, {50, 54}, {49, 42}, {19, 25}, {26, 18}, {61, 18}, {77, 33}, {88, 54}, {127, 57}, {127, 57}}, []Enemy{}, []TowerFamily{}},
	2: {ebiten.NewImage(128, 72), 0, []Pos{{12, -10}, {12, 10}, {13, 33}, {31, 57}, {62, 17}, {88, 58}, {114, 40}, {127, 43}, {127, 43}}, []Enemy{}, []TowerFamily{}},
	3: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 57}, {0, 57}, {32, 56}, {25, 34}, {30, 13}, {64, 6}, {92, 13}, {97, 31}, {85, 55}, {81, 71}, {81, 71}}, []Enemy{}, []TowerFamily{}},
	4: {ebiten.NewImage(128, 72), 0, []Pos{{-7, 31}, {0, 31}, {40, 32}, {83, 31}, {99, 24}, {116, 35}, {118, 59}, {84, 57}, {78, 49}, {40, 47}, {0, 47}, {0, 47}}, []Enemy{}, []TowerFamily{}},
	5: {ebiten.NewImage(128, 72), 0, []Pos{{32, -5}, {32, 0}, {33, 13}, {51, 14}, {84, 18}, {84, 54}, {58, 54}, {57, 36}, {29, 34}, {32, 59}, {0, 55}, {0, 55}}, []Enemy{}, []TowerFamily{}},
	6: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 23}, {0, 23}, {32, 24}, {33, 0}, {33, -10}}, []Enemy{}, []TowerFamily{}}, 7: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 29}, {0, 29}, {30, 36}, {39, 55}, {60, 55}, {69, 46}, {67, 36}, {87, 37}, {91, 22}, {102, 20}, {112, 38}, {127, 37}, {127, 37}}, []Enemy{}, []TowerFamily{}},
	8:  {ebiten.NewImage(128, 72), 0, []Pos{{-10, 68}, {0, 68}, {26, 53}, {54, 42}, {85, 27}, {111, 17}, {120, 37}, {113, 56}, {79, 48}, {57, 36}, {37, 24}, {17, 10}, {0, 0}, {0, 0}}, []Enemy{}, []TowerFamily{}},
	9:  {ebiten.NewImage(128, 72), 0, []Pos{{31, -10}, {31, 0}, {48, 27}, {65, 47}, {90, 65}, {116, 62}, {125, 46}, {115, 21}, {62, 22}, {27, 14}, {13, 31}, {30, 68}, {63, 68}, {103, 60}, {110, 40}, {76, 37}, {36, 40}, {0, 45}, {0, 45}}, []Enemy{}, []TowerFamily{}},
	10: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 15}, {0, 15}, {39, 23}, {39, 23}, {32, 41}, {17, 33}, {30, 13}, {82, 23}, {91, 28}, {81, 37}, {67, 30}, {82, 20}, {106, 24}, {106, 56}, {86, 60}, {73, 56}, {84, 46}, {93, 56}, {80, 55}, {22, 55}, {23, 52}, {53, 52}, {53, 58}, {0, 57}, {0, 57}}, []Enemy{}, []TowerFamily{}},
	11: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 32}, {0, 32}, {36, 34}, {37, 65}, {79, 44}, {114, 15}, {63, 15}, {59, 0}, {59, 0}}, []Enemy{}, []TowerFamily{}},
	12: {ebiten.NewImage(128, 72), 0, []Pos{{60, -10}, {50, 0}, {27, 11}, {15, 34}, {25, 60}, {67, 65}, {97, 55}, {100, 42}, {100, 26}, {70, 16}, {45, 21}, {47, 43}, {66, 47}, {77, 41}, {63, 35}, {63, 35}}, []Enemy{}, []TowerFamily{}},
	13: {ebiten.NewImage(128, 72), 0, []Pos{{-10, 64}, {0, 64}, {33, 37}, {62, 11}, {85, 29}, {114, 50}, {73, 50}, {30, 53}, {32, 72}, {32, 72}}, []Enemy{}, []TowerFamily{}},
}
var current_level = 1
var changed_level = false
var selected_map = 0

var map_debug = false
var map_debug_clicked = false

var started = false

var health = 100
var money = 500
var money_multiplier = 1
var placed = false
var can_place = true

var secret_clicked = false

var map_12_rock *ebiten.Image

var deflation_start = false

func setUpMaps() {
	map_img1, _, _ := ebitenutil.NewImageFromFile("./art/maps/map1.png")
	levels[1] = level{map_img1, levels[1].id, levels[1].path, levels[1].enemies, levels[1].tower}
	map_img2, _, _ := ebitenutil.NewImageFromFile("./art/maps/map2.png")
	levels[2] = level{map_img2, levels[2].id, levels[2].path, levels[2].enemies, levels[2].tower}
	map_img3, _, _ := ebitenutil.NewImageFromFile("./art/maps/map3.png")
	levels[3] = level{map_img3, levels[3].id, levels[3].path, levels[3].enemies, levels[3].tower}
	map_img4, _, _ := ebitenutil.NewImageFromFile("./art/maps/map4.png")
	levels[4] = level{map_img4, levels[4].id, levels[4].path, levels[4].enemies, levels[4].tower}
	map_img5, _, _ := ebitenutil.NewImageFromFile("./art/maps/map5.png")
	levels[5] = level{map_img5, levels[5].id, levels[5].path, levels[5].enemies, levels[5].tower}
	map_img6, _, _ := ebitenutil.NewImageFromFile("./art/maps/map6.png")
	levels[6] = level{map_img6, levels[6].id, levels[6].path, levels[6].enemies, levels[6].tower}
	map_img7, _, _ := ebitenutil.NewImageFromFile("./art/maps/map7.png")
	levels[7] = level{map_img7, levels[7].id, levels[7].path, levels[7].enemies, levels[7].tower}
	map_img8, _, _ := ebitenutil.NewImageFromFile("./art/maps/map8.png")
	levels[8] = level{map_img8, levels[8].id, levels[8].path, levels[8].enemies, levels[8].tower}
	map_img9, _, _ := ebitenutil.NewImageFromFile("./art/maps/map9.png")
	levels[9] = level{map_img9, levels[9].id, levels[9].path, levels[9].enemies, levels[9].tower}
	map_img10, _, _ := ebitenutil.NewImageFromFile("./art/maps/map10.png")
	levels[10] = level{map_img10, levels[10].id, levels[10].path, levels[10].enemies, levels[10].tower}
	map_img11, _, _ := ebitenutil.NewImageFromFile("./art/maps/map11.png")
	levels[11] = level{map_img11, levels[11].id, levels[11].path, levels[11].enemies, levels[11].tower}
	map_img12, _, _ := ebitenutil.NewImageFromFile("./art/maps/map12.png")
	levels[12] = level{map_img12, levels[12].id, levels[12].path, levels[12].enemies, levels[12].tower}
	map_img13, _, _ := ebitenutil.NewImageFromFile("./art/maps/map13.png")
	levels[13] = level{map_img13, levels[13].id, levels[13].path, levels[13].enemies, levels[13].tower}

	new_map_12_rock, _, err := ebitenutil.NewImageFromFile("./art/maps/props/rock.png")
	if err != nil {
		log.Fatal(err)
	}
	map_12_rock = ebiten.NewImageFromImage(new_map_12_rock)
}

func mapLogic() {
	if health > 0 {
		ordered_enemies := BubbleSort(levels[int32(current_level)].enemies)
		for e := 0; e < len(levels[int32(current_level)].enemies); e++ {
			levels[int32(current_level)].enemies[e] = ordered_enemies[e]
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) && current_level+1 < len(levels)+1 && !changed_level && current_round == 0 && selected_map == 0 {
			current_level += 1
			changed_level = true
		} else if ebiten.IsKeyPressed(ebiten.KeyA) && current_level-1 > 0 && !changed_level && selected_map == 0 {
			current_level -= 1
			changed_level = true
		}
		if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
			changed_level = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyTab) && !map_debug_clicked {
			map_debug = !map_debug
			map_debug_clicked = true
		}
		if !ebiten.IsKeyPressed(ebiten.KeyTab) {
			map_debug_clicked = false
		}

		if ebiten.IsKeyPressed(ebiten.KeyL) && !changed_level && current_round == 0 {
			money_multiplier = 2
		}

		if !changed_level && !deflation_start && current_round == 0 && current_mode == 1 {
			money = 15000
			deflation_start = true
		} else if !changed_level && current_round == 0 && current_mode == 2 {
			health = 10
			deflation_start = true
		} else if !changed_level && current_round == 0 && current_mode == 3 {
			damage_divider = 2
			deflation_start = true
		} else if !changed_level && current_round == 0 && current_mode == 4 {
			can_upgrade = false
			deflation_start = true
		} else if !changed_level && current_round == 0 && current_mode == 5 {
			thanos_timer = 0
			deflation_start = true
		} else if !changed_level && current_round == 0 && current_mode == 6 {
			if selected_round == 1 {
				rounds = round_1_rainbow
			}
			if selected_round == 2 {
				rounds = round_2_rainbow
			}
			if selected_round == 3 {
				rounds = round_3_rainbow
			}
			if selected_round == 4 {
				rounds = round_4_rainbow
			}
			if selected_round == 5 {
				rounds = round_5_rainbow
			}
		}
		if !changed_level && current_round == 0 && current_mode == 7 && !deflation_start {
			money = 1000
			deflation_start = true
			if selected_round == 1 {
				rounds = round_1_boss
			}
			if selected_round == 2 {
				rounds = round_2_boss
			}
			if selected_round == 3 {
				rounds = round_3_boss
			}
			if selected_round == 4 {
				rounds = round_4_boss
			}
			if selected_round == 5 {
				rounds = round_5_boss
			}
		}

		if ebiten.IsKeyPressed(ebiten.Key1) {
			selected_tower = 1
		} else if ebiten.IsKeyPressed(ebiten.Key2) {
			selected_tower = 2
		} else if ebiten.IsKeyPressed(ebiten.Key3) {
			selected_tower = 3
		} else if ebiten.IsKeyPressed(ebiten.Key4) {
			selected_tower = 4
		} else if ebiten.IsKeyPressed(ebiten.Key5) {
			selected_tower = 5
		} else if ebiten.IsKeyPressed(ebiten.Key6) {
			selected_tower = 6
		} else if ebiten.IsKeyPressed(ebiten.Key7) {
			selected_tower = 7
		} else if ebiten.IsKeyPressed(ebiten.Key8) {
			selected_tower = 8
		} else if ebiten.IsKeyPressed(ebiten.Key9) {
			selected_tower = 9
		} else if ebiten.IsKeyPressed(ebiten.Key0) {
			selected_tower = 10
		} else if ebiten.IsKeyPressed(ebiten.KeyQ) {
			selected_tower = 11
		} else if ebiten.IsKeyPressed(ebiten.KeyP) {
			selected_tower = 12
		} else if ebiten.IsKeyPressed(ebiten.KeyB) {
			selected_tower = 13
		} else if ebiten.IsKeyPressed(ebiten.KeyW) {
			selected_tower = 14
		} else if ebiten.IsKeyPressed(ebiten.KeyE) {
			selected_tower = 15
		} else if ebiten.IsKeyPressed(ebiten.KeyG) {
			selected_tower = 16
		} else if ebiten.IsKeyPressed(ebiten.KeyH) {
			selected_tower = 17
		}

		if !secret_clicked {
			for k := 0; k < len(gober_secret_keys); k++ {
				if !ebiten.IsKeyPressed(r_gober_secret_keys[k]) {
					gober_secret_keys[k] = false
				}
				if ebiten.IsKeyPressed(r_gober_secret_keys[k]) {
					gober_secret_keys[k] = true
				}
			}

			working := true

			for k := 0; k < len(gober_secret_keys); k++ {
				if !gober_secret_keys[k] {
					working = false
				}
			}

			if working {
				money += 500
				secret_clicked = true
			}
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && money >= tower_cost[selected_tower] && !placed && can_place {
			new_towers := levels[int32(current_level)].tower
			mx, my := ebiten.CursorPosition()
			new_towers = append(new_towers, newTower(int32(selected_tower), FPos{float64(mx) - 32, float64(my) - 32}))
			levels[int32(current_level)] = level{levels[int32(current_level)].img, levels[int32(current_level)].id, levels[int32(current_level)].path, levels[int32(current_level)].enemies, new_towers}
			money -= tower_cost[selected_tower]
			placed = true
		}
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			placed = false
		}

		if !started && ebiten.IsKeyPressed(ebiten.KeySpace) {
			selected_map = current_level
			if thanos_timer != -1 && thanos_timer < 5 {
				thanos_timer += 1
				if thanos_timer >= 5 {
					num := rand.Float64() * float64(len(levels[int32(current_level)].tower))
					levels[int32(current_level)].tower[int(num)] = newTower(1, FPos{-1000, 1000})
					thanos_timer = 0
				}
			}
			if current_round < 1 {
				started = true
				current_round += 1
			} else {
				started = true
			}
			new_enemies := []Enemy{newEnemy(0, FPos{1000000, 100000})}
			for i := 0; i < len(levels[int32(current_level)].enemies); i++ {
				for j := 0; j < len(levels[int32(current_level)].tower[i].stages); j++ {
					levels[int32(current_level)].tower[i].stages[j].attacks = []Attack{}
				}
				new_enemies = append(new_enemies, levels[int32(current_level)].enemies[i])
			}
			for i := 0; i < len(rounds[int32(current_round)]); i++ {
				fpos := FPos{float64(levels[int32(current_level)].path[0].x) - float64(i*15), float64(levels[int32(current_level)].path[0].y)}
				new_enemies = append(new_enemies, newEnemy(rounds[int32(current_round)][i], fpos))
			}
			levels[int32(current_level)] = level{levels[int32(current_level)].img, levels[int32(current_level)].id, levels[int32(current_level)].path, new_enemies, levels[int32(current_level)].tower}
		}

		bad := 0

		for i := 0; i < len(levels[int32(current_level)].enemies); i++ {
			en := levels[int32(current_level)].enemies[i]
			if en.money_given {
				bad += 1
			}
			if bad == len(levels[int32(current_level)].enemies) && current_round < len(rounds) {
				started = false
				levels[int32(current_level)] = level{levels[int32(current_level)].img, levels[int32(current_level)].id, levels[int32(current_level)].path, []Enemy{}, levels[int32(current_level)].tower}
				current_round += 1
			}
		}

		if len(levels[int32(current_level)].enemies) > 0 {
			for i := 0; i < len(levels[int32(current_level)].enemies); i++ {
				updateEnemy(levels[int32(current_level)].path, &levels[int32(current_level)].enemies[i], &health)
			}
		}

		if len(levels[int32(current_level)].enemies) >= 0 {
			for i := 0; i < len(levels[int32(current_level)].enemies); i++ {
				en := &levels[int32(current_level)].enemies[i]
				if en.r_health <= 0 && !en.money_given && current_round > 0 {
					if current_mode != 1 {
						money += 50 * money_multiplier
					}

					en.money_given = true
				}
			}
		}

		for i := 0; i < len(levels[int32(current_level)].tower); i++ {
			updateTower(&levels[int32(current_level)].tower[i], levels[int32(current_level)].enemies)
		}

		if current_round >= len(rounds) {
			levels[int32(current_level)] = level{levels[int32(current_level)].img, levels[int32(current_level)].id, levels[int32(current_level)].path, []Enemy{}, []TowerFamily{}}
			current_round = 0
			state = 0
			current_level = 1
			changed_level = false
			selected_map = 0
			map_debug = false
			map_debug_clicked = false
			started = false
			health = 100
			money = 500
			money_multiplier = 1
			placed = false
			can_place = true
			state = 3
			thanos_timer = -1
		}
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			rmx, rmy := ebiten.CursorPosition()
			mx, my := float64(rmx), float64(rmy)
			if fcollide(FPos{44 * 10, 57 * 10}, FPos{40 * 10, 13 * 10}, FPos{mx, my}, FPos{1, 1}) {
				levels[int32(current_level)] = level{levels[int32(current_level)].img, levels[int32(current_level)].id, levels[int32(current_level)].path, []Enemy{}, []TowerFamily{}}
				current_round = 0
				state = 0
				current_level = 1
				changed_level = false
				selected_map = 0
				map_debug = false
				map_debug_clicked = false
				started = false
				health = 100
				money = 500
				money_multiplier = 1
				placed = false
				can_place = true

				deflation_start = false
				can_upgrade = true
				damage_divider = 1
				current_mode = 0
				thanos_timer = -1
			}
		}
	}
}

func mapDrawing(s *ebiten.Image) {
	if health > 0 {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(10, 10)
		s.DrawImage(levels[int32(current_level)].img, &op)
		for i := 0; i < len(levels[int32(current_level)].enemies); i++ {
			info := levels[int32(current_level)].enemies[i]
			drawEnemy(s, info)
		}
		if current_level == 12 {
			op.GeoM.Reset()
			op.GeoM.Translate(57, 27)
			op.GeoM.Scale(10, 10)
			s.DrawImage(map_12_rock, &op)
		}
		for i := 0; i < len(levels[int32(current_level)].tower); i++ {
			info := levels[int32(current_level)].tower[i]
			drawTower(s, info)
		}
		if map_debug {
			for i := 0; i < len(levels[int32(current_level)].path); i++ {
				info := levels[int32(current_level)].path[i]
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(info.x), float64(info.y))
				op.GeoM.Scale(10, 10)

				debug_img, _, _ := ebitenutil.NewImageFromFile("./art/ui/point.png")

				s.DrawImage(debug_img, &op)
			}
		}
		if current_round == 0 {
			top := text.DrawOptions{}
			top.GeoM.Translate(500, 10)
			text.Draw(s, "hit a or d to cycle throuh the maps", &text.GoTextFace{
				Source: font,
				Size:   48,
			}, &top)
		}
	} else {
		img, _, err := ebitenutil.NewImageFromFile("./art/ui/lose_screen.png")
		if err != nil {
			panic(err)
		}
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(10, 10)
		s.DrawImage(img, &op)
	}
}

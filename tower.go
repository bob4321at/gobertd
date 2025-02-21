package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var upgraded = false

var selected_tower = 1

func newTower(id int32, pos FPos) TowerFamily {
	t := TowerFamily{[]Tower{{tower_images[id], id, pos, tower_attack[id], []Attack{}}, {tower_upgrade_images[id], id, pos, tower_upgrade_attack[id], []Attack{}}, {tower_upgrade_upgrade_images[id], id, pos, tower_upgrade_upgrade_attack[id], []Attack{}}, {tower_upgrade_upgrade_upgrade_images[id], id, pos, tower_upgrade_upgrade_upgrade_attack[id], []Attack{}}}, 0, false}
	return t
}

func updateTower(tf *TowerFamily, es []Enemy) {
	t := &tf.stages[tf.current_stage]
	if !tf.sold {

		if t.id == 0 {
			return
		}

		t.attack.timer -= 0.1

		// target
		if t.attack.timer < 0 {
			t.attack.timer = t.attack.o_timer
			t_atck := tower_attack[t.id]
			t_atck.pos = FPos{t.pos.x, t.pos.y}
			if len(es) != 0 {
				first_index := 0
				for e := 0; e < len(es); e++ {
					en := &es[e]
					if t.pos.x > (en.pos.x*10)-300 && t.pos.x <= (en.pos.x*10)+300 {
						if t.pos.y > (en.pos.y*10)-300 && t.pos.y <= (en.pos.y*10)+300 {
							if es[e].img != enemy_images[0] && !es[e].money_given {
								if t.pos.x > (en.pos.x*10)-300 && t.pos.x <= (en.pos.x*10)+300 {
									if t.pos.y > (en.pos.y*10)-300 && t.pos.y <= (en.pos.y*10)+300 {
										first_index = e
										e = len(es) + 1
									}
								}
							}
						}
					}
				}
				if first_index != -1 && es[first_index].id != 0 {
					dir := math.Atan2((es[first_index].pos.x*10-32)-t.pos.x, (es[first_index].pos.y*10-32)-t.pos.y)
					t_atck.vel = FPos{math.Sin(dir), math.Cos(dir)}
					t.attacks = append(t.attacks, t_atck)
					if t.id == 12 {
						money += 10 * int(t.attack.damage)
					}
				}
			}
		}

		// attack
		for i := 0; i < len(t.attacks); i++ {
			if !t.attacks[i].hit {
				t.attacks[i].pos.x += t.attacks[i].vel.x * t.attacks[i].projectile_speed
				t.attacks[i].pos.y += t.attacks[i].vel.y * t.attacks[i].projectile_speed
				if t.attacks[i].pos.x < 0 && t.attacks[i].pos.x > 1280 {
					if t.attacks[i].pos.y < 0 && t.attacks[i].pos.y > 720 {
						t.attacks[i].pos = FPos{1000, 1000}
						t.attacks[i].img = ebiten.NewImage(8, 8)
					}
				}
				for j := 0; j < len(es); j++ {
					if es[j].img != ebiten.NewImage(8, 8) {
						e_pos := es[j].pos
						e_pos.x *= 10
						e_pos.y *= 10
						if es[j].r_health >= 0 {
							if fcollide(t.attacks[i].pos, FPos{32, 32}, FPos{e_pos.x - 32, e_pos.y - 32}, FPos{32, 32}) {
								es[j].r_health -= t.attack.damage / int32(damage_divider)
								t.attacks[i].hit = true
								t.attacks[i].pos = FPos{1000, 1000}
								t.attacks[i].img = ebiten.NewImage(8, 8)
								if t.id == 13 {
									t.id = 0
								}
							}
						}
					}
				}
			}
		}
	}

	mx, my := ebiten.CursorPosition()
	rmx, rmy := float64(mx), float64(my)

	if can_upgrade {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !tf.sold && !upgraded && tower_upgrade_cost[int(t.id)] != -1 {
			if fcollide(t.pos, FPos{64, 64}, FPos{rmx, rmy}, FPos{1, 1}) {
				if tf.current_stage < len(tf.stages)-1 {
					if tf.current_stage == 0 {
						if money >= tower_upgrade_cost[int(t.id)] {
							tf.current_stage += 1
							money -= tower_upgrade_cost[int(t.id)]
							upgraded = true
						}
					} else if tf.current_stage == 1 {
						if money >= tower_upgrade_upgrade_cost[int(t.id)] {
							tf.current_stage += 1
							money -= tower_upgrade_upgrade_cost[int(t.id)]
							upgraded = true
						}
					} else if tf.current_stage == 2 {
						if money >= tower_upgrade_upgrade_upgrade_cost[int(t.id)] {
							tf.current_stage += 1
							money -= tower_upgrade_upgrade_upgrade_cost[int(t.id)]
							upgraded = true
						}
					}
				}
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && !tf.sold {
		if fcollide(t.pos, FPos{64, 64}, FPos{rmx, rmy}, FPos{1, 1}) {
			tf.sold = true
			if tf.current_stage == 0 {
				money += tower_cost[int(t.id)] / 2
			} else if tf.current_stage == 1 {
				money += (tower_cost[int(t.id)] / 2) + (tower_upgrade_cost[int(t.id)] / 2)
			} else if tf.current_stage == 2 {
				money += (tower_cost[int(t.id)] / 2) + (tower_upgrade_cost[int(t.id)] / 2) + (tower_upgrade_upgrade_cost[int(t.id)] / 2)
			}
		}
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		upgraded = false
	}
}

func drawOutline(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	mx, my := ebiten.CursorPosition()
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(float64(mx)-32, float64(my)-32)
	s.DrawImage(tower_images_outlines[int32(selected_tower)], &op)
}

func drawTower(s *ebiten.Image, tf TowerFamily) {
	t := tf.stages[tf.current_stage]
	if t.id != 0 && !tf.sold {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(5, 5)
		op.GeoM.Translate(t.pos.x, t.pos.y)
		s.DrawImage(t.img, &op)
		for i := 0; i < len(t.attacks); i++ {
			op.GeoM.Reset()
			op.GeoM.Scale(4, 4)
			if t.id == 13 {
				op.GeoM.Scale(5, 5)
			}
			op.GeoM.Translate(t.attacks[i].pos.x+16, t.attacks[i].pos.y+16)
			s.DrawImage(t.attack.img, &op)
		}
	}
}

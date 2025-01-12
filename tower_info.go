package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tower struct {
	img     *ebiten.Image
	id      int32
	pos     FPos
	attack  Attack
	attacks []Attack
}

type TowerFamily struct {
	stages        []Tower
	current_stage int
	sold          bool
}

type Attack struct {
	img              *ebiten.Image
	id               int32
	pos              FPos
	vel              FPos
	timer            float64
	o_timer          float64
	projectile_speed float64
	damage           int32
	hit              bool
}

var tower_images = map[int32]*ebiten.Image{}
var tower_upgrade_images = map[int32]*ebiten.Image{}
var tower_upgrade_upgrade_images = map[int32]*ebiten.Image{}
var tower_images_outlines = map[int32]*ebiten.Image{}
var tower_cost = map[int]int{
	1:  300,
	2:  300,
	3:  1250,
	4:  600,
	5:  1500,
	6:  700,
	7:  500,
	8:  750,
	9:  500,
	10: 400,
	11: 1000,
	12: 600,
	13: 750,
	14: 800,
	15: 750,
}
var tower_upgrade_cost = map[int]int{
	1:  200,
	2:  300,
	3:  900,
	4:  500,
	5:  2250,
	6:  400,
	7:  250,
	8:  1000,
	9:  300,
	10: 500,
	11: 1250,
	12: 800,
	13: -1,
	14: 1200,
	15: 1000,
}
var tower_upgrade_upgrade_cost = map[int]int{
	1:  250,
	2:  300,
	3:  1000,
	4:  500,
	5:  3000,
	6:  600,
	7:  550,
	8:  1250,
	9:  500,
	10: 650,
	11: 1250,
	12: 1250,
	14: 1500,
	15: 1250,
}
var tower_attack = map[int32]Attack{
	1: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		3,
		3,
		8,
		1,
		false,
	},
	2: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		6,
		6,
		4,
		3,
		false,
	},
	3: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		0.6,
		0.6,
		40,
		2,
		false,
	},
	4: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		2,
		2,
		7,
		2,
		false,
	},
	5: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		0.2,
		0.2,
		9,
		1,
		false,
	},
	6: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		2,
		2,
		4,
		3,
		false,
	},
	7: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		2.5,
		2.5,
		8,
		2,
		false,
	},
	8: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		6,
		6,
		12,
		9,
		false,
	},
	9: {
		ebiten.NewImage(8, 8),
		4,
		FPos{0, 0},
		FPos{0, 0},
		12, //cooldown
		12,
		20,    //projectile speed
		4,     //damgage
		false, //idk
	},
	10: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		5.5,
		5.5,
		10,
		2,
		false,
	},
	11: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		2.5,
		2.5,
		8,
		3,
		false,
	},
	12: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		5,
		5,
		8,
		1,
		false,
	},
	13: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		10,
		10,
		8,
		150,
		false,
	},
	14: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		7,
		7,
		3,
		12,
		false,
	},
	15: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		3,
		3,
		10,
		2,
		false,
	},
}
var tower_upgrade_attack = map[int32]Attack{
	1: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		3,
		3,
		8,
		2,
		false,
	},
	2: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		5.5,
		5.5,
		4,
		5,
		false,
	},
	3: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		0.5,
		0.5,
		40,
		3,
		false,
	},
	4: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		1.5,
		1.5,
		7,
		3,
		false,
	},
	5: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		0.1,
		0.1,
		9,
		2,
		false,
	},
	6: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		2,
		2,
		4,
		4,
		false,
	},
	7: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		1.5,
		1.5,
		8,
		2,
		false,
	},
	8: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		5.5,
		5.5,
		12,
		15,
		false,
	},
	9: {
		ebiten.NewImage(8, 8),
		4,
		FPos{0, 0},
		FPos{0, 0},
		12, //cooldown
		12,
		20,    //projectile speed
		6,     //damgage
		false, //idk
	},
	10: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		5.5,
		5.5,
		10,
		5,
		false,
	},
	11: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		1.25,
		1.25,
		8,
		4,
		false,
	},
	12: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		5,
		5,
		8,
		3,
		false,
	},
	14: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		8,
		8,
		2,
		24,
		false,
	},
	15: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		2.5,
		2.5,
		15,
		4,
		false,
	},
}
var tower_upgrade_upgrade_attack = map[int32]Attack{
	1: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		2.5,
		2.5,
		8,
		3,
		false,
	},
	2: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		5,
		5,
		4,
		7,
		false,
	},
	3: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		0.4,
		0.4,
		60,
		4,
		false,
	},
	4: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		1,
		1,
		7,
		3,
		false,
	},
	5: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		0.05,
		0.05,
		9,
		3,
		false,
	},
	6: {
		ebiten.NewImage(8, 8),
		3,
		FPos{0, 0},
		FPos{0, 0},
		1.5,
		1.5,
		4,
		5,
		false,
	},
	7: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		1.2,
		1.2,
		8,
		3,
		false,
	},
	8: {
		ebiten.NewImage(8, 8),
		1,
		FPos{0, 0},
		FPos{0, 0},
		4.5,
		4.5,
		12,
		21,
		false,
	},
	9: {
		ebiten.NewImage(8, 8),
		4,
		FPos{0, 0},
		FPos{0, 0},
		8, //cooldown
		8,
		20,    //projectile speed
		8,     //damgage
		false, //idk
	},
	10: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		3.5,
		3.5,
		10,
		5,
		false,
	},
	11: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		1,
		1,
		8,
		6,
		false,
	},
	12: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		6,
		6,
		8,
		5,
		false,
	},
	14: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		9,
		9,
		1,
		39,
		false,
	},
	15: {
		ebiten.NewImage(8, 8),
		2,
		FPos{0, 0},
		FPos{0, 0},
		2,
		2,
		15,
		6,
		false,
	},
}

func init() {
	//towerimg
	tower_img1, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower1.png")
	tower_images[1] = tower_img1
	tower_img2, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower2.png")
	tower_images[2] = tower_img2
	tower_img3, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower3.png")
	tower_images[3] = tower_img3
	tower_img4, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower4.png")
	tower_images[4] = tower_img4
	tower_img5, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower5.png")
	tower_images[5] = tower_img5
	tower_img6, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower6.png")
	tower_images[6] = tower_img6
	tower_img7, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower7.png")
	tower_images[7] = tower_img7
	tower_img8, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower8.png")
	tower_images[8] = tower_img8
	tower_img9, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower9.png")
	tower_images[9] = tower_img9
	tower_img10, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower10.png")
	tower_images[10] = tower_img10
	tower_img11, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower11.png")
	tower_images[11] = tower_img11
	tower_img12, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower12.png")
	tower_images[12] = tower_img12
	tower_img13, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower13.png")
	tower_images[13] = tower_img13
	tower_img14, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower14.png")
	tower_images[14] = tower_img14
	tower_img15, _, _ := ebitenutil.NewImageFromFile("./art/towers/tower15.png")
	tower_images[15] = tower_img15

	//towerupgradeimg
	tower_upgrade_img1, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower1.png")
	tower_upgrade_images[1] = tower_upgrade_img1
	ttower_upgrade_img2, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower2.png")
	tower_upgrade_images[2] = ttower_upgrade_img2
	tower_upgrade_img3, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower3.png")
	tower_upgrade_images[3] = tower_upgrade_img3
	tower_upgrade_img4, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower4.png")
	tower_upgrade_images[4] = tower_upgrade_img4
	tower_upgrade_img5, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower5.png")
	tower_upgrade_images[5] = tower_upgrade_img5
	tower_upgrade_img6, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower6.png")
	tower_upgrade_images[6] = tower_upgrade_img6
	tower_upgrade_img7, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower7.png")
	tower_upgrade_images[7] = tower_upgrade_img7
	tower_upgrade_img8, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower8.png")
	tower_upgrade_images[8] = tower_upgrade_img8
	tower_upgrade_img9, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower9.png")
	tower_upgrade_images[9] = tower_upgrade_img9
	tower_upgrade_img10, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower10.png")
	tower_upgrade_images[10] = tower_upgrade_img10
	tower_upgrade_img11, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower11.png")
	tower_upgrade_images[11] = tower_upgrade_img11
	tower_upgrade_img12, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower12.png")
	tower_upgrade_images[12] = tower_upgrade_img12
	tower_upgrade_img14, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower14.png")
	tower_upgrade_images[14] = tower_upgrade_img14
	tower_upgrade_img15, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade/tower15.png")
	tower_upgrade_images[15] = tower_upgrade_img15

	//towerupgradeupgradeimg
	tower_upgrade_upgrade_img1, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower1.png")
	tower_upgrade_upgrade_images[1] = tower_upgrade_upgrade_img1
	tower_upgrade_upgrade_img2, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower2.png")
	tower_upgrade_upgrade_images[2] = tower_upgrade_upgrade_img2
	tower_upgrade_upgrade_img3, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower3.png")
	tower_upgrade_upgrade_images[3] = tower_upgrade_upgrade_img3
	tower_upgrade_upgrade_img4, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower4.png")
	tower_upgrade_upgrade_images[4] = tower_upgrade_upgrade_img4
	tower_upgrade_upgrade_img5, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower5.png")
	tower_upgrade_upgrade_images[5] = tower_upgrade_upgrade_img5
	tower_upgrade_upgrade_img6, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower6.png")
	tower_upgrade_upgrade_images[6] = tower_upgrade_upgrade_img6
	tower_upgrade_upgrade_img7, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower7.png")
	tower_upgrade_upgrade_images[7] = tower_upgrade_upgrade_img7
	tower_upgrade_upgrade_img8, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower8.png")
	tower_upgrade_upgrade_images[8] = tower_upgrade_upgrade_img8
	tower_upgrade_upgrade_img9, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower9.png")
	tower_upgrade_upgrade_images[9] = tower_upgrade_upgrade_img9
	tower_upgrade_upgrade_img10, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower10.png")
	tower_upgrade_upgrade_images[10] = tower_upgrade_upgrade_img10
	tower_upgrade_upgrade_img11, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower11.png")
	tower_upgrade_upgrade_images[11] = tower_upgrade_upgrade_img11
	tower_upgrade_upgrade_img12, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower12.png")
	tower_upgrade_upgrade_images[12] = tower_upgrade_upgrade_img12
	tower_upgrade_upgrade_img14, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower14.png")
	tower_upgrade_upgrade_images[14] = tower_upgrade_upgrade_img14
	tower_upgrade_upgrade_img15, _, _ := ebitenutil.NewImageFromFile("./art/tower_upgrade_upgrade/tower15.png")
	tower_upgrade_upgrade_images[15] = tower_upgrade_upgrade_img15

	//outline
	tower_img_outline_1, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower1outline.png")
	tower_images_outlines[1] = tower_img_outline_1
	tower_img_outline_2, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower2outline.png")
	tower_images_outlines[2] = tower_img_outline_2
	tower_img_outline_3, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower3outline.png")
	tower_images_outlines[3] = tower_img_outline_3
	tower_img_outline_4, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower4outline.png")
	tower_images_outlines[4] = tower_img_outline_4
	tower_img_outline_5, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower5outline.png")
	tower_images_outlines[5] = tower_img_outline_5
	tower_img_outline_6, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower6outline.png")
	tower_images_outlines[6] = tower_img_outline_6
	tower_img_outline_7, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower7outline.png")
	tower_images_outlines[7] = tower_img_outline_7
	tower_img_outline_8, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower8outline.png")
	tower_images_outlines[8] = tower_img_outline_8
	tower_img_outline_9, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower9outline.png")
	tower_images_outlines[9] = tower_img_outline_9
	tower_img_outline_10, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower10outline.png")
	tower_images_outlines[10] = tower_img_outline_10
	tower_img_outline_11, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower11outline.png")
	tower_images_outlines[11] = tower_img_outline_11
	tower_img_outline_12, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower12outline.png")
	tower_images_outlines[12] = tower_img_outline_12
	tower_img_outline_13, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower13outline.png")
	tower_images_outlines[13] = tower_img_outline_13
	tower_img_outline_14, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower14outline.png")
	tower_images_outlines[14] = tower_img_outline_14
	tower_img_outline_15, _, _ := ebitenutil.NewImageFromFile("./art/tower_outlines/tower15outline.png")
	tower_images_outlines[15] = tower_img_outline_15

	//attack
	attack_img1, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack1.png")
	tower_attack[1] = Attack{attack_img1, 1, FPos{0, 0}, FPos{0, 0}, tower_attack[1].timer, tower_attack[1].o_timer, tower_attack[1].projectile_speed, tower_attack[1].damage, false}
	attack_img2, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack2.png")
	tower_attack[2] = Attack{attack_img2, 2, FPos{0, 0}, FPos{0, 0}, tower_attack[2].timer, tower_attack[2].o_timer, tower_attack[2].projectile_speed, tower_attack[2].damage, false}
	attack_img3, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack3.png")
	tower_attack[3] = Attack{attack_img3, 3, FPos{0, 0}, FPos{0, 0}, tower_attack[3].timer, tower_attack[3].o_timer, tower_attack[3].projectile_speed, tower_attack[3].damage, false}
	attack_img4, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack4.png")
	tower_attack[4] = Attack{attack_img4, 4, FPos{0, 0}, FPos{0, 0}, tower_attack[4].timer, tower_attack[4].o_timer, tower_attack[4].projectile_speed, tower_attack[4].damage, false}
	attack_img5, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack5.png")
	tower_attack[5] = Attack{attack_img5, 5, FPos{0, 0}, FPos{0, 0}, tower_attack[5].timer, tower_attack[5].o_timer, tower_attack[5].projectile_speed, tower_attack[5].damage, false}
	attack_img6, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack6.png")
	tower_attack[6] = Attack{attack_img6, 6, FPos{0, 0}, FPos{0, 0}, tower_attack[6].timer, tower_attack[6].o_timer, tower_attack[6].projectile_speed, tower_attack[6].damage, false}
	attack_img7, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack7.png")
	tower_attack[7] = Attack{attack_img7, 7, FPos{0, 0}, FPos{0, 0}, tower_attack[7].timer, tower_attack[7].o_timer, tower_attack[7].projectile_speed, tower_attack[7].damage, false}
	attack_img8, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack8.png")
	tower_attack[8] = Attack{attack_img8, 8, FPos{0, 0}, FPos{0, 0}, tower_attack[8].timer, tower_attack[8].o_timer, tower_attack[8].projectile_speed, tower_attack[8].damage, false}
	attack_img9, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack9.png")
	tower_attack[9] = Attack{attack_img9, 9, FPos{0, 0}, FPos{0, 0}, tower_attack[9].timer, tower_attack[9].o_timer, tower_attack[9].projectile_speed, tower_attack[9].damage, false}
	attack_img10, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack10.png")
	tower_attack[10] = Attack{attack_img10, 10, FPos{0, 0}, FPos{0, 0}, tower_attack[10].timer, tower_attack[10].o_timer, tower_attack[10].projectile_speed, tower_attack[10].damage, false}
	attack_img11, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack11.png")
	tower_attack[11] = Attack{attack_img11, 11, FPos{0, 0}, FPos{0, 0}, tower_attack[11].timer, tower_attack[11].o_timer, tower_attack[11].projectile_speed, tower_attack[11].damage, false}
	attack_img12, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack12.png")
	tower_attack[12] = Attack{attack_img12, 12, FPos{0, 0}, FPos{0, 0}, tower_attack[12].timer, tower_attack[12].o_timer, tower_attack[12].projectile_speed, tower_attack[12].damage, false}
	attack_img13, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack13.png")
	tower_attack[13] = Attack{attack_img13, 13, FPos{0, 0}, FPos{0, 0}, tower_attack[13].timer, tower_attack[13].o_timer, tower_attack[13].projectile_speed, tower_attack[13].damage, false}
	attack_img14, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack14.png")
	tower_attack[14] = Attack{attack_img14, 14, FPos{0, 0}, FPos{0, 0}, tower_attack[14].timer, tower_attack[14].o_timer, tower_attack[14].projectile_speed, tower_attack[14].damage, false}
	attack_img15, _, _ := ebitenutil.NewImageFromFile("./art/attacks/attack15.png")
	tower_attack[15] = Attack{attack_img15, 15, FPos{0, 0}, FPos{0, 0}, tower_attack[15].timer, tower_attack[15].o_timer, tower_attack[15].projectile_speed, tower_attack[15].damage, false}

	//upgraded attacks
	attack_upgrade_img1, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack1.png")
	tower_upgrade_attack[1] = Attack{attack_upgrade_img1, 1, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[1].timer, tower_upgrade_attack[1].o_timer, tower_upgrade_attack[1].projectile_speed, tower_upgrade_attack[1].damage, false}
	attack_upgrade_img2, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack2.png")
	tower_upgrade_attack[2] = Attack{attack_upgrade_img2, 2, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[2].timer, tower_upgrade_attack[2].o_timer, tower_upgrade_attack[2].projectile_speed, tower_upgrade_attack[2].damage, false}
	attack_upgrade_img3, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack3.png")
	tower_upgrade_attack[3] = Attack{attack_upgrade_img3, 3, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[3].timer, tower_upgrade_attack[3].o_timer, tower_upgrade_attack[3].projectile_speed, tower_upgrade_attack[3].damage, false}
	attack_upgrade_img4, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack4.png")
	tower_upgrade_attack[4] = Attack{attack_upgrade_img4, 4, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[4].timer, tower_upgrade_attack[4].o_timer, tower_upgrade_attack[4].projectile_speed, tower_upgrade_attack[4].damage, false}
	attack_upgrade_img5, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack5.png")
	tower_upgrade_attack[5] = Attack{attack_upgrade_img5, 5, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[5].timer, tower_upgrade_attack[5].o_timer, tower_upgrade_attack[5].projectile_speed, tower_upgrade_attack[5].damage, false}
	attack_upgrade_img6, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack6.png")
	tower_upgrade_attack[6] = Attack{attack_upgrade_img6, 6, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[6].timer, tower_upgrade_attack[6].o_timer, tower_upgrade_attack[6].projectile_speed, tower_upgrade_attack[6].damage, false}
	attack_upgrade_img7, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack7.png")
	tower_upgrade_attack[7] = Attack{attack_upgrade_img7, 7, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[7].timer, tower_upgrade_attack[7].o_timer, tower_upgrade_attack[7].projectile_speed, tower_upgrade_attack[7].damage, false}
	attack_upgrade_img8, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack8.png")
	tower_upgrade_attack[8] = Attack{attack_upgrade_img8, 8, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[8].timer, tower_upgrade_attack[8].o_timer, tower_upgrade_attack[8].projectile_speed, tower_upgrade_attack[8].damage, false}
	attack_upgrade_img9, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack9.png")
	tower_upgrade_attack[9] = Attack{attack_upgrade_img9, 9, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[9].timer, tower_upgrade_attack[9].o_timer, tower_upgrade_attack[9].projectile_speed, tower_upgrade_attack[9].damage, false}
	attack_upgrade_img10, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack10.png")
	tower_upgrade_attack[10] = Attack{attack_upgrade_img10, 10, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[10].timer, tower_upgrade_attack[10].o_timer, tower_upgrade_attack[10].projectile_speed, tower_upgrade_attack[10].damage, false}
	attack_upgrade_img11, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack11.png")
	tower_upgrade_attack[11] = Attack{attack_upgrade_img11, 11, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[11].timer, tower_upgrade_attack[11].o_timer, tower_upgrade_attack[11].projectile_speed, tower_upgrade_attack[11].damage, false}
	attack_upgrade_img12, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack12.png")
	tower_upgrade_attack[12] = Attack{attack_upgrade_img12, 12, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[12].timer, tower_upgrade_attack[12].o_timer, tower_upgrade_attack[12].projectile_speed, tower_upgrade_attack[12].damage, false}
	attack_upgrade_img14, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack14.png")
	tower_upgrade_attack[14] = Attack{attack_upgrade_img14, 14, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[14].timer, tower_upgrade_attack[14].o_timer, tower_upgrade_attack[14].projectile_speed, tower_upgrade_attack[14].damage, false}
	attack_upgrade_img15, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_attacks/attack15.png")
	tower_upgrade_attack[15] = Attack{attack_upgrade_img15, 15, FPos{0, 0}, FPos{0, 0}, tower_upgrade_attack[15].timer, tower_upgrade_attack[15].o_timer, tower_upgrade_attack[15].projectile_speed, tower_upgrade_attack[15].damage, false}

	//upgraded upgraded attacks
	attack_upgrade_upgrade_img1, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack1.png")
	tower_upgrade_upgrade_attack[1] = Attack{attack_upgrade_upgrade_img1, 1, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[1].timer, tower_upgrade_upgrade_attack[1].o_timer, tower_upgrade_upgrade_attack[1].projectile_speed, tower_upgrade_upgrade_attack[1].damage, false}
	attack_upgrade_upgrade_img2, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack2.png")
	tower_upgrade_upgrade_attack[2] = Attack{attack_upgrade_upgrade_img2, 2, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[2].timer, tower_upgrade_upgrade_attack[2].o_timer, tower_upgrade_upgrade_attack[2].projectile_speed, tower_upgrade_upgrade_attack[2].damage, false}
	attack_upgrade_upgrade_img3, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack3.png")
	tower_upgrade_upgrade_attack[3] = Attack{attack_upgrade_upgrade_img3, 3, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[3].timer, tower_upgrade_upgrade_attack[3].o_timer, tower_upgrade_upgrade_attack[3].projectile_speed, tower_upgrade_upgrade_attack[3].damage, false}
	attack_upgrade_upgrade_img4, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack4.png")
	tower_upgrade_upgrade_attack[4] = Attack{attack_upgrade_upgrade_img4, 4, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[4].timer, tower_upgrade_upgrade_attack[4].o_timer, tower_upgrade_upgrade_attack[4].projectile_speed, tower_upgrade_upgrade_attack[4].damage, false}
	attack_upgrade_upgrade_img5, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack5.png")
	tower_upgrade_upgrade_attack[5] = Attack{attack_upgrade_upgrade_img5, 5, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[5].timer, tower_upgrade_upgrade_attack[5].o_timer, tower_upgrade_upgrade_attack[5].projectile_speed, tower_upgrade_upgrade_attack[5].damage, false}
	attack_upgrade_upgrade_img6, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack6.png")
	tower_upgrade_upgrade_attack[6] = Attack{attack_upgrade_upgrade_img6, 6, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[6].timer, tower_upgrade_upgrade_attack[6].o_timer, tower_upgrade_upgrade_attack[6].projectile_speed, tower_upgrade_upgrade_attack[6].damage, false}
	attack_upgrade_upgrade_img7, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack7.png")
	tower_upgrade_upgrade_attack[7] = Attack{attack_upgrade_upgrade_img7, 7, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[7].timer, tower_upgrade_upgrade_attack[7].o_timer, tower_upgrade_upgrade_attack[7].projectile_speed, tower_upgrade_upgrade_attack[7].damage, false}
	attack_upgrade_upgrade_img8, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack8.png")
	tower_upgrade_upgrade_attack[8] = Attack{attack_upgrade_upgrade_img8, 8, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[8].timer, tower_upgrade_upgrade_attack[8].o_timer, tower_upgrade_upgrade_attack[8].projectile_speed, tower_upgrade_upgrade_attack[8].damage, false}
	attack_upgrade_upgrade_img9, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack9.png")
	tower_upgrade_upgrade_attack[9] = Attack{attack_upgrade_upgrade_img9, 9, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[9].timer, tower_upgrade_upgrade_attack[9].o_timer, tower_upgrade_upgrade_attack[9].projectile_speed, tower_upgrade_upgrade_attack[9].damage, false}
	attack_upgrade_upgrade_img10, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack10.png")
	tower_upgrade_upgrade_attack[10] = Attack{attack_upgrade_upgrade_img10, 10, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[10].timer, tower_upgrade_upgrade_attack[10].o_timer, tower_upgrade_upgrade_attack[10].projectile_speed, tower_upgrade_upgrade_attack[10].damage, false}
	attack_upgrade_upgrade_img11, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack11.png")
	tower_upgrade_upgrade_attack[11] = Attack{attack_upgrade_upgrade_img11, 11, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[11].timer, tower_upgrade_upgrade_attack[11].o_timer, tower_upgrade_upgrade_attack[11].projectile_speed, tower_upgrade_upgrade_attack[11].damage, false}
	attack_upgrade_upgrade_img12, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack12.png")
	tower_upgrade_upgrade_attack[12] = Attack{attack_upgrade_upgrade_img12, 12, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[12].timer, tower_upgrade_upgrade_attack[12].o_timer, tower_upgrade_upgrade_attack[12].projectile_speed, tower_upgrade_upgrade_attack[12].damage, false}
	attack_upgrade_upgrade_img14, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack14.png")
	tower_upgrade_upgrade_attack[14] = Attack{attack_upgrade_upgrade_img14, 14, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[14].timer, tower_upgrade_upgrade_attack[14].o_timer, tower_upgrade_upgrade_attack[14].projectile_speed, tower_upgrade_upgrade_attack[14].damage, false}
	attack_upgrade_upgrade_img15, _, _ := ebitenutil.NewImageFromFile("./art/upgraded_upgraded_attacks/attack15.png")
	tower_upgrade_upgrade_attack[15] = Attack{attack_upgrade_upgrade_img15, 15, FPos{0, 0}, FPos{0, 0}, tower_upgrade_upgrade_attack[15].timer, tower_upgrade_upgrade_attack[15].o_timer, tower_upgrade_upgrade_attack[15].projectile_speed, tower_upgrade_upgrade_attack[15].damage, false}
}

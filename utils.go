package main

type Pos struct {
	x int32
	y int32
}
type FPos struct {
	x float64
	y float64
}

func deg2rad(num float64) float64 {
	return (num * (180 / 3.14159))
}
func rad2deg(num float64) float64 {
	return (num * (3.14159 / 180))
}

// func collide(p1, s1, p2, s2 Pos) bool {
// 	return (p1.x < p2.x+s2.x &&
// 		p1.x+s1.x > p2.x &&
// 		p1.y < p2.y+s2.y &&
// 		p1.y+s1.y > p2.y)
// }

func fcollide(p1, s1, p2, s2 FPos) bool {
	return (p1.x < p2.x+s2.x &&
		p1.x+s1.x > p2.x &&
		p1.y < p2.y+s2.y &&
		p1.y+s1.y > p2.y)
}

func BubbleSort(array []Enemy) []Enemy {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j].travelled > array[j+1].travelled {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

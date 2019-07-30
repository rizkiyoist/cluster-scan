package main

import "fmt"

type cluster struct {
	x int64
	y int64
	q int
}

func main() {
	kitchenC := cluster{
		x: 1,
		y: 1,
		q: 3,
	}
	kitchenPoints, _ := getKitchenArea(kitchenC, 4)
	for _, kp := range kitchenPoints {
		fmt.Println(kp.x, kp.y, kp.q)
	}
}

func getKitchenArea(kc cluster, size int64) ([]cluster, error) {
	var c []cluster
	var temp cluster
	// get the rest of 2x2 cluster
	temp.x = kc.x
	temp.y = kc.y
	for q := 1; q <= 4; q++ {
		temp.q = q
		c = append(c, temp)
	}

	return c, nil
}

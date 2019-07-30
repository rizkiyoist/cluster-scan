package main

import "fmt"

type cluster struct {
	x int64
	y int64
	q int
}

type kitchen struct {
	x int64
	y int64
	q int
}

type diagonalCluster struct {
	xc int64
	yc int64
	qc int
	xf int64
	yf int64
	qf int
}

func main() {
	k := new(kitchen)
	k.q = 1
	k.x = 0
	k.y = 0
	kitchenPoints, _ := k.getKitchenArea(6)
	for _, kp := range kitchenPoints {
		fmt.Println(kp.x, kp.y, kp.q)
	}
}

func (kc kitchen) getKitchenArea(size int) ([]cluster, error) {
	var c []cluster
	var temp cluster
	// get the 2x2 cluster
	temp.x = kc.x
	temp.y = kc.y
	for q := 1; q <= 4; q++ {
		temp.q = q
		c = append(c, temp)
	}
	// get the rest of cluster
	roc, _ := kc.scanBySize(size)
	c = append(c, roc...)
	return c, nil
}

// scan every ring
func (kc kitchen) scanBySize(size int) ([]cluster, error) {
	// var c cluster
	var cs []cluster
	var cc cluster
	var rc cluster
	var ccDir string
	var rcDir string
	//current cluster is kitchen cluster
	cc.x = kc.x
	cc.y = kc.y
	cc.q = kc.q
	//reverse cluster is the opposite quadrant of current cluster
	rc.x = kc.x
	rc.y = kc.y
	switch kc.q {
	case 1:
		rc.q = 3
	case 2:
		rc.q = 4
	case 3:
		rc.q = 1
	case 4:
		rc.q = 2
	}

	//foreach size, get the diagonal then scan half ring
	for curSize := 3; curSize <= size; curSize++ {
		fmt.Printf("scanning size %v x %v \n", curSize, curSize)
		switch kc.q {
		case 1:
			ccDir = "upright"
			rcDir = "downleft"
		case 2:
			ccDir = "upleft"
			rcDir = "downright"
		case 3:
			ccDir = "downleft"
			rcDir = "upright"
		case 4:
			ccDir = "downright"
			rcDir = "upleft"
		}
		switch curSize % 2 {
		case 0:
			rc, _ = rc.getOneDiagonal(rcDir)
			fmt.Println("getting reverse diagonal", rc.x, rc.y, rc.q)
			cs = append(cs, rc)
			// c, _ = cc.scanRing(diagonal, curSize)
			// cs = append(cs, c...)
		case 1:
			cc, _ = cc.getOneDiagonal(ccDir)
			fmt.Println("getting diagonal", cc.x, cc.y, cc.q)
			cs = append(cs, cc)
			// c, _ = cc.scanRing(diagonal, curSize)
			// cs = append(cs, c...)
		}

	}
	return cs, nil
}

func (cc cluster) getOneDiagonal(direction string) (cluster, error) {
	var c cluster
	// check quadrant of current cluster, then get one diagonal based on direction
	switch direction {
	case "upright":
		// get upright
		c = getOne(cc, "up")
		c = getOne(c, "right")
	case "upleft":
		// get upleft
		c = getOne(cc, "up")
		c = getOne(c, "left")
	case "downleft":
		// get downleft
		c = getOne(cc, "down")
		c = getOne(c, "left")
	case "downright":
		// get downright
		c = getOne(cc, "down")
		c = getOne(c, "right")
	}
	return c, nil
}

func getOne(cc cluster, direction string) cluster {
	var c cluster
	c = cc
	switch direction {
	case "up":
		switch cc.q {
		case 1:
			c.y = cc.y + 1
			c.q = 4
		case 2:
			c.y = cc.y + 1
			c.q = 3
		case 3:
			c.q = 2
		case 4:
			c.q = 1
		}
	case "down":
		switch cc.q {
		case 1:
			c.q = 4
		case 2:
			c.q = 3
		case 3:
			c.q = 2
			c.y = cc.y - 1
		case 4:
			c.q = 1
			c.y = cc.y - 1
		}
	case "left":
		switch cc.q {
		case 1:
			c.q = 2
		case 2:
			c.q = 1
			c.x = cc.x - 1
		case 3:
			c.q = 2
			c.x = cc.x - 1
		case 4:
			c.q = 3
		}
	case "right":
		switch cc.q {
		case 1:
			c.q = 2
			c.x = cc.x + 1
		case 2:
			c.q = 1
		case 3:
			c.q = 4
		case 4:
			c.q = 3
			c.x = cc.x + 1
		}
	}
	return c
}

func (cc cluster) scanRing(diagonal cluster, size int) ([]cluster, error) {
	var cs []cluster
	return cs, nil
}

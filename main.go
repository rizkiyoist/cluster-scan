package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type cluster struct {
	x int
	y int
	q int
}

type kitchen struct {
	x int
	y int
	q int
}

func main() {
	k := new(kitchen)
	k.q = 4
	k.x = 1
	k.y = 1
	kitchenPoints, _ := k.getKitchenArea(8)
	// Create canvas image
	img := image.NewRGBA(image.Rect(-10, -10, 50, 50))
	var count int
	for c, kp := range kitchenPoints {

		// Draw a red dot at x y
		switch kp.q {
		case 1:
			img.Set(kp.x, kp.y, color.RGBA{255, 0, 0, 255})
		case 2:
			img.Set(kp.x, kp.y, color.RGBA{0, 255, 0, 255})
		case 3:
			img.Set(kp.x, kp.y, color.RGBA{0, 0, 255, 255})
		case 4:
			img.Set(kp.x, kp.y, color.RGBA{255, 255, 0, 255})
		}
		fmt.Println(kp.x, kp.y, kp.q)
		count = c + 1
	}
	fmt.Println("there are", count, "nodes")
	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
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
	var tempCs []cluster
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
			// reverse diagonal
			rc, _ = rc.getOneDiagonal(rcDir)
			cs = append(cs, rc)
			tempCs, _ = rc.scanRing(curSize, rcDir)
			cs = append(cs, tempCs...)
		case 1:
			// diagonal
			cc, _ = cc.getOneDiagonal(ccDir)
			cs = append(cs, cc)
			tempCs, _ = cc.scanRing(curSize, ccDir)
			cs = append(cs, tempCs...)
		}

	}
	return cs, nil
}

func (cc cluster) getOneDiagonal(direction string) (cluster, error) {
	var c cluster
	// check quadrant of current cluster, then get one diagonal based on direction
	switch direction {
	case "upright":
		c = getOne(cc, "up")
		c = getOne(c, "right")
	case "upleft":
		c = getOne(cc, "up")
		c = getOne(c, "left")
	case "downleft":
		c = getOne(cc, "down")
		c = getOne(c, "left")
	case "downright":
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
			c.q = 4
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

func (cc cluster) scanRing(size int, start string) ([]cluster, error) {
	var cs []cluster
	cx := cc
	cy := cc
	for i := 1; i < size; i++ {
		switch start {
		//scan from
		case "upright":
			cy = getOne(cy, "down")
			cs = append(cs, cy)
			cx = getOne(cx, "left")
			fmt.Println("get one left ", cx.x, cx.y, cx.q)
			cs = append(cs, cx)
		case "upleft":
			cy = getOne(cy, "down")
			cs = append(cs, cy)
			cx = getOne(cx, "right")
			cs = append(cs, cx)
		case "downleft":
			cx = getOne(cx, "right")
			cs = append(cs, cx)
			cy = getOne(cy, "up")
			cs = append(cs, cy)
		case "downright":
			cy = getOne(cy, "up")
			cs = append(cs, cy)
			cx = getOne(cx, "left")
			cs = append(cs, cx)
		}
	}
	return cs, nil
}

package main

import (
	"fmt"
	"math"
)

const (
	width          int     = 160
	height         int     = 44
	cubeWidth      float64 = 10
	bufferSize     int     = width * height
	background     string  = " "
	objectDistance int     = 60
	K1             float64 = 40
	speed          float64 = 0.5
)

var (
	ooz, x, y, z float64
	xp, yp, idx  int
	buff         []string
	zBuff        []float64
	A, B, C      float64
)

func main() {
	fmt.Print("\x1b[2J")
	buff = make([]string, int(bufferSize))
	zBuff = make([]float64, int(bufferSize))
	for {
		memSet(buff, background)
		memSetFloat(zBuff, 0)
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += speed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += speed {
				calculateSurface(cubeX, cubeY, -cubeWidth, "#")
				calculateSurface(cubeX, cubeY, -cubeWidth, "@")
				calculateSurface(cubeWidth, cubeY, cubeX, "$")
				calculateSurface(-cubeWidth, cubeY, -cubeX, "~")
				calculateSurface(-cubeX, cubeY, cubeWidth, "#")
				calculateSurface(cubeX, -cubeWidth, -cubeY, ";")
				calculateSurface(cubeX, cubeWidth, cubeY, "+")
			}
		}
		fmt.Print("\x1b[H")
		for k := 0; k < width*height; k++ {
			c := k % width
			if c != 0 {
				fmt.Printf("%s", buff[k])
			} else {
				fmt.Printf("\n")
			}

		}
		A += 0.005
		B += 0.005
		C += 0.01
	}

}

func calculateX(i, j, k float64) float64 {
	return j*math.Sin(A)*math.Sin(B)*math.Cos(C) - k*math.Cos(A)*math.Sin(B)*math.Cos(C) + j*math.Cos(A)*math.Sin(C) + k*math.Sin(A)*math.Sin(C) + i*math.Cos(B)*math.Cos(C)

}

func calculateY(i, j, k float64) float64 {
	return j*math.Cos(A)*math.Cos(C) + k*math.Sin(A)*math.Cos(C) - j*math.Sin(A)*math.Sin(B)*math.Sin(C) + k*math.Cos(A)*math.Sin(B)*math.Sin(C) - i*math.Cos(B)*math.Sin(C)
}

func calculateZ(i, j, k float64) float64 {
	return k*math.Cos(A)*math.Cos(B) - j*math.Sin(A)*math.Cos(B) + i*math.Sin(B)
}

func calculateSurface(cubeX, cubeY, cubeZ float64, ch string) {
	x = calculateX(cubeX, cubeY, cubeZ)
	y = calculateY(cubeX, cubeY, cubeZ)
	z = calculateZ(cubeX, cubeY, cubeZ) + float64(objectDistance)

	ooz = 1 / z
	xp = int(float64(width)/2 + K1*ooz*x*2)
	yp = int(float64(height)/2 + K1*ooz*y)
	idx = xp + yp*width
	if idx >= 0 && idx < width*height {
		if ooz > zBuff[idx] {
			zBuff[idx] = ooz
			buff[idx] = ch
		}
	}
}

// memSet fills map's space with a given character. Mimics the behavior of C memset function
func memSet(a []string, v string) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

// memSetFloat fills map's space with a given float. Mimics the behavior of C memset function
func memSetFloat(a []float64, v float64) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

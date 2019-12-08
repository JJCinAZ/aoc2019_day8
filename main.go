package main

import (
	"cloud.google.com/aoc2019/day8/sif"
	"fmt"
	"io/ioutil"
)

func main() {
	part1()
	part2()
}

func part1() {
	bytes, err := ioutil.ReadFile("image.dat")
	if err != nil {
		panic(err)
	}
	img, err := sif.BuildImage(string(bytes), 25, 6)
	fewestZerosLayer := 0
	fewestZerosCount := 999999
	for i := range img.LayerInfo {
		if img.LayerInfo[i].ColorCounts[0] < fewestZerosCount {
			fewestZerosCount = img.LayerInfo[i].ColorCounts[0]
			fewestZerosLayer = i
		}
	}
	fmt.Println(img.LayerInfo[fewestZerosLayer].ColorCounts[1] * img.LayerInfo[fewestZerosLayer].ColorCounts[2])
}

func part2() {
	bytes, err := ioutil.ReadFile("image.dat")
	if err != nil {
		panic(err)
	}
	img, err := sif.BuildImage(string(bytes), 25, 6)
	newImg := make([]byte, img.Width * img.Height)
	for i := 0; i < len(newImg); i++ {
		newImg[i] = 2
	}
	offset := 0
	layerCount := 0
	for _, p := range img.Pixels {
		if newImg[offset] == 2 {
			newImg[offset] = p
		}
		offset++
		if offset == img.Width*img.Height {
			offset = 0
			layerCount++
		}
	}
	printImg(newImg, img.Width, img.Height)
}

func printImg(pixels []byte, w, h int) {
	i := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := "  "
			if pixels[i] == 1 {
				c = "🀫🀫"
			}
			fmt.Print(c)
			i++
		}
		fmt.Println("")
	}
}
package sif

import "fmt"

type LayerInfo struct {
	ColorCounts [10]int
}

type Image struct {
	Pixels        []byte
	Width, Height int
	Layers        int
	LayerInfo     []LayerInfo
}

func BuildImage(input string, width int, height int) (Image, error) {
	var (
		img           Image
		bytesPerLayer int
	)
	bytesPerLayer = width * height
	if len(input) % bytesPerLayer != 0 {
		return img, fmt.Errorf("input data length (%d) is not a multiple of the layer size (%d)",
			len(input), bytesPerLayer)
	}
	img.Layers = len(input) / bytesPerLayer
	img.Pixels = make([]byte, len(input))
	for j, c := range input {
		if c < '0' || c > '9' {
			return img, fmt.Errorf("invalid character %c in input at offset %d", c, j)
		}
		img.Pixels[j] = byte(c - '0')
	}
	img.Width = width
	img.Height = height
	curLayer := 0
	count := 0
	img.LayerInfo = make([]LayerInfo, img.Layers)
	for _, c := range img.Pixels {
		img.LayerInfo[curLayer].ColorCounts[c]++
		count++
		if count == bytesPerLayer {
			count = 0
			curLayer++
		}
	}
	return img, nil
}


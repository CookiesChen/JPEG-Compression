package JPEG

import (
	"fmt"
	"image/jpeg"
	"os"
)

type yuv struct {
	y int
	u int
	v int
}

func Exec(){
	file, err := os.Open("img/动物照片.jpg")
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	var arr [][]yuv
	for i := 0; i < bounds.Dx(); i++ {
		tmp := make([]yuv, bounds.Dy())
		arr = append(arr, tmp)
	}

	// YUV颜色转换
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			r, g, b, _ := img.At(i,j).RGBA()
			r = r/256
			g = g/256
			b = b/256
			y, u, v := YUV(int(r),int(g),int(b))
			setColor(&arr[i][j], y, u, v)
		}
	}

	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}

}

func setColor(c *yuv, y int, u int, v int)  {
	c.y = y
	c.u = u
	c.v = v
}
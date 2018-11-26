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

var(
	width int
	heigth int
)

type nodeF struct {
	yF int
	uF int
	vF int
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
	width = bounds.Dx()
	heigth = bounds.Dy()
	var arr [][]yuv
	var F [][]nodeF
	for i := 0; i < width; i++ {
		tmp := make([]yuv, heigth)
		tempF := make([]nodeF, heigth)
		arr = append(arr, tmp)
		F = append(F, tempF)
	}

	// YUV颜色转换
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			r, g, b, _ := img.At(i,j).RGBA()
			r = r/256
			g = g/256
			b = b/256
			y, u, v := YUV(int(r),int(g),int(b))
			setColor(&arr[i][j], y, u, v)
		}
	}

	// 二次采样
	TwiceSample(arr[:][:])

	// DCT变换
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {
			DCT(i, j, arr[i*8:(i+1)*8][j*8:(j+1)*8],F[i*8:(i+1)*8][j*8:(j+1)*8])
		}
	}


	for i := 0; i < width; i++ {
		for j := 0; j < 1; j++ {
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}

	for i := 0; i < width; i++ {
		for j := 0; j < 1; j++ {
			fmt.Print(F[i][j])
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
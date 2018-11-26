package JPEG

import (
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
	extendx int
	extendy int
	AC []nodeAC
)

type nodeF struct {
	yF int
	uF int
	vF int
}

type nodeAC struct {
	zS int // #-zero-to-skip
	next int // next non-zero value
}

type nodeDC struct {
	s int // SIZE
	a int // AMPLITUDE
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
	extendx = 8 - width % 8
	if extendx == 8 {
		extendx = 0
	}
	extendy = 8 - heigth % 8
	if extendy == 8 {
		extendy = 0
	}
	for i := 0; i < width + extendx; i++ {
		tmp := make([]yuv, heigth + extendy)
		tempF := make([]nodeF, heigth + extendy)
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

			var f [8][8]yuv
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					f[k][t] = arr[i*8 + k][j*8 + t]
				}
			}

			FDCT := DCT(f)

			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					F[i*8 + k][j*8 + t] = FDCT[k][t]
				}
			}
		}
	}

	// 量化
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {

			var f [8][8]nodeF
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					f[k][t] = F[i*8 + k][j*8 + t]
				}
			}

			QF := quantification(f)

			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					F[i*8 + k][j*8 + t] = QF[k][t]
				}
			}
		}
	}


	// RLC
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {
			var f [8][8]nodeF
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					f[k][t] = F[i*8 + k][j*8 + t]
				}
			}
			RLC(f)
		}
	}

	// DPCM
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {
			var f []nodeF
			f = append(f, F[i*8][j*8])
			DPCM(f)
		}
	}
}

func setColor(c *yuv, y int, u int, v int)  {
	c.y = y
	c.u = u
	c.v = v
}
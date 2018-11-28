package JPEG

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

func Decode(F [][]nodeF)  {
	file, err := os.Open("img/动物卡通图片.jpg")
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
	extendx = 8 - width % 8
	if extendx == 8 {
		extendx = 0
	}
	extendy = 8 - heigth % 8
	if extendy == 8 {
		extendy = 0
	}
	var arr [][]yuv
	for i := 0; i < width + extendx; i++ {
		tmp := make([]yuv, heigth + extendy)
		arr = append(arr, tmp)
	}

	// 反量化
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {

			var f [8][8]nodeF
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					f[k][t] = F[i*8 + k][j*8 + t]
				}
			}

			antiQF := antiQuantification(f)

			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					F[i*8 + k][j*8 + t] = antiQF[k][t]
				}
			}
		}
	}

	// 逆DCT
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {

			var f [8][8]nodeF
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					f[k][t] = F[i*8 + k][j*8 + t]
				}
			}

			antiFDCT := antiDCT(f)

			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					arr[i*8 + k][j*8 + t] = antiFDCT[k][t]
				}
			}
		}
	}

	newImage := image.NewRGBA(image.Rect(0,0, width, heigth))
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			r, g, b := antiYUV(arr[i][j].y, arr[i][j].u, arr[i][j].v)
			newImage.Set(i, j, color.RGBA{R:uint8(r), G:uint8(g), B:uint8(b)})
		}
	}

	outputfile, _ := os.Create("new.jpg")
	jpeg.Encode(outputfile, newImage, &jpeg.Options{Quality:100})

	outputPng, _ := os.Create("动物图片.png")
	png.Encode(outputPng, img)
}
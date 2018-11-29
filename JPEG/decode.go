package JPEG

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"os"
)

func Decode(F [][]nodeF)  {
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

	// 均方差mse
	mse := 0
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			r, g, b := antiYUV(arr[i][j].y, arr[i][j].u, arr[i][j].v)
			r1, g1, b1, _ := img.At(i,j).RGBA()
			R, G, B := int(r1/257), int(g1/257), int(b1/257)
			mse += (R - r)*(R - r)
			mse += (G - g)*(G - g)
			mse += (B - b)*(B - b)
			newImage.Set(i, j, color.RGBA{R:uint8(r), G:uint8(g), B:uint8(b)})
		}
	}

	mse = mse / width / heigth

	fmt.Println("jpeg mse:", mse)

	outputfile, _ := os.Create("new.jpg")
	jpeg.Encode(outputfile, newImage, &jpeg.Options{Quality:100})


	mse = 0
	file1, err := os.Open("img/动物卡通图片.gif")
	img1, err := gif.Decode(file1)
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			r, g, b, _ := img.At(i,j).RGBA()
			R, G, B := int(r/257), int(g/257), int(b/257)
			r1, g1, b1, _ := img1.At(i,j).RGBA()
			R1, G1, B1 := int(r1/257), int(g1/257), int(b1/257)
			mse += (R - R1)*(R - R1)
			mse += (G - G1)*(G - G1)
			mse += (B - B1)*(B - B1)
		}
	}
	mse = mse / width / heigth
	fmt.Println("gif mse:", mse)
}
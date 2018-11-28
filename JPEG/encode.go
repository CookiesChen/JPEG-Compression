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
	extendx int
	extendy int
)

type nodeF struct {
	yF int
	uF int
	vF int
}

type nodeAC struct {
	rs int // (runlength, size)
	next []int // next non-zero value
}

type nodeDC struct {
	s int // SIZE
	a []int // AMPLITUDE
}

func Exec()(F [][]nodeF){
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
	var arr [][]yuv
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
			R, G, B := float32(r/257), float32(g/257), float32(b/257)
			y, u, v := YUV(R, G, B)
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

	//// RLC
	var AC []nodeAC
	for i := 0; i < 1; i++ {
		for j:= 0; j < 1; j++ {
			var f [8][8]nodeF
			for k := 0; k < 8; k++ {
				for t := 0; t < 8; t++ {
					f[k][t] = F[i*8 + k][j*8 + t]
				}
			}
			AC  = append(AC, RLC(f)...)
		}
	}

	//// DPCM
	var f []nodeF
	for i := 0; i < width/8; i++ {
		for j:= 0; j < heigth/8; j++ {
			f = append(f, F[i*8][j*8])
		}
	}
	DC := DPCM(f)

	//// 哈夫曼编码
	fileObj,err := os.OpenFile("text.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	var data []byte

	m := make(map[int]int)
	for i, length := 0, len(AC); i < length; i++ {
		m[AC[i].rs]++
	}
	ACTable := huffman(m)
	for i, length := 0, len(AC); i < length; i++ {
		code := ACTable[AC[i].rs]
		count := 0
		var onebyte byte
		for j, len1 := 0, len(code); j < len1; j++ {
			count++
			onebyte = onebyte << 1 + code[j]-'0'
			if count == 8 {
				count = 0
				data = append(data, onebyte)
			}
		}
		count = 0
		for j, len1 := 0, len(AC[i].next); j < len1; j++ {
			count++
			onebyte = onebyte << 1 + byte(AC[i].next[j])
			if count == 8 {
				count = 0
				data = append(data, onebyte)
			}
		}
	}
	m = make(map[int]int)
	for i, length := 0, len(DC); i < length; i++ {
		m[DC[i].s]++
	}
	DCTable := huffman(m)
	for i, length := 0, len(DC); i < length; i++ {
		code := DCTable[DC[i].s]
		count := 0
		var onebyte byte
		for j, len1 := 0, len(code); j < len1; j++ {
			count++
			onebyte = onebyte << 1 + code[j]-'0'
			if count == 8 {
				count = 0
				data = append(data, onebyte)
			}
		}
		count = 0
		for j, len1 := 0, len(DC[i].a); j < len1; j++ {
			count++
			onebyte = onebyte << 1 + byte(DC[i].a[j])
			if count == 8 {
				count = 0
				data = append(data, onebyte)
			}
		}
	}
	fileObj.Write(data)

	return F
}

func setColor(c *yuv, y int, u int, v int)  {
	c.y = y
	c.u = u
	c.v = v
}
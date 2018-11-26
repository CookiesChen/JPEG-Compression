package JPEG

import "fmt"

func RLC(F [8][8]nodeF)  {
	zY, zU, zV := zScan(F)
	changeToAC(zY)
	changeToAC(zU)
	changeToAC(zV)
}

func zScan(F [8][8]nodeF)(zY []int, zU []int, zV []int) {
	for k := 0; k < 8; k++ {
		for t := 0; t < 8; t++ {
			fmt.Print(F[t][k].yF)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	flag := false
	x := 0
	y := 0
	for i := 0; i < 64; i++ {
		if !flag {
			zY = append(zY, F[x][y].yF)
			zU = append(zU, F[x][y].uF)
			zV = append(zV, F[x][y].vF)
			if x + 1 >= 8 {
				flag = !flag
				y = y + 1
			} else if y <= 0 {
				flag = !flag
				x = x + 1
			} else {
				x = x + 1
				y = y - 1
			}
		} else {
			zY = append(zY, F[x][y].yF)
			zU = append(zU, F[x][y].uF)
			zV = append(zV, F[x][y].vF)
			if y + 1 >= 8 {
				flag = !flag
				x = x + 1
			} else if x <= 0 {
				flag = !flag
				y = y + 1
			} else {
				x = x - 1
				y = y + 1
			}
		}
	}

	return zY, zU, zV

}

func changeToAC(arr []int)  {
	zeroNum := 0
	for i := 1; i < 64; i++ {
		if arr[i] == 0{
			zeroNum++
		} else {
			AC = append(AC, nodeAC{zeroNum, arr[i]})
		}
	}
	AC = append(AC, nodeAC{0, 0})
}
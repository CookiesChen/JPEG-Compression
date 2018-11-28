package JPEG

import "fmt"

func RLC(F [8][8]nodeF)(AC []nodeAC)  {
	zY, zU, zV := zScan(F)
	AC = append(AC, changeToAC(zY)...)
	fmt.Println(AC)
	AC = append(AC, changeToAC(zU)...)
	AC = append(AC, changeToAC(zV)...)
	return AC
}

func zScan(F [8][8]nodeF)(zY []int, zU []int, zV []int) {
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

func changeToAC(arr []int) (AC []nodeAC) {
	fmt.Println(arr)
	zeroNum := 0
	for i := 1; i < 64; i++ {
		if arr[i] == 0 && zeroNum <= 15{
			zeroNum++
		} else {
			if zeroNum > 15 {
				extendTimes := zeroNum / 16
				for j := 0; j < extendTimes; j++{
					AC = append(AC, nodeAC{240, append([]int{}, 0)})
				}
				zeroNum = zeroNum % 16
			}
			num := arr[i]
			if num < 0 {
				count, dst := intToBits(-num)
				AC = append(AC, nodeAC{getSymbol1(zeroNum, count), dst})
				// 反转
			} else if num > 0{
				count, dst := intToBits(num)
				AC = append(AC, nodeAC{getSymbol1(zeroNum, count), dst})
			} else {
				dst := append([]int{}, 0)
				AC = append(AC, nodeAC{getSymbol1(zeroNum, 0), dst})
			}
			zeroNum = 0
		}
	}
	return append(AC, nodeAC{0, append([]int{}, 0)})
}

func getSymbol1(zeroNum int, count int) (rs int){
	fmt.Println(zeroNum, count)
	dst := make([]int, 0)
	for i := 0; i < 4; i++ {
		move := uint(3 - i)
		dst = append(dst, int((zeroNum>>move)&1))
	}
	for i := 0; i < 4; i++ {
		move := uint(3 - i)
		dst = append(dst, int((count>>move)&1))
	}
	for i := 0; i < 8; i ++ {
		rs = rs * 2 + dst[i]
	}
	return rs
}
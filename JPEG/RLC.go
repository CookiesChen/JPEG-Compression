package JPEG

func RLC(F [8][8]nodeF)(AC []nodeAC)  {
	zY, zU, zV := zScan(F)
	AC = append(AC, changeToAC(zY)...)
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
	zeroNum := 0
	for i := 1; i < 64; i++ {
		if arr[i] == 0{
			zeroNum++
		} else {
			AC = append(AC, nodeAC{zeroNum, arr[i]})
		}
	}
	return append(AC, nodeAC{0, 0})
}
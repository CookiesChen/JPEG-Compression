package JPEG

func DPCM(f []nodeF)(DC []nodeDC)  {
	var yDC []int
	var uDC []int
	var vDC []int
	var d []int
	yDC = append(yDC, f[0].yF)
	uDC = append(uDC, f[0].uF)
	vDC = append(vDC, f[0].vF)
	for i, length:= 1, len(f); i < length; i++ {
		yDC = append(yDC, f[i].yF - f[i-1].yF)
		uDC = append(uDC, f[i].uF - f[i-1].uF)
		vDC = append(vDC, f[i].vF - f[i-1].vF)
	}
	d = append(d, yDC...)
	d = append(d, uDC...)
	d = append(d, vDC...)
	DC = dToDC(d)
	return DC
}

func dToDC(d []int)(DC []nodeDC) {
	for v := range d {
		num := d[v]
		if num < 0 {
			num = -num
			count, dst := intToBits(num)
			DC = append(DC, nodeDC{count, inverse(dst)})
			// 反转
		} else if num > 0{
			count, dst := intToBits(num)
			DC = append(DC, nodeDC{count, dst})
		} else {
			dst := append([]int{}, 0)
			DC = append(DC, nodeDC{0, dst})
		}
	}
	return DC
}

func intToBits(v int) (count int, dst []int) {
	for ; v != 0;{
		count++
		dst = append(dst, int(v % 2))
		v = v / 2
	}
	for i, j := 0, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return count, dst
}

func inverse(dst []int) []int  {
	for v := range dst {
		dst[v] = dst[v]^1
	}
	return dst
}
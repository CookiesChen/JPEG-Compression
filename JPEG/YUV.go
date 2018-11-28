package JPEG

func YUV(r float32,g float32,b float32)(y int, u int, v int)  {
	y = int(uint8(0.299*r + 0.587*g + 0.114*b))
	u = int(uint8(-0.1687*r - 0.3313*g + 0.5*b + 128))
	v = int(uint8(0.5*r - 0.4187*g - 0.0813*b + 128))
	if y < 0 {
		y = 0
	} else if r > 255 {
		y = 255
	}
	if u < 0 {
		u = 0
	} else if b > 255 {
		u = 255
	}
	if v < 0 {
		v = 0
	} else if b > 255 {
		v = 255
	}
	return y, u, v
}
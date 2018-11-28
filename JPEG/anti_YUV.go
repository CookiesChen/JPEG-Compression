package JPEG

func antiYUV(y int, u int, v int)(r int, g int, b int)  {
	r = y + int(1.402 * float32(v - 128))
	g = y - int(0.34414 * float32(u - 128)) - int(0.71414 * float32(v - 128))
	b = y + int(1.772 * float32(u - 128))
	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}
	return r, g, b
}
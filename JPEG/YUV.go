package JPEG

func YUV(r int,g int,b int)(y int, u int, v int)  {
	y = int(0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b))
	u = int(-0.169*float32(r) - 0.331*float32(g) + 0.5*float32(b)) + 128
	v = int(0.5*float32(r) - 0.419*float32(g) - 0.081*float32(b)) + 128
	return y, u, v
}
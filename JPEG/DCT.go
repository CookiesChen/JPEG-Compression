package JPEG

import "math"

func DCT(u int, v int, f [][]yuv, F [][]nodeF)  {
	var sumY float64 = 0.0
	var sumU float64 = 0.0
	var sumV float64 = 0.0
	for i := 0; i < 7; i++ {
		var sumy float64 = 0.0
		var sumu float64 = 0.0
		var sumv float64 = 0.0
		for j := 0; j < 7; j++ {
			factor1 := math.Cos(float64((2*float64(i))*float64(u)*math.Pi)/16)
			factor2 := math.Cos(float64((2*float64(j))*float64(v)*math.Pi)/16)
			sumy += factor1 * factor2 * float64(f[i][j].y)
			sumu += factor1 * factor2 * float64(f[i][j].u)
			sumv += factor1 * factor2 * float64(f[i][j].v)
		}
		sumY += sumy
		sumU += sumu
		sumV += sumv
	}
	F[u][v].yF = int(sumY)
	F[u][v].uF = int(sumU)
	F[u][v].vF = int(sumV)
}
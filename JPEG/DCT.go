package JPEG

import (
	"math"
)

func DCT(f [8][8]yuv) (F [8][8]nodeF) {
	for u := 0; u < 8; u++ {
		for v:= 0; v < 8; v++ {
			var cU float64
			var cV float64
			if u == 0 {
				cU = math.Sqrt(2) / 2.0
			} else {
				cU = 1
			}
			if v == 0 {
				cV = math.Sqrt(2) / 2.0
			} else {
				cV = 1
			}
			var sumY float64
			var sumU float64
			var sumV float64
			for i := 0; i < 8; i++ {
				var sumy float64
				var sumu float64
				var sumv float64
				for j := 0; j < 8; j++ {
					factor1 := math.Cos(float64((2*i+1)*u)*float64(math.Pi)/16.0)
					factor2 := math.Cos(float64((2*j+1)*v)*float64(math.Pi)/16.0)
					sumy += factor1 * factor2 * float64(f[i][j].y)
					sumu += factor1 * factor2 * float64(f[i][j].u)
					sumv += factor1 * factor2 * float64(f[i][j].v)
				}
				sumY += sumy
				sumU += sumu
				sumV += sumv
			}
			F[u][v].yF = int(math.Round(sumY * cU * cV / 4.0))
			F[u][v].uF = int(math.Round(sumU * cU * cV / 4.0))
			F[u][v].vF = int(math.Round(sumV * cU * cV / 4.0))
		}
	}
	return F
}
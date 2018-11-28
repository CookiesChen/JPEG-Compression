package JPEG

import (
	"math"
)

func antiDCT(F [8][8]nodeF) (f [8][8]yuv) {
	for i := 0; i < 8; i++ {
		for j:= 0; j < 8; j++ {
			var sumY float64
			var sumU float64
			var sumV float64
			for u := 0; u < 8; u++ {
				var sumy float64
				var sumu float64
				var sumv float64
				for v := 0; v < 8; v++ {
					var cU float64
					var cV float64
					if u == 0 {
						cU = math.Sqrt(2) / 2
					} else {
						cU = 1
					}
					if v == 0 {
						cV = math.Sqrt(2) / 2
					} else {
						cV = 1
					}
					factor1 := math.Cos(float64((2*i+1)*u)*math.Pi/16)
					factor2 := math.Cos(float64((2*j+1)*v)*math.Pi/16)
					sumy += factor1 * factor2 * float64(F[u][v].yF) * cU * cV / 4
					sumu += factor1 * factor2 * float64(F[u][v].uF) * cU * cV / 4
					sumv += factor1 * factor2 * float64(F[u][v].vF) * cU * cV / 4
				}
				sumY += sumy
				sumU += sumu
				sumV += sumv
			}
			f[i][j].y = int(sumY)
			f[i][j].u = int(sumU)
			f[i][j].v = int(sumV)
		}
	}
	return f
}
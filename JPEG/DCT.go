package JPEG

import (
	"math"
)

func DCT(f [8][8]yuv) (F [8][8]nodeF) {
	for u := 0; u < 8; u++ {
		for v:= 0; v < 8; v++ {
			sumY := 0.0
			sumU := 0.0
			sumV := 0.0
			cU := 0.0
			cV := 0.0
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
			for i := 0; i < 8; i++ {
				sumy := 0.0
				sumu := 0.0
				sumv := 0.0
				for j := 0; j < 8; j++ {
					factor1 := math.Cos((2*float64(i) + 1)*float64(u)*math.Pi/16)
					factor2 := math.Cos((2*float64(j) + 1)*float64(v)*math.Pi/16)
					sumy += factor1 * factor2 * float64(f[i][j].y)
					sumu += factor1 * factor2 * float64(f[i][j].u)
					sumv += factor1 * factor2 * float64(f[i][j].v)
				}
				sumY += sumy
				sumU += sumu
				sumV += sumv
			}
			F[u][v].yF = int(sumY * cU * cV / 4)
			F[u][v].uF = int(sumU * cU * cV / 4)
			F[u][v].vF = int(sumV  * cU * cV / 4)
		}
	}
	return F
}
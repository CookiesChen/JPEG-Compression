package JPEG

import "math"

func antiQuantification(QF [8][8]nodeF)(F [8][8]nodeF) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			F[i][j].yF = int(math.Round(float64(QF[i][j].yF) * float64(tableY[i][j])))
			F[i][j].uF = int(math.Round(float64(QF[i][j].uF) * float64(tableU[i][j])))
			F[i][j].vF = int(math.Round(float64(QF[i][j].vF) * float64(tableU[i][j])))
		}
	}
	return F
}
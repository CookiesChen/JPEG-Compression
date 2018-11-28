package JPEG


func antiQuantification(QF [8][8]nodeF)(F [8][8]nodeF) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			F[i][j].yF = QF[i][j].yF * tableY[i][j]
			F[i][j].uF = QF[i][j].uF * tableU[i][j]
			F[i][j].vF = QF[i][j].vF * tableU[i][j]
		}
	}
	return F
}
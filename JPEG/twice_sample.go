package JPEG

func TwiceSample(arr [][]yuv){
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			if j % 2 == 0 {
				arr[i][j].u = 0
			} else {
				arr[i][j].v = 0
			}
		}
	}
}
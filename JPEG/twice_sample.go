package JPEG

func TwiceSample(arr [][]yuv){
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			if j % 2 == 0 && j != 0 {
				arr[i][j].u = arr[i][j - 1].u
			}
			if j % 2 == 1 {
				arr[i][j].v = arr[i][j - 1].v
			}
		}
	}
}
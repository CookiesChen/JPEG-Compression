package JPEG

func TwiceSample(arr [][]yuv){
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			if j % 2 == 0 && i % 2 == 0 {
				arr[i][j].u = arr[i][j+1].u
				arr[i+1][j].v = arr[i][j].v
				arr[i+1][j].u = arr[i][j].u
				arr[i][j+1].u = arr[i][j].u
				arr[i][j+1].v = arr[i][j].v
				arr[i+1][j+1].u = arr[i][j].u
				arr[i+1][j+1].v = arr[i][j].v
			}
		}
	}
}
package JPEG

func TwiceSample(arr [][]yuv){
	for i := 0; i < width; i++ {
		for j := 0; j < heigth; j++ {
			if j % 2 == 0 {
				//if i % 4 == 0 || i % 4 == 2{
				//	arr[i][j].u = arr[i][j].v
				//} else {
				//	arr[i][j].u = arr[i-1][j].u
				//	arr[i][j].v = arr[i-1][j].v
				//}
			} else {
				//if i % 4 == 0 || i % 4 == 2 {
				//	arr[i][j].v = arr[i][j].u
				//} else {
				//	arr[i][j].u = arr[i-1][j].u
				//	arr[i][j].v = arr[i-1][j].v
				//}
			}
		}
	}
}
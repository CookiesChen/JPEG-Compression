package main

import "github.com/CookiesChen/JPEG-Compression/JPEG"

func main()  {
	F := JPEG.Exec()
	JPEG.Decode(F)
}

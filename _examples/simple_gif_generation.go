package main

import (
	"bunnYmage"
	"fmt"
)

func main() {
	test, _ := bunnYmage.ReadImageFromDisk("bunny_gif_1.jpg")
	test2, _ := bunnYmage.ReadImageFromDisk("bunny_gif_2.jpg")
	test3, _ := bunnYmage.ReadImageFromDisk("bunny_gif_3.jpg")
	err2 := bunnYmage.GenerateGIFAndWriteToDisk(1.5, "bunny.gif", test, test2, test3)
	if err2 != nil {
		fmt.Println(err2)
	}
}

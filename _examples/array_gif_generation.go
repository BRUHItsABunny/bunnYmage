package main

import (
	"bunnYmage"
	"fmt"
	"io/ioutil"
)

func main() {
	test, _ := bunnYmage.ReadImageFromDisk("bunny_gif_1.jpg")
	test2, _ := bunnYmage.ReadImageFromDisk("bunny_gif_2.jpg")
	test3, _ := bunnYmage.ReadImageFromDisk("bunny_gif_3.jpg")
	bunArray, err2 := bunnYmage.GenerateGIFAndWriteToByteArray(1.5, test, test2, test3)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		_ = ioutil.WriteFile("bunny_array.gif", bunArray, 0777)
	}
}

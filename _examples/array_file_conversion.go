package main

import (
	"bunnYmage"
	"fmt"
	"io/ioutil"
)

func main() {
	bunBytes, _ := ioutil.ReadFile("bunny.png")
	test, err := bunnYmage.ReadImageFromByteArray(bunBytes, "bunny", ".png")
	if err == nil {
		bunJPGBytes, err2 := test.ConvertAndWriteToByteArray(bunnYmage.JPG)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			_ = ioutil.WriteFile("bunny_bytearray.jpg", bunJPGBytes, 0777)
		}
	} else {
		fmt.Println(err)
	}
}

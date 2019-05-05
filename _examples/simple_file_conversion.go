package main

import (
	"bunnYmage"
	"fmt"
)

func main() {
	test, err := bunnYmage.ReadImageFromDisk("bunny.jpg")
	if err == nil {
		_, err2 := test.ConvertAndWriteToDisk("bunny.png", bunnYmage.PNG)
		if err2 != nil {
			fmt.Println(err2)
		}
	} else {
		fmt.Println(err)
	}
}

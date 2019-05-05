package bunnYmage

import "image"

type BunnyImage struct {
	ImageType FileType
	Image     image.Image
	FileName  string
	FilePath  string
}

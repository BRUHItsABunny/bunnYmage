package bunnYmage

import "image"

// BunnyImage is the global object for images
type BunnyImage struct {
	ImageType FileType
	Image     image.Image
	FileName  string
	FilePath  string
}

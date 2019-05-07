package bunnYmage

import (
	"errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Check if image exists
func ImageExists(location string) (bool, string, string) {
	var extension string
	var name string
	if stats, err := os.Stat(location); err != nil {
		if os.IsNotExist(err) {
			return false, "", ""
		}
		splitName := strings.Split(stats.Name(), ".")
		extension = "." + splitName[len(splitName)-1]
		name = strings.ReplaceAll(stats.Name(), extension, "")
		extension = strings.ToLower(extension)
	}
	return true, extension, name
}

/*
This function reads an image from disk into a BunnyImage object
It takes a location (eg: "foo/bar/in.png")
It returns a BunnyImage object and an error
*/
func ReadImageFromDisk(location string) (*BunnyImage, error) {
	exists, extension, name := ImageExists(location)
	if !exists {
		return nil, errors.New("file doesn't exist")
	}
	fImage, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	bunImage := new(BunnyImage)
	var decodedImage image.Image
	var decodeErr error
	switch extension {
	case ".png":
		bunImage.ImageType = PNG
		decodedImage, decodeErr = png.Decode(fImage)
	case ".gif":
		bunImage.ImageType = GIF
		decodedImage, decodeErr = gif.Decode(fImage)
	case ".bmp":
		bunImage.ImageType = BMP
		decodedImage, decodeErr = bmp.Decode(fImage)
	case ".tiff":
		bunImage.ImageType = TIFF
		decodedImage, decodeErr = tiff.Decode(fImage)
	case ".wepb":
		bunImage.ImageType = WEBP
		decodedImage, decodeErr = webp.Decode(fImage)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	default:
		bunImage.ImageType = JPG
		decodedImage, decodeErr = jpeg.Decode(fImage)
	}
	_ = fImage.Close()
	if decodeErr != nil {
		return nil, decodeErr
	}
	bunImage.Image = decodedImage
	bunImage.FilePath = filepath.Dir(location)
	bunImage.FileName = name
	return bunImage, nil
}

/*
This function writes a BunnyImage object to disk
It takes a location (eg: "foo/bar/out.png")
It returns an error
*/
func (bunImage BunnyImage) WriteImageToDisk(location string) error {
	var f *os.File
	var fileErr, writeErr error
	if location == "" {
		location = bunImage.FilePath + bunImage.FileName + bunImage.ImageType.Extension()
	}
	if _, err := os.Stat(location); os.IsNotExist(err) {
		f, fileErr = os.Create(location)
	} else {
		f, fileErr = os.Open(location)
	}
	if fileErr != nil {
		return fileErr
	}
	switch bunImage.ImageType {
	case PNG:
		writeErr = png.Encode(f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(f, bunImage.Image, nil)
	}
	_ = f.Close()
	return writeErr
}

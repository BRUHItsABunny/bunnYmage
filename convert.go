package bunnYmage

import (
	"bytes"
	"errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

/*
This function converts a BunnyImage object from one format to another
It takes a location (eg: "foo/bar/out.jpeg") and a FileType object
It returns an error
*/
func (bunImage BunnyImage) ConvertAndWriteToDisk(location string, toType FileType) error {
	var f *os.File
	var fileErr, writeErr error
	if location == "" {
		location = bunImage.FilePath + bunImage.FileName + toType.Extension()
	}
	if _, err := os.Stat(location); os.IsNotExist(err) {
		f, fileErr = os.Create(location)
	} else {
		f, fileErr = os.Open(location)
	}
	if fileErr != nil {
		return fileErr
	}
	switch toType {
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

/*
This function converts a BunnyImage from one format to another and returns the result as a byte array
It takes a FileType object
It returns a byte array and an error
*/
func (bunImage BunnyImage) ConvertAndWriteToByteArray(toType FileType) ([]byte, error) {
	var b bytes.Buffer
	var writeErr error
	f := io.Writer(&b)
	switch toType {
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
	if writeErr == nil {
		return b.Bytes(), writeErr
	}
	return nil, writeErr
}

/*
This function converts a BunnyImage from one format to another and returns the *io.Writer
It takes an *io.Writer and a FileType object
It returns an *io.Writer and an error
*/
func (bunImage BunnyImage) ConvertAndWriteToWriter(f *io.Writer, toType FileType) (*io.Writer, error) {
	var writeErr error
	switch toType {
	case PNG:
		writeErr = png.Encode(*f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(*f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(*f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(*f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(*f, bunImage.Image, nil)
	}
	if writeErr == nil {
		return f, writeErr
	}
	return nil, writeErr
}

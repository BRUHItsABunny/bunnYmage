package bunnYmage

import (
	"bytes"
	"errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

/*
This function reads a BunnyImage from a byte array
It takes a byte array and extension (eg: ".jpeg")
It returns a BunnyImage and an error
*/
func ReadImageFromByteArray(data []byte, name, extension string) (*BunnyImage, error) {
	extension = strings.ToLower(extension)
	fImage := bytes.NewReader(data)
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
	if decodeErr != nil {
		return nil, decodeErr
	}
	bunImage.Image = decodedImage
	bunImage.FilePath = name
	bunImage.FileName = name
	return bunImage, nil
}

/*
This function writes a BunnyImage to a byte array
It takes nothing
It returns a byte array and an error
*/
func (bunImage BunnyImage) WriteImageToByteArray() ([]byte, error) {
	var b bytes.Buffer
	var writeErr error
	f := io.Writer(&b)
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
	if writeErr == nil {
		return b.Bytes(), writeErr
	}
	return nil, writeErr
}

/*
This function reads a BunnyImage object from an *io.Reader
It takes an *io.Reader, name (eg: "in") and an extension (eg: ".png")
It returns a BunnyImage object and an error
*/
func ReadImageFromReader(fImage *io.Reader, name, extension string) (*BunnyImage, error) {
	extension = strings.ToLower(extension)
	bunImage := new(BunnyImage)
	var decodedImage image.Image
	var decodeErr error
	switch extension {
	case ".png":
		bunImage.ImageType = PNG
		decodedImage, decodeErr = png.Decode(*fImage)
	case ".gif":
		bunImage.ImageType = GIF
		decodedImage, decodeErr = gif.Decode(*fImage)
	case ".bmp":
		bunImage.ImageType = BMP
		decodedImage, decodeErr = bmp.Decode(*fImage)
	case ".tiff":
		bunImage.ImageType = TIFF
		decodedImage, decodeErr = tiff.Decode(*fImage)
	case ".wepb":
		bunImage.ImageType = WEBP
		decodedImage, decodeErr = webp.Decode(*fImage)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	default:
		bunImage.ImageType = JPG
		decodedImage, decodeErr = jpeg.Decode(*fImage)
	}
	if decodeErr != nil {
		return nil, decodeErr
	}
	bunImage.Image = decodedImage
	bunImage.FilePath = name
	bunImage.FileName = name
	return bunImage, nil
}

/*
This function writes a BunnyImage object to an *io.Writer
It takes an *io.Writer
It returns an *io.Writer and error
*/
func (bunImage BunnyImage) WriteImageToWriter(f *io.Writer) (*io.Writer, error) {
	var writeErr error
	switch bunImage.ImageType {
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

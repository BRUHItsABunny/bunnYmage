package bunnYmage

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"io"
	"math"
	"os"
)

/*
This function generates a GIF from multiple BunnyImage objects
It takes an interval in SECONDS, a location (eg: "foo/bar/out.gif") and several BunnyImage objects
It returns an error
*/
func GenerateGIFAndWriteToDisk(interval float64, location string, images ...*BunnyImage) error {
	actualInterval := int(math.RoundToEven(100.0 * interval))
	var gifs []BunnyImage
	for _, i := range images {
		if i.ImageType == GIF {
			gifs = append(gifs, *i)
		} else {
			conBytes, err := i.ConvertAndWriteToByteArray(GIF)
			if err == nil {
				conImage, err := ReadImageFromByteArray(conBytes, i.FileName, ".gif")
				if err == nil {
					gifs = append(gifs, *conImage)
				}
			}
		}
	}
	if len(gifs) > 1 {
		outGif := &gif.GIF{}
		for _, inGif := range gifs {
			outGif.Image = append(outGif.Image, inGif.Image.(*image.Paletted))
			outGif.Delay = append(outGif.Delay, actualInterval)
		}
		var f *os.File
		var fileErr error
		if _, err := os.Stat(location); os.IsNotExist(err) {
			f, fileErr = os.Create(location)
		} else {
			f, fileErr = os.Open(location)
		}
		if fileErr != nil {
			return fileErr
		}
		err := gif.EncodeAll(f, outGif)
		_ = f.Close()
		return err
	}
	return errors.New("no images to merge into GIF")
}

/*
This function generates a GIF from multiple BunnyImage objects
It takes an interval in SECONDS and several BunnyImage objects
It returns a byte array along with an error
*/
func GenerateGIFAndWriteToByteArray(interval float64, images ...*BunnyImage) ([]byte, error) {
	actualInterval := int(math.RoundToEven(100.0 * interval))
	var gifs []BunnyImage
	for _, i := range images {
		if i.ImageType == GIF {
			gifs = append(gifs, *i)
		} else {
			conBytes, err := i.ConvertAndWriteToByteArray(GIF)
			if err == nil {
				conImage, err := ReadImageFromByteArray(conBytes, i.FileName, ".gif")
				if err == nil {
					gifs = append(gifs, *conImage)
				}
			}
		}
	}
	if len(gifs) > 1 {
		outGif := &gif.GIF{}
		for _, inGif := range gifs {
			outGif.Image = append(outGif.Image, inGif.Image.(*image.Paletted))
			outGif.Delay = append(outGif.Delay, actualInterval)
		}
		var b bytes.Buffer
		f := io.Writer(&b)
		err := gif.EncodeAll(f, outGif)
		return b.Bytes(), err
	}
	return nil, errors.New("no images to merge into GIF")
}

/*
This function generates a GIF from multiple BunnyImage objects
It takes an interval in SECONDS, *io.Writer and several BunnyImage ibjects
It returns the *io.Writer along with an error
*/
func GenerateGIFAndWriteToWriter(interval float64, f *io.Writer, images ...*BunnyImage) (*io.Writer, error) {
	actualInterval := int(math.RoundToEven(100.0 * interval))
	var gifs []BunnyImage
	for _, i := range images {
		if i.ImageType == GIF {
			gifs = append(gifs, *i)
		} else {
			conBytes, err := i.ConvertAndWriteToByteArray(GIF)
			if err == nil {
				conImage, err := ReadImageFromByteArray(conBytes, i.FileName, ".gif")
				if err == nil {
					gifs = append(gifs, *conImage)
				}
			}
		}
	}
	if len(gifs) > 1 {
		outGif := &gif.GIF{}
		for _, inGif := range gifs {
			outGif.Image = append(outGif.Image, inGif.Image.(*image.Paletted))
			outGif.Delay = append(outGif.Delay, actualInterval)
		}
		err := gif.EncodeAll(*f, outGif)
		return f, err
	}
	return nil, errors.New("no images to merge into GIF")
}

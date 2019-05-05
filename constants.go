package bunnYmage

type FileType int

const (
	JPG FileType = iota
	PNG
	GIF
	BMP
	TIFF
	WEBP
)

func (fT FileType) Extension() string {
	switch fT {
	case JPG:
		return ".jpg"
	case PNG:
		return ".png"
	case GIF:
		return ".gif"
	case BMP:
		return ".bmp"
	case TIFF:
		return ".tiff"
	case WEBP:
		return ".webp"
	default:
		//... what else?
		return ".bunny"
	}
}

func (fT FileType) String() string {
	switch fT {
	case JPG:
		return "jpg"
	case PNG:
		return "png"
	case GIF:
		return "gif"
	case BMP:
		return "bmp"
	case TIFF:
		return "tiff"
	case WEBP:
		return "webp"
	default:
		//... what else?
		return "bunny"
	}
}

func GetFileType(extension string) FileType {
	switch extension {
	case ".jpeg":
		fallthrough
	case ".jpg":
		return JPG
	case ".png":
		return PNG
	case ".gif":
		return GIF
	case ".bmp":
		return BMP
	case ".tiff":
		return TIFF
	case ".wepb":
		return WEBP
	default:
		return JPG
	}
}

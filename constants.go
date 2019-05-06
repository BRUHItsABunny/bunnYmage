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

func (fT FileType) MIMEType() string {
	switch fT {
	case JPG:
		return "image/jpeg"
	case PNG:
		return "image/png"
	case GIF:
		return "image/gif"
	case BMP:
		return "image/bmp"
	case TIFF:
		return "image/tiff"
	case WEBP:
		return "image/webp"
	default:
		//... what else?
		return "animal/bunny"
	}
}

func GetFileTypeByExtension(extension string) FileType {
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

func GetFileTypeByMimeType(mime string) FileType {
	switch mime {
	case "image/jpeg":
		fallthrough
	case "image/jpg":
		// Technically invalid
		return JPG
	case "image/png":
		return PNG
	case "image/gif":
		return GIF
	case "image/bmp":
		return BMP
	case "image/tiff":
		return TIFF
	case "image/wepb":
		return WEBP
	default:
		return JPG
	}
}

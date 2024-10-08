package ansicode

import "image/color"

var _ color.Color = RGBColor{}

// IndexedColor is a color that can be represented by an index.
type IndexedColor struct {
	Index byte
}

// RGBColor is a color that can be represented by RGB values.
type RGBColor struct {
	R, G, B uint8
}

// NamedColor is a color that can be referenced by name.
type NamedColor int

// NamedColor values.
const (
	NamedColorBlack         NamedColor = 0
	NamedColorRed           NamedColor = 1
	NamedColorGreen         NamedColor = 2
	NamedColorYellow        NamedColor = 3
	NamedColorBlue          NamedColor = 4
	NamedColorMagenta       NamedColor = 5
	NamedColorCyan          NamedColor = 6
	NamedColorWhite         NamedColor = 7
	NamedColorBrightBlack   NamedColor = 8
	NamedColorBrightRed     NamedColor = 9
	NamedColorBrightGreen   NamedColor = 10
	NamedColorBrightYellow  NamedColor = 11
	NamedColorBrightBlue    NamedColor = 12
	NamedColorBrightMagenta NamedColor = 13
	NamedColorBrightCyan    NamedColor = 14
	NamedColorBrightWhite   NamedColor = 15

	NamedColorCubeStart NamedColor = 16
	NamedColorCubeEnd   NamedColor = 232

	NamedColorGrayscaleStart NamedColor = 233
	NamedColorGrayscaleEnd   NamedColor = 255

	NamedColorForeground NamedColor = 256
	NamedColorBackground NamedColor = 257
	NamedColorCursor     NamedColor = 258

	NamedColorDimBlack   NamedColor = 259
	NamedColorDimRed     NamedColor = 260
	NamedColorDimGreen   NamedColor = 261
	NamedColorDimYellow  NamedColor = 262
	NamedColorDimBlue    NamedColor = 263
	NamedColorDimMagenta NamedColor = 264
	NamedColorDimCyan    NamedColor = 265
	NamedColorDimWhite   NamedColor = 266

	NamedColorBrightForeground NamedColor = 267
	NamedColorDimForeground    NamedColor = 268
)

// ToBright returns the bright version of the color.
func (c NamedColor) ToBright() NamedColor {
	switch c {
	case NamedColorForeground:
		return NamedColorBrightForeground
	case NamedColorBlack:
		return NamedColorBrightBlack
	case NamedColorRed:
		return NamedColorBrightRed
	case NamedColorGreen:
		return NamedColorBrightGreen
	case NamedColorYellow:
		return NamedColorBrightYellow
	case NamedColorBlue:
		return NamedColorBrightBlue
	case NamedColorMagenta:
		return NamedColorBrightMagenta
	case NamedColorCyan:
		return NamedColorBrightCyan
	case NamedColorWhite:
		return NamedColorBrightWhite
	case NamedColorDimForeground:
		return NamedColorForeground
	case NamedColorDimBlack:
		return NamedColorBlack
	case NamedColorDimRed:
		return NamedColorRed
	case NamedColorDimGreen:
		return NamedColorGreen
	case NamedColorDimYellow:
		return NamedColorYellow
	case NamedColorDimBlue:
		return NamedColorBlue
	case NamedColorDimMagenta:
		return NamedColorMagenta
	case NamedColorDimCyan:
		return NamedColorCyan
	case NamedColorDimWhite:
		return NamedColorWhite
	default:
		return c
	}
}

// ToDim returns the dim version of the color.
func (c NamedColor) ToDim() NamedColor {
	switch c {
	case NamedColorBlack:
		return NamedColorDimBlack
	case NamedColorRed:
		return NamedColorDimRed
	case NamedColorGreen:
		return NamedColorDimGreen
	case NamedColorYellow:
		return NamedColorDimYellow
	case NamedColorBlue:
		return NamedColorDimBlue
	case NamedColorMagenta:
		return NamedColorDimMagenta
	case NamedColorCyan:
		return NamedColorDimCyan
	case NamedColorWhite:
		return NamedColorDimWhite
	case NamedColorForeground:
		return NamedColorDimForeground
	case NamedColorBrightBlack:
		return NamedColorBlack
	case NamedColorBrightRed:
		return NamedColorRed
	case NamedColorBrightGreen:
		return NamedColorGreen
	case NamedColorBrightYellow:
		return NamedColorYellow
	case NamedColorBrightBlue:
		return NamedColorBlue
	case NamedColorBrightMagenta:
		return NamedColorMagenta
	case NamedColorBrightCyan:
		return NamedColorCyan
	case NamedColorBrightWhite:
		return NamedColorWhite
	default:
		return c
	}
}

// RGBA returns the RGBA value of the color.
func (c RGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) * 0x101
	g = uint32(c.G) * 0x101
	b = uint32(c.B) * 0x101
	a = 0xffff
	return
}

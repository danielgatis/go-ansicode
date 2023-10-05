package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamedColor_ToBright(t *testing.T) {
	tests := []struct {
		name string
		c    NamedColor
		want NamedColor
	}{
		{
			name: "Foreground",
			c:    NamedColorForeground,
			want: NamedColorBrightForeground,
		},
		{
			name: "Black",
			c:    NamedColorBlack,
			want: NamedColorBrightBlack,
		},
		{
			name: "Red",
			c:    NamedColorRed,
			want: NamedColorBrightRed,
		},
		{
			name: "Green",
			c:    NamedColorGreen,
			want: NamedColorBrightGreen,
		},
		{
			name: "Yellow",
			c:    NamedColorYellow,
			want: NamedColorBrightYellow,
		},
		{
			name: "Blue",
			c:    NamedColorBlue,
			want: NamedColorBrightBlue,
		},
		{
			name: "Magenta",
			c:    NamedColorMagenta,
			want: NamedColorBrightMagenta,
		},
		{
			name: "Cyan",
			c:    NamedColorCyan,
			want: NamedColorBrightCyan,
		},
		{
			name: "White",
			c:    NamedColorWhite,
			want: NamedColorBrightWhite,
		},
		{
			name: "DimForeground",
			c:    NamedColorDimForeground,
			want: NamedColorForeground,
		},
		{
			name: "DimBlack",
			c:    NamedColorDimBlack,
			want: NamedColorBlack,
		},
		{
			name: "DimRed",
			c:    NamedColorDimRed,
			want: NamedColorRed,
		},
		{
			name: "DimGreen",
			c:    NamedColorDimGreen,
			want: NamedColorGreen,
		},
		{
			name: "DimYellow",
			c:    NamedColorDimYellow,
			want: NamedColorYellow,
		},
		{
			name: "DimBlue",
			c:    NamedColorDimBlue,
			want: NamedColorBlue,
		},
		{
			name: "DimMagenta",
			c:    NamedColorDimMagenta,
			want: NamedColorMagenta,
		},
		{
			name: "DimCyan",
			c:    NamedColorDimCyan,
			want: NamedColorCyan,
		},
		{
			name: "DimWhite",
			c:    NamedColorDimWhite,
			want: NamedColorWhite,
		},
		{
			name: "Unknown",
			c:    NamedColor(42),
			want: NamedColor(42),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.c.ToBright())
		})
	}
}
func TestNamedColor_ToDim(t *testing.T) {
	tests := []struct {
		name string
		c    NamedColor
		want NamedColor
	}{
		{
			name: "Foreground",
			c:    NamedColorForeground,
			want: NamedColorDimForeground,
		},
		{
			name: "Black",
			c:    NamedColorBlack,
			want: NamedColorDimBlack,
		},
		{
			name: "Red",
			c:    NamedColorRed,
			want: NamedColorDimRed,
		},
		{
			name: "Green",
			c:    NamedColorGreen,
			want: NamedColorDimGreen,
		},
		{
			name: "Yellow",
			c:    NamedColorYellow,
			want: NamedColorDimYellow,
		},
		{
			name: "Blue",
			c:    NamedColorBlue,
			want: NamedColorDimBlue,
		},
		{
			name: "Magenta",
			c:    NamedColorMagenta,
			want: NamedColorDimMagenta,
		},
		{
			name: "Cyan",
			c:    NamedColorCyan,
			want: NamedColorDimCyan,
		},
		{
			name: "White",
			c:    NamedColorWhite,
			want: NamedColorDimWhite,
		},
		{
			name: "BrightBlack",
			c:    NamedColorBrightBlack,
			want: NamedColorBlack,
		},
		{
			name: "BrightRed",
			c:    NamedColorBrightRed,
			want: NamedColorRed,
		},
		{
			name: "BrightGreen",
			c:    NamedColorBrightGreen,
			want: NamedColorGreen,
		},
		{
			name: "BrightYellow",
			c:    NamedColorBrightYellow,
			want: NamedColorYellow,
		},
		{
			name: "BrightBlue",
			c:    NamedColorBrightBlue,
			want: NamedColorBlue,
		},
		{
			name: "BrightMagenta",
			c:    NamedColorBrightMagenta,
			want: NamedColorMagenta,
		},
		{
			name: "BrightCyan",
			c:    NamedColorBrightCyan,
			want: NamedColorCyan,
		},
		{
			name: "BrightWhite",
			c:    NamedColorBrightWhite,
			want: NamedColorWhite,
		},
		{
			name: "Unknown",
			c:    NamedColor(42),
			want: NamedColor(42),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.c.ToDim())
		})
	}
}

package color

import (
	"math"
	"strconv"
)

// Color represents an RGB color.
type Color struct {
	R int
	G int
	B int
}

// ANSI code representation of some common colors
const (
	BLACK   = "0"
	RED     = "1"
	GREEN   = "2"
	YELLOW  = "3"
	BLUE    = "4"
	MAGENTA = "5"
	CYAN    = "6"
	WHITE   = "7"
)

// ANSI code representations
const (
	ansiReset               = "\033[0m"
	ansiStart               = "\033["
	ansi256ForegroundPrefix = "38;5;"
	ansi256BackgroundPrefix = "48;5;"
)

// Function to calculate the color distance
func colorDistance(c1 Color, c2 Color) float64 {
	rDiff := c1.R - c2.R
	gDiff := c1.G - c2.G
	bDiff := c1.B - c2.B
	return math.Sqrt(float64(rDiff*rDiff + gDiff*gDiff + bDiff*bDiff))
}

// Function to find the nearest ANSI color (256 colors)
func nearestANSI256Color(color Color) int {
	var nearestIndex int
	minDistance := math.MaxFloat64

	for i := 0; i < 256; i++ {
		r, g, b := ansi256ToRGB(i)
		ansiColor := Color{R: r, G: g, B: b}
		distance := colorDistance(color, ansiColor)

		if distance < minDistance {
			minDistance = distance
			nearestIndex = i
		}
	}

	return nearestIndex
}

// Function to convert ANSI 256 color index to RGB
func ansi256ToRGB(code int) (r, g, b int) {
	if code < 16 {
		// Standard and high-intensity colors
		r, g, b = map[int]int{
			0: 0, 1: 128, 2: 0, 3: 128, 4: 0, 5: 128, 6: 0, 7: 192,
			8: 128, 9: 255, 10: 128, 11: 255, 12: 128, 13: 255, 14: 128, 15: 255,
		}[code%16], map[int]int{
			0: 0, 1: 0, 2: 128, 3: 128, 4: 0, 5: 0, 6: 128, 7: 192,
			8: 128, 9: 128, 10: 255, 11: 255, 12: 0, 13: 0, 14: 255, 15: 255,
		}[code%16], map[int]int{
			0: 0, 1: 0, 2: 0, 3: 0, 4: 128, 5: 128, 6: 128, 7: 192,
			8: 128, 9: 128, 10: 128, 11: 128, 12: 255, 13: 255, 14: 255, 15: 255,
		}[code%16]
	} else if code < 232 {
		// 6x6x6 color cube
		code -= 16
		r = ((code / 36) % 6) * 255 / 5
		g = ((code / 6) % 6) * 255 / 5
		b = (code % 6) * 255 / 5
	} else if code < 256 {
		// Grayscale ramp
		gray := (code-232)*10 + 8
		r, g, b = gray, gray, gray
	}
	return
}

// RgbToANSI256 Function to convert RGB to ANSI 256 color code
func RgbToANSI256(color Color) string {
	ansiCode := nearestANSI256Color(color)
	return strconv.Itoa(ansiCode)
}

// FormatWithANSI Function to format text with ANSI color
func FormatWithANSI(text string, foregroundColorCode string, backgroundColorCode string, foreground bool, background bool) string {
	var ansi string
	if foreground {
		ansi += ansiStart + ansi256ForegroundPrefix + foregroundColorCode + "m"
	}
	if background {
		ansi += ansiStart + ansi256BackgroundPrefix + backgroundColorCode + "m"
	}
	return ansi + text + ansiReset
}

func FormatWithRGB(text string, foregroundColor Color, backgroundColor Color, foreground bool, background bool) string {
	foregroundColorCode := RgbToANSI256(foregroundColor)
	backgroundColorCode := RgbToANSI256(backgroundColor)

	return FormatWithANSI(text, foregroundColorCode, backgroundColorCode, foreground, background)
}

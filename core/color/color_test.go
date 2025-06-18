package color

import (
	"math"
	"testing"
)

func TestColorDistance(t *testing.T) {
	tests := []struct {
		name     string
		color1   Color
		color2   Color
		expected float64
	}{
		{
			name:     "Same color",
			color1:   Color{R: 255, G: 255, B: 255},
			color2:   Color{R: 255, G: 255, B: 255},
			expected: 0,
		},
		{
			name:     "Different colors",
			color1:   Color{R: 255, G: 0, B: 0},
			color2:   Color{R: 0, G: 255, B: 0},
			expected: math.Sqrt(130050), // √(255² + (-255)² + 0²)
		},
		{
			name:     "Black and white",
			color1:   Color{R: 0, G: 0, B: 0},
			color2:   Color{R: 255, G: 255, B: 255},
			expected: math.Sqrt(195075), // √(255² + 255² + 255²)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := colorDistance(tt.color1, tt.color2)
			if math.Abs(got-tt.expected) > 0.0001 {
				t.Errorf("colorDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAnsi256ToRGB(t *testing.T) {
	tests := []struct {
		name  string
		code  int
		wantR int
		wantG int
		wantB int
	}{
		{
			name:  "Basic black (0)",
			code:  0,
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "Basic red (1)",
			code:  1,
			wantR: 128,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "High-intensity white (15)",
			code:  15,
			wantR: 255,
			wantG: 255,
			wantB: 255,
		},
		{
			name:  "Color cube (216)",
			code:  216,
			wantR: 255,
			wantG: 153,
			wantB: 102,
		},
		{
			name:  "Grayscale (232)",
			code:  232,
			wantR: 8,
			wantG: 8,
			wantB: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := ansi256ToRGB(tt.code)
			if gotR != tt.wantR || gotG != tt.wantG || gotB != tt.wantB {
				t.Errorf("ansi256ToRGB() = (%v, %v, %v), want (%v, %v, %v)",
					gotR, gotG, gotB, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}
}

func TestRgbToANSI256(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{
			name:     "Pure black",
			color:    Color{R: 0, G: 0, B: 0},
			expected: "0",
		},
		{
			name:     "Pure white",
			color:    Color{R: 255, G: 255, B: 255},
			expected: "15",
		},
		{
			name:     "Pure red",
			color:    Color{R: 255, G: 0, B: 0},
			expected: "196",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RgbToANSI256(tt.color)
			if got != tt.expected {
				t.Errorf("RgbToANSI256() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatWithANSI(t *testing.T) {
	tests := []struct {
		name            string
		text            string
		foregroundColor string
		backgroundColor string
		foreground      bool
		background      bool
		expected        string
	}{
		{
			name:            "Foreground only",
			text:            "test",
			foregroundColor: "1",
			backgroundColor: "",
			foreground:      true,
			background:      false,
			expected:        "\033[38;5;1mtest\033[0m",
		},
		{
			name:            "Background only",
			text:            "test",
			foregroundColor: "",
			backgroundColor: "2",
			foreground:      false,
			background:      true,
			expected:        "\033[48;5;2mtest\033[0m",
		},
		{
			name:            "Both colors",
			text:            "test",
			foregroundColor: "1",
			backgroundColor: "2",
			foreground:      true,
			background:      true,
			expected:        "\033[38;5;1m\033[48;5;2mtest\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatWithANSI(tt.text, tt.foregroundColor, tt.backgroundColor, tt.foreground, tt.background)
			if got != tt.expected {
				t.Errorf("FormatWithANSI() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatWithRGB(t *testing.T) {
	tests := []struct {
		name            string
		text            string
		foregroundColor Color
		backgroundColor Color
		foreground      bool
		background      bool
		expected        string
	}{
		{
			name:            "Black text on white background",
			text:            "test",
			foregroundColor: Color{R: 0, G: 0, B: 0},
			backgroundColor: Color{R: 255, G: 255, B: 255},
			foreground:      true,
			background:      true,
			expected:        "\033[38;5;0m\033[48;5;15mtest\033[0m",
		},
		{
			name:            "Red text only",
			text:            "test",
			foregroundColor: Color{R: 255, G: 0, B: 0},
			backgroundColor: Color{},
			foreground:      true,
			background:      false,
			expected:        "\033[38;5;196mtest\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatWithRGB(tt.text, tt.foregroundColor, tt.backgroundColor, tt.foreground, tt.background)
			if got != tt.expected {
				t.Errorf("FormatWithRGB() = %v, want %v", got, tt.expected)
			}
		})
	}
}

package main

import (
	"fmt"
	"ragol-cli/core"
	"ragol-cli/core/color"
	"ragol-cli/core/command"
	"ragol-cli/core/flag"
)

func testColors() {
	red := color.Color{R: 255, G: 0, B: 0}
	green := color.Color{R: 0, G: 255, B: 0}
	blue := color.Color{R: 0, G: 0, B: 255}

	// Convert RGB to ANSI 256 color codes
	ansiRed := color.RgbToANSI256(red)
	ansiGreen := color.RgbToANSI256(green)
	ansiBlue := color.RgbToANSI256(blue)
	ansiTest := color.RgbToANSI256(color.Color{R: 95, G: 255, B: 255})

	// Format text with ANSI colors
	formattedRed := color.FormatWithANSI("This is red", ansiRed, ansiGreen, true, true)
	formattedGreen := color.FormatWithANSI("This is green", ansiGreen, ansiBlue, true, true)
	formattedBlue := color.FormatWithANSI("This is blue", ansiBlue, ansiRed, true, true)
	formattedTest := color.FormatWithANSI("This is test", ansiTest, ansiGreen, true, true)
	formattedRedTest := color.FormatWithANSI("This is red const test", color.RED, color.GREEN, true, true)

	formattedRGBTest := color.FormatWithRGB("This is red with green background", red, green, true, true)

	// Print the formatted text
	fmt.Println(formattedRed)
	fmt.Println(formattedGreen)
	fmt.Println(formattedBlue)
	fmt.Println(formattedTest)
	fmt.Println(formattedRedTest)
	fmt.Println(formattedRGBTest)
}

func main() {
	app := core.NewApp(
		"test cli",
		"0.0.1",
		[]command.BaseCommand{
			*command.NewCommand(
				"test",
				testCommand,
				[]flag.Flag{
					*flag.NewFlag("test", "t", false),
				}),
			*command.NewCommand(
				"test2",
				testCommand2,
				[]flag.Flag{
					*flag.NewFlag("test", "t", false),
				}),
		},
	)
	app.Run()
}

func testCommand(flags []flag.Flag) error {
	fmt.Println("Testing command runs")
	for _, f := range flags {
		fmt.Println(f.Name)
	}
	return nil
}

func testCommand2(flags []flag.Flag) error {
	fmt.Println("Testing command 2 runs")
	return nil
}

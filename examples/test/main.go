package main

import (
	"fmt"
	"github.com/ragolsnagol/ragol-cli/core"
	"github.com/ragolsnagol/ragol-cli/core/action"
	"github.com/ragolsnagol/ragol-cli/core/color"
	"github.com/ragolsnagol/ragol-cli/core/command"
	"github.com/ragolsnagol/ragol-cli/core/context"
	"github.com/ragolsnagol/ragol-cli/core/flag"
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
	f, err := flag.NewFlag("--task", "-t", true, true)
	if err != nil {
		panic(err)
	}
	f2, err := flag.NewFlag("--test", "-t", false, true)
	if err != nil {
		panic(err)
	}

	app := core.NewApp(
		"test cli",
		"0.0.1",
		[]command.BaseCommand{
			*command.NewCommand(
				"test",
				"Test command",
				action.NewAction(testCommand),
				[]flag.Flag{
					*f,
				}),
			*command.NewCommand(
				"test2",
				"Test command 2",
				action.NewAction(testCommand2),
				[]flag.Flag{
					*f2,
				}),
		},
	)
	err = app.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func testCommand(ctx context.Context) error {
	fmt.Println("Testing command runs")
	for _, f := range ctx.Flags {
		fmt.Println(f.Name)
	}
	return nil
}

func testCommand2(ctx context.Context) error {
	fmt.Println("Testing command 2 runs")
	for _, f := range ctx.Flags {
		fmt.Printf("%v: %v\n", f.Name, f.Value)
	}
	return nil
}

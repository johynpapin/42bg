package main

import (
	"os"
	"image"
	_ "image/jpeg"
	_ "image/gif"
	"errors"
	"image/color"
	"image/draw"
	"log"
	"image/png"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]

	if len(args) < 3 {
		return errors.New("you need to pass an image file, an output file and a number of kanjis in argument")
	}

	wallpaperPath := args[0]             // Path to the base wallpaper
	outputPath := args[1]                // Output path
	kanjis, err := strconv.Atoi(args[2]) // Number of kanji studied in the KKLC book
	if err != nil {
		return err
	}

	reader, err := os.Open(wallpaperPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	wallpaper, _, err := image.Decode(reader)

	bounds := wallpaper.Bounds()
	output := image.NewRGBA(bounds)

	draw.Draw(output, bounds, wallpaper, image.ZP, draw.Over)

	rouge := color.RGBA{R: 200}

	draw.DrawMask(output, image.Rect(bounds.Min.X, bounds.Min.Y, kanjis*bounds.Max.X/2300, bounds.Max.Y), &image.Uniform{rouge}, image.ZP, image.NewUniform(color.Alpha{128}), image.ZP, draw.Over) // Added a red rectangle showing the progress of learning on the basic wallpaper

	outputFile, _ := os.Create(outputPath)
	png.Encode(outputFile, output)

	return nil
}

//       ___  _____  ____
//      / _ \/  _/ |/_/ /____ ______ _
//     / ___// /_>  </ __/ -_) __/  ' \
//    /_/  /___/_/|_|\__/\__/_/ /_/_/_/
//
//    Copyright 2017 Eliuk Blau
//
//    This Source Code Form is subject to the terms of the Mozilla Public
//    License, v. 2.0. If a copy of the MPL was not distributed with this
//    file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/eliukblau/pixterm/ansimage"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	pxtVersion = "1.0.0"
	pxtLogo    = `

   ___  _____  ____
  / _ \/  _/ |/_/ /____ ______ _    Made with love by Eliuk Blau
 / ___// /_>  </ __/ -_) __/  ' \   github.com/eliukblau/pixterm
/_/  /___/_/|_|\__/\__/_/ /_/_/_/   v{{VERSION}}

`
)

var (
	flagRows  uint
	flagCols  uint
	flagScale uint
)

func main() {
	validateFlags()
	checkTerminal()
	runPixterm()
}

func printLogo() {
	fmt.Print(strings.Trim(strings.Replace(pxtLogo, "{{VERSION}}", pxtVersion, 1), "\n"), "\n\n")
}

func throwError(code int, v ...interface{}) {
	printLogo()
	log.New(os.Stderr, "[PIXTERM ERROR] ", log.LstdFlags).Println(v...)
	os.Exit(code)
}

func configureFlags() {
	flag.CommandLine.Usage = func() {
		printLogo()

		_, file := filepath.Split(os.Args[0])
		fmt.Print("USAGE:\n\n")
		fmt.Printf("  %v [options] image (JPEG, PNG, GIF, BMP, TIFF, WebP)\n\n", file)

		fmt.Print("OPTIONS:\n\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.CommandLine.PrintDefaults()
		flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
		fmt.Println()
	}

	flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)

	flag.CommandLine.UintVar(&flagScale, "s", 0, "scale `method`:\n\t  0 - resize (default)\n\t  1 - fill\n\t  2 - fit")
	flag.CommandLine.UintVar(&flagRows, "tr", 0, "terminal `rows` (optional, >=2)")
	flag.CommandLine.UintVar(&flagCols, "tc", 0, "terminal `columns` (optional, >=2)")

	flag.CommandLine.Parse(os.Args[1:])
}

func validateFlags() {
	if flag.CommandLine.Arg(0) == "" {
		flag.CommandLine.Usage()
		os.Exit(2)
	}
	if flagScale != 0 && flagScale != 1 && flagScale != 2 {
		flag.CommandLine.Usage()
		os.Exit(2)
	}
	if (flagRows > 0 && flagRows < 2) || (flagCols > 0 && flagCols < 2) {
		flag.CommandLine.Usage()
		os.Exit(2)
	}
}

func checkTerminal() {
	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		throwError(1, "Not running on terminal :(")
	}
}

func getTerminalSize() (width, height int, err error) {
	return terminal.GetSize(int(os.Stdout.Fd()))
}

func runPixterm() {
	var pix *ansimage.ANSImage
	var err error

	// get terminal size
	tx, ty, err := getTerminalSize()
	if err != nil {
		throwError(1, err)
	}

	// use custom terminal size (if applies)
	if flagRows != 0 {
		ty = int(flagRows) + 1
	}
	if flagCols != 0 {
		tx = int(flagCols)
	}

	// set scale mode and create new ANSImage from file
	file := flag.CommandLine.Arg(0)
	switch flagScale {
	case 0:
		pix, err = ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeResize, file)
	case 1:
		pix, err = ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeFill, file)
	case 2:
		pix, err = ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeFit, file)
	}
	if err != nil {
		throwError(1, err)
	}

	// draw ANSImage to terminal
	ansimage.ClearTerminal()
	pix.SetMaxProcs(runtime.NumCPU()) // maximum number of parallel goroutines!
	pix.Draw()
	fmt.Println()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use paralelism for goroutines!
	configureFlags()
}

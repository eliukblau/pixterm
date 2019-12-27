//       ___  _____  ____
//      / _ \/  _/ |/_/ /____ ______ _
//     / ___// /_>  </ __/ -_) __/  ' \
//    /_/  /___/_/|_|\__/\__/_/ /_/_/_/
//
//    Copyright 2019 Eliuk Blau
//
//    This Source Code Form is subject to the terms of the Mozilla Public
//    License, v. 2.0. If a copy of the MPL was not distributed with this
//    file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import "fmt"

func printContributors() {
	fmt.Print("CONTRIBUTORS:\n\n")

	fmt.Print("  > @disq - https://github.com/disq\n")
	fmt.Print("      + Original code for image transparency support.\n")
	fmt.Println()

	fmt.Print("  > @timob - https://github.com/timob\n")
	fmt.Print("      + Fix for ANSIpixel type: use 8bit color component for output.\n")
	fmt.Println()

	fmt.Print("  > @HongjiangHuang - https://github.com/HongjiangHuang\n")
	fmt.Print("      + Original code for image download support.\n")
	fmt.Println()

	fmt.Print("  > @brutestack - https://github.com/brutestack\n")
	fmt.Print("      + Color support for Windows (Command Prompt & PowerShell).\n")
	fmt.Print("      + Original code for disable background color in dithering mode.\n")
	fmt.Print("      + Original code for output Go code to 'fmt.Print()' the image.\n")
	fmt.Println()

	fmt.Print("  > @diamondburned - https://github.com/diamondburned\n")
	fmt.Print("      + NewFromImage() & NewScaledFromImage() for ANSImage API.\n")
	fmt.Println()

	fmt.Print("  > @MichaelMure - https://github.com/MichaelMure\n")
	fmt.Print("      + More conventional 'go.mod' file at repository.\n")
	fmt.Println()

	fmt.Print("  > @Calinou - https://github.com/Calinou\n")
	fmt.Print("      + Use HTTPS URLs everywhere.\n")
	fmt.Print("      + Other awesome contributions.\n")
	fmt.Println()

	fmt.Print("  > @danirod - https://github.com/danirod\n")
	fmt.Print("  > @Xpktro - https://github.com/Xpktro\n")
	fmt.Print("      + Moral support.\n")
	fmt.Println()
}

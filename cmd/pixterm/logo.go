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

import (
	"regexp"
	"strings"

	_ "embed"
)

var (
	//go:embed VERSION
	pxtVersion string

	pxtLogo = `
   ___  _____  ____
  / _ \/  _/ |/_/ /____ ______ _      Made with love by Eliuk Blau
 / ___// /_>  </ __/ -_) __/  ' \ https://github.com/eliukblau/pixterm
/_/  /___/_/|_|\__/\__/_/ /_/_/_/                {{VERSION}}
`
)

func prepareLogoStuff() {
	rx, err := regexp.Compile(`\s+`)
	if err != nil {
		throwError(1, err)
	}
	pxtVersion = rx.ReplaceAllString(pxtVersion, "")
	pxtLogo = strings.Trim(strings.Replace(pxtLogo, "{{VERSION}}", pxtVersion, 1), "\n")
}
